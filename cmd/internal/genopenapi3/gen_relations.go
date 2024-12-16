package genopenapi3

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/getkin/kin-openapi/openapi3"
	"go.acuvity.ai/regolithe/spec"
)

var noDesc = "n/a"

func (c *converter) convertRelationsForRootSpec(relations []*spec.Relation) map[string]*openapi3.PathItem {

	paths := make(map[string]*openapi3.PathItem)

	for _, relation := range relations {

		model := relation.Specification().Model()

		if c.skipPrivateModels && model.Private {
			continue
		}

		if relation.Get == nil && relation.Create == nil {
			continue
		}

		pathItem := &openapi3.PathItem{
			Get:  c.extractOperationGetAll("", relation),
			Post: c.extractOperationPost("", relation),
		}

		uri := "/" + model.ResourceName
		// check if the openapi_path_extension is set, and we use it if it is
		if path, ok := getPathExtension(relation.Extensions); ok {
			uri = "/" + path
		}
		paths[uri] = pathItem
	}

	return paths
}

func (c *converter) convertRelationsForNonRootSpec(resourceName string, relations []*spec.Relation) map[string]*openapi3.PathItem {

	paths := make(map[string]*openapi3.PathItem)
	parentRestName := c.resourceToRest[resourceName]

	for _, relation := range relations {

		if relation.Get == nil && relation.Create == nil {
			continue
		}

		pathItem := &openapi3.PathItem{
			Get:  c.extractOperationGetAll(parentRestName, relation),
			Post: c.extractOperationPost(parentRestName, relation),
		}

		c.insertParamID(&pathItem.Parameters)

		childModel := c.inSpecSet.Specification(relation.RestName).Model()
		uri := fmt.Sprintf("/%s/{%s}/%s", resourceName, paramNameID, childModel.ResourceName)
		// check if the openapi_path_extension is set, and we use it if it is
		if path, ok := getPathExtension(relation.Extensions); ok {
			uri = fmt.Sprintf("/%s/{%s}/%s", resourceName, paramNameID, path)
		}
		paths[uri] = pathItem
	}

	return paths
}

func (c *converter) convertRelationsForNonRootModel(model *spec.Model) map[string]*openapi3.PathItem {

	if model.Get == nil && model.Update == nil && model.Delete == nil {
		return nil
	}

	pathItem := &openapi3.PathItem{
		Get:    c.extractOperationGetByID(model),
		Delete: c.extractOperationDeleteByID(model),
		Put:    c.extractOperationPutByID(model),
	}
	c.insertParamID(&pathItem.Parameters)

	uri := fmt.Sprintf("/%s/{%s}", model.ResourceName, paramNameID)
	if path, ok := getPathExtension(model.Extensions); ok {
		uri = "/" + path
	}
	pathItems := map[string]*openapi3.PathItem{uri: pathItem}
	return pathItems
}

func (c *converter) extractOperationGetAll(parentRestName string, relation *spec.Relation) *openapi3.Operation {

	if relation == nil || relation.Get == nil {
		return nil
	}
	relationAction := relation.Get

	model := relation.Specification().Model()

	respBodySchemaRef := c.getSchemaRef(parentRestName, model.RestName, true)

	resp200 := openapi3.WithStatus(200,
		&openapi3.ResponseRef{
			Value: &openapi3.Response{
				Description: &noDesc,
				Content: openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: respBodySchemaRef,
					},
				},
			},
		},
	)

	op := &openapi3.Operation{
		OperationID: "get-all-" + model.EntityNamePlural,
		Tags:        []string{model.Group},
		Description: relationAction.Description,
		Responses:   openapi3.NewResponses(resp200),
		Parameters:  c.convertParamDefAsQueryParams(relationAction.ParameterDefinition),
	}

	if parentRestName != "" {
		op.OperationID = "get-all-" + model.EntityNamePlural + "-in-" + parentRestName
	}

	extensionResponses, newOperationID := c.createResponsesFromExtension(relationAction, resp200, parentRestName, op.OperationID)
	if extensionResponses != nil {
		op.Responses = extensionResponses
	}
	if newOperationID != "" {
		op.OperationID = newOperationID
	}

	return op
}

func (c *converter) extractOperationPost(parentRestName string, relation *spec.Relation) *openapi3.Operation {

	if relation == nil || relation.Create == nil {
		return nil
	}
	relationAction := relation.Create

	model := relation.Specification().Model()

	schemaRef := c.getSchemaRef(parentRestName, relation.RestName, false)

	resp200 := openapi3.WithStatus(200,
		&openapi3.ResponseRef{
			Value: &openapi3.Response{
				Description: &noDesc,
				Content: openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: schemaRef,
					},
				},
			},
		},
	)

	op := &openapi3.Operation{
		OperationID: "create-" + model.EntityName,
		Tags:        []string{model.Group},
		Description: relationAction.Description,
		RequestBody: &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Required: true,
				Content: openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: schemaRef,
					},
				},
			},
		},
		Responses:  openapi3.NewResponses(resp200),
		Parameters: c.convertParamDefAsQueryParams(relationAction.ParameterDefinition),
	}

	if parentRestName != "" {
		op.OperationID = "create-" + model.EntityName + "-in-" + parentRestName
	}

	extensionResponses, newOperationID := c.createResponsesFromExtension(relationAction, resp200, parentRestName, op.OperationID)
	if extensionResponses != nil {
		op.Responses = extensionResponses
	}
	if newOperationID != "" {
		op.OperationID = newOperationID
	}

	return op
}

func (c *converter) extractOperationGetByID(model *spec.Model) *openapi3.Operation {

	if model == nil || model.Get == nil {
		return nil
	}
	relationAction := model.Get

	respBodySchemaRef := c.getSchemaRef("", model.RestName, false)

	resp200 := openapi3.WithStatus(200,
		&openapi3.ResponseRef{
			Value: &openapi3.Response{
				Description: &noDesc,
				Content: openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: respBodySchemaRef,
					},
				},
			},
		},
	)

	op := &openapi3.Operation{
		OperationID: fmt.Sprintf("get-%s", model.EntityName),
		Tags:        []string{model.Group},
		Description: relationAction.Description,
		Responses:   openapi3.NewResponses(resp200),
		Parameters:  c.convertParamDefAsQueryParams(relationAction.ParameterDefinition),
	}

	extensionResponses, newOperationID := c.createResponsesFromExtension(relationAction, resp200, "", op.OperationID)
	if extensionResponses != nil {
		op.Responses = extensionResponses
	}
	if newOperationID != "" {
		op.OperationID = newOperationID
	}

	return op
}

func (c *converter) extractOperationDeleteByID(model *spec.Model) *openapi3.Operation {

	if model == nil || model.Delete == nil {
		return nil
	}
	relationAction := model.Delete

	respBodySchemaRef := c.getSchemaRef("", model.RestName, false)

	resp200 := openapi3.WithStatus(200,
		&openapi3.ResponseRef{
			Value: &openapi3.Response{
				Description: &noDesc,
				Content: openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: respBodySchemaRef,
					},
				},
			},
		},
	)

	op := &openapi3.Operation{
		OperationID: fmt.Sprintf("delete-%s", model.EntityName),
		Tags:        []string{model.Group},
		Description: relationAction.Description,
		Responses:   openapi3.NewResponses(resp200),
		Parameters:  c.convertParamDefAsQueryParams(relationAction.ParameterDefinition),
	}

	extensionResponses, newOperationID := c.createResponsesFromExtension(relationAction, resp200, "", op.OperationID)
	if extensionResponses != nil {
		op.Responses = extensionResponses
	}
	if newOperationID != "" {
		op.OperationID = newOperationID
	}

	return op
}

func (c *converter) extractOperationPutByID(model *spec.Model) *openapi3.Operation {

	if model == nil || model.Update == nil {
		return nil
	}
	relationAction := model.Update

	schemaRef := c.getSchemaRef("", model.RestName, false)

	resp200 := openapi3.WithStatus(200,
		&openapi3.ResponseRef{
			Value: &openapi3.Response{
				Description: &noDesc,
				Content: openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: schemaRef,
					},
				},
			},
		},
	)

	op := &openapi3.Operation{
		OperationID: fmt.Sprintf("update-%s", model.EntityName),
		Tags:        []string{model.Group},
		Description: relationAction.Description,
		RequestBody: &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Required: true,
				Content: openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: schemaRef,
					},
				},
			},
		},
		Responses:  openapi3.NewResponses(resp200),
		Parameters: c.convertParamDefAsQueryParams(relationAction.ParameterDefinition),
	}

	extensionResponses, newOperationID := c.createResponsesFromExtension(relationAction, resp200, "", op.OperationID)
	if extensionResponses != nil {
		op.Responses = extensionResponses
	}
	if newOperationID != "" {
		op.OperationID = newOperationID
	}

	return op
}

func (c *converter) getSchemaRef(parentRestName, restName string, isArray bool) *openapi3.SchemaRef {

	if isArray {
		arrSchema := openapi3.NewArraySchema()
		if !c.splitOutput || parentRestName == "" {
			arrSchema.Items = openapi3.NewSchemaRef("#/components/schemas/"+restName, nil)
		} else {
			arrSchema.Items = openapi3.NewSchemaRef("./"+restName+"#/components/schemas/"+restName, nil)
		}
		return arrSchema.NewRef()
	}

	if !c.splitOutput || parentRestName == "" {
		return openapi3.NewSchemaRef("#/components/schemas/"+restName, nil)
	}
	return openapi3.NewSchemaRef("./"+restName+"#/components/schemas/"+restName, nil)

}

// createResponsesFromExtension creates a responses object from the openapi_response_map extension
// It returns the new responses if the extension was found and nil otherwise.
// The returned string refers to the new operation ID if it needs updating because of the newly returned type from the 200 response.
func (c *converter) createResponsesFromExtension(relationAction *spec.RelationAction, response200 openapi3.NewResponsesOption, parentRestName, operationID string) (*openapi3.Responses, string) {
	responseMap, ok := getResponseMapExtension(relationAction.Extensions)
	if !ok {
		return nil, ""
	}

	var modelOutEntityName string
	respOpts := make([]openapi3.NewResponsesOption, 0, len(responseMap))
	if responseSpec200, ok := responseMap["200"]; ok {
		responseSpec200Spec := c.inSpecSet.Specification(responseSpec200.Spec)
		if responseSpec200Spec != nil {
			modelOut := responseSpec200Spec.Model()
			if modelOut != nil {
				modelOutEntityName = modelOut.EntityName
			}
		}
	} else {
		respOpts = append(respOpts, response200)
	}
	for statusCodeStr, responseSpec := range responseMap {
		statusCode, err := strconv.Atoi(statusCodeStr)
		if err != nil {
			panic(fmt.Sprintf("invalid status code string '%s': %s", statusCodeStr, err))
		}

		var resp openapi3.NewResponsesOption

		if responseSpec.Spec == "" {
			// if the spec is empty, we assume that this is a text/plain type of response
			// as we don't have a type
			resp = openapi3.WithStatus(statusCode,
				&openapi3.ResponseRef{
					Value: &openapi3.Response{
						Description: stringPtr(responseSpec.Description),
						Content: openapi3.Content{
							"text/plain": &openapi3.MediaType{
								Schema: openapi3.NewStringSchema().NewRef(),
							},
						},
					},
				},
			)
		} else {
			var respSchemaRef *openapi3.SchemaRef
			responseSpecSpec := c.inSpecSet.Specification(responseSpec.Spec)
			if responseSpecSpec != nil {
				// if there is a spec, then there will be a component and we simply use it
				responseModel := responseSpecSpec.Model()
				if responseModel == nil {
					panic(fmt.Sprintf("invalid response spec '%s': no model", responseSpec.Spec))
				}
				respSchemaRef = c.getSchemaRef(parentRestName, responseModel.RestName, responseSpec.IsArray)
			} else if _, ok = c.outRootDoc.Components.Schemas[responseSpec.Spec]; ok {
				// otherwise there might have been a component with this spec registered through the openapi_component_registry extension
				// we create a new reference for that in this case
				respSchemaRef = openapi3.NewSchemaRef("#/components/schemas/"+responseSpec.Spec, nil)
			} else {
				// or otherwise this could be referring to a registered type mapping
				// and as it was not embedded in the components, we create a new schema for it and embed it
				mapping, err := c.inSpecSet.TypeMapping().Mapping("openapi3", responseSpec.Spec)
				if err != nil {
					panic(fmt.Sprintf("invalid response spec '%s': no spec found, and no type mapping found: %s", responseSpec.Spec, err))
				}
				attrSchema := new(openapi3.Schema)
				if err := json.Unmarshal([]byte(mapping.Type), attrSchema); err != nil {
					panic(fmt.Sprintf("invalid response spec '%s': no spec found, and type mapping unmarshaling failed: %s", responseSpec.Spec, err))
				}
				respSchemaRef = attrSchema.NewRef()
			}

			resp = openapi3.WithStatus(statusCode,
				&openapi3.ResponseRef{
					Value: &openapi3.Response{
						Description: stringPtr(responseSpec.Description),
						Content: openapi3.Content{
							"application/json": &openapi3.MediaType{
								Schema: respSchemaRef,
							},
						},
					},
				},
			)
		}

		respOpts = append(respOpts, resp)
	}

	var retOperationID string
	if modelOutEntityName != "" {
		retOperationID = operationID + "-as-" + modelOutEntityName
	}

	return openapi3.NewResponses(respOpts...), retOperationID
}

// stringPtr creates a dangling string pointer which is perfect if you are iterating over something
// and need to refer to the pointer of a loop variable
func stringPtr(str string) *string { return &str }
