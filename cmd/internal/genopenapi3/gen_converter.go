package genopenapi3

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"go.acuvity.ai/regolithe/spec"
)

const (
	paramNameID    = "id"
	defaultDocName = "toplevel"
)

type converter struct {
	skipPrivateModels bool
	splitOutput       bool
	inSpecSet         spec.SpecificationSet
	resourceToRest    map[string]string
	tagsForModel      map[string]openapi3.Tags
	globalTagSet      map[string]*openapi3.Tag
	outRootDoc        openapi3.T
}

func newConverter(inSpecSet spec.SpecificationSet, cfg Config) *converter {
	c := &converter{
		skipPrivateModels: cfg.Public,
		splitOutput:       cfg.SplitOutput,
		inSpecSet:         inSpecSet,
		resourceToRest:    make(map[string]string),
		tagsForModel:      make(map[string]openapi3.Tags),
		globalTagSet:      make(map[string]*openapi3.Tag),
		outRootDoc:        newOpenAPI3Template(inSpecSet.Configuration()),
	}

	for _, spec := range inSpecSet.Specifications() {
		model := spec.Model()
		c.resourceToRest[model.ResourceName] = model.RestName
	}

	return c
}

func (c *converter) Do(newWriter func(name string) (io.WriteCloser, error)) error {

	// register external types within the components first
	for _, s := range c.inSpecSet.Specifications() {
		if err := c.processSpecComponentRegistryExtension(s); err != nil {
			return fmt.Errorf("unable to process spec component registry extension: %w", err)
		}
	}

	for _, s := range c.inSpecSet.Specifications() {
		if err := c.processSpec(s); err != nil {
			return fmt.Errorf("unable to to process spec: %w", err)
		}
		c.cacheTags(s.Model())
	}

	for name, doc := range c.convertedDocs() {
		dest, err := newWriter(name)
		if err != nil {
			return fmt.Errorf("'%s': unable to create write destination: %w", name, err)
		}

		enc := json.NewEncoder(dest)
		enc.SetIndent("", "  ")
		if err := enc.Encode(&doc); err != nil {
			return fmt.Errorf("'%s': marshaling openapi3 document: %w", name, err)
		}
	}

	return nil
}

func (c *converter) processSpecComponentRegistryExtension(s spec.Specification) error {

	model := s.Model()

	if types, ok := getComponentRegistryExtension(model.Extensions); ok {
		for _, typ := range types {
			if _, ok := c.outRootDoc.Components.Schemas[typ]; ok {
				continue
			}
			typMapping, err := c.inSpecSet.TypeMapping().Mapping("openapi3", typ)
			if err != nil {
				return fmt.Errorf("invalid model spec '%s': invalid %s extension: no type mapping found for %s: %w", model.ResourceName, componentRegistryExtensionName, typ, err)
			}
			typSchema := new(openapi3.Schema)
			if err := json.Unmarshal([]byte(typMapping.Type), typSchema); err != nil {
				return fmt.Errorf("invalid model spec '%s': invalid %s extension: type mapping unmarshaling failed: %w", model.ResourceName, componentRegistryExtensionName, err)
			}
			c.outRootDoc.Components.Schemas[typ] = typSchema.NewRef()
		}
	}

	if license, ok := getLicenseExtension(model.Extensions); ok {
		c.outRootDoc.Info.License = license
	}

	if servers, ok := getServersExtension(model.Extensions); ok {
		c.outRootDoc.Servers = servers
	}

	if securityRequirements, ok := getSecurityRequirementsExtension(model.Extensions); ok {
		c.outRootDoc.Security = securityRequirements
	}

	if securitySchemes, ok := getSecuritySchemesExtension(model.Extensions); ok {
		c.outRootDoc.Components.SecuritySchemes = securitySchemes
	}

	if openAPIExtensions, ok := getOpenAPIExtensionsExtension(model.Extensions); ok {
		c.outRootDoc.Extensions = openAPIExtensions
	}

	return nil
}

func (c *converter) processSpec(s spec.Specification) error {

	model := s.Model()

	if c.skipPrivateModels && model.Private {
		return nil
	}

	if model.IsRoot {
		pathItems := c.convertRelationsForRootSpec(s.Relations())
		for path, item := range pathItems {
			c.outRootDoc.Paths.Set(path, item)
		}
		// we don't care about root model's relations for now, so we are done for root spec
		return nil
	}

	schema, err := c.convertModel(s)
	if err != nil {
		return fmt.Errorf("model '%s': %w", model.RestName, err)
	}
	c.outRootDoc.Components.Schemas[model.RestName] = schema

	pathItems := c.convertRelationsForNonRootModel(model)
	for path, item := range pathItems {
		c.outRootDoc.Paths.Set(path, item)
	}

	pathItems = c.convertRelationsForNonRootSpec(model.ResourceName, s.Relations())
	for path, item := range pathItems {
		c.outRootDoc.Paths.Set(path, item)
	}

	return nil
}

func (c *converter) convertedDocs() map[string]openapi3.T {

	if !c.splitOutput || len(c.outRootDoc.Components.Schemas) == 0 {
		c.outRootDoc.Tags = c.globalTags()
		return map[string]openapi3.T{defaultDocName: c.outRootDoc}
	}

	docs := make(map[string]openapi3.T)
	specConfig := c.inSpecSet.Configuration()
	for name, schema := range c.outRootDoc.Components.Schemas {
		template := newOpenAPI3Template(specConfig)
		template.Components.Schemas[name] = schema
		template.Info.Title = name
		template.Tags = c.tagsForModel[name]
		docs[name] = template
	}

	for path, item := range c.outRootDoc.Paths.Map() {
		pathRoot := strings.SplitN(strings.Trim(path, "/"), "/", 2)[0]
		docName := c.resourceToRest[pathRoot]
		docs[docName].Paths.Set(path, item)
	}

	return docs
}

func (c *converter) cacheTags(model *spec.Model) {

	if model.IsRoot {
		return
	}
	if c.skipPrivateModels && model.Private {
		return
	}

	tags := openapi3.Tags{
		{
			Name:        model.Group,
			Description: fmt.Sprintf("This tag is for group '%s'", model.Group),
		},
	}

	for _, t := range tags {
		c.globalTagSet[t.Name] = t
	}
	c.tagsForModel[model.RestName] = tags
}

func (c *converter) globalTags() openapi3.Tags {
	tags := make(openapi3.Tags, 0, len(c.globalTagSet))
	for _, t := range c.globalTagSet {
		tags = append(tags, t)
	}
	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Name < tags[j].Name
	})
	return tags
}

func newOpenAPI3Template(specConfig *spec.Config) openapi3.T {
	return openapi3.T{
		OpenAPI: "3.1.0",
		Servers: openapi3.Servers{
			&openapi3.Server{
				URL: "https://api.acuvity.ai",
			},
		},
		Info: &openapi3.Info{
			Title:       specConfig.ProductName,
			Version:     specConfig.Version,
			Description: specConfig.Description,
			Contact: &openapi3.Contact{
				Name:  specConfig.Author,
				URL:   specConfig.URL,
				Email: specConfig.Email,
			},
		},
		Paths: &openapi3.Paths{},
		Components: &openapi3.Components{
			Schemas: make(openapi3.Schemas),
			SecuritySchemes: openapi3.SecuritySchemes{
				"BearerAuth": &openapi3.SecuritySchemeRef{
					Value: &openapi3.SecurityScheme{
						Type:   "http",
						Scheme: "bearer",
					},
				},
				"NamespaceHeader": &openapi3.SecuritySchemeRef{
					Value: &openapi3.SecurityScheme{
						Type: "apiKey",
						Name: "X-Namespace",
						In:   "header",
					},
				},
			},
		},
		Security: openapi3.SecurityRequirements{
			openapi3.SecurityRequirement{
				"BearerAuth":      {},
				"NamespaceHeader": {},
			},
		},
	}
}
