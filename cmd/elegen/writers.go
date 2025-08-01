// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"go/scanner"
	"path"
	"strings"
	"text/template"

	"go.acuvity.ai/regolithe/spec"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/tools/imports"
)

var functions = template.FuncMap{
	"upper":                           strings.ToUpper,
	"lower":                           strings.ToLower,
	"capitalize":                      cases.Title(language.Und, cases.NoLower).String,
	"join":                            strings.Join,
	"hasPrefix":                       strings.HasPrefix,
	"attrBSONFieldName":               attrBSONFieldName,
	"attrToType":                      attrToType,
	"attrToField":                     attrToField,
	"attrToMongoField":                attrToMongoField,
	"escBackticks":                    escapeBackticks,
	"buildEnums":                      buildEnums,
	"shouldGenerateGetter":            shouldGenerateGetter,
	"shouldGenerateSetter":            shouldGenerateSetter,
	"shouldWriteInitializer":          shouldWriteInitializer,
	"shouldWriteAttributeMap":         shouldWriteAttributeMap,
	"shouldRegisterSpecification":     shouldRegisterSpecification,
	"shouldRegisterRelationship":      shouldRegisterRelationship,
	"shouldRegisterInnerRelationship": shouldRegisterInnerRelationship,
	"writeInitializer":                writeInitializer,
	"writeDefaultValue":               writeDefaultValue,
	"sortAttributes":                  sortAttributes,
	"sortIndexes":                     sortIndexes,
	"modelCommentFlags":               modelCommentFlags,
}

func writeModel(set spec.SpecificationSet, name string, outFolder string, publicMode bool) error {

	tmpl, err := makeTemplate("templates/model.gotpl")
	if err != nil {
		return err
	}

	s := set.Specification(name)

	bnames := map[string]struct{}{}
	for _, attr := range s.Attributes(s.LatestAttributesVersion()) {
		item, ok := attr.Extensions["bson_name"]
		if !ok {
			continue
		}
		bname := item.(string)
		if _, ok = bnames[bname]; ok {
			return fmt.Errorf("invalid bson name. '%s' reused", bname)
		}
		bnames[bname] = struct{}{}
	}

	if (s.Model().Private || s.Model().Undocumented) && publicMode {
		return nil
	}

	var buf bytes.Buffer

	if err = tmpl.Execute(
		&buf,
		struct {
			Set        spec.SpecificationSet
			Spec       spec.Specification
			PublicMode bool
		}{
			PublicMode: publicMode,
			Set:        set,
			Spec:       s,
		}); err != nil {
		return fmt.Errorf("unable to generate model '%s': %w", name, err)
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		var errs scanner.ErrorList
		if errors.As(err, &errs) {
			lines := strings.Split(buf.String(), "\n")
			for i := range errs.Len() {
				fmt.Printf("Error in '%s' near:\n\n\t%s\n\n", name, lines[errs[i].Pos.Line-1])
			}
		}
		return fmt.Errorf("unable to format model '%s': %w", name, err)
	}

	p, err = imports.Process("", p, nil)
	if err != nil {
		return err
	}

	if err := writeFile(path.Join(outFolder, name+".go"), p); err != nil {
		return fmt.Errorf("unable to write file for spec: %s", name)
	}

	return nil
}

func writeIdentitiesRegistry(set spec.SpecificationSet, outFolder string, publicMode bool) error {

	tmpl, err := makeTemplate("templates/identities_registry.gotpl")
	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err = tmpl.Execute(
		&buf,
		struct {
			Set        spec.SpecificationSet
			PublicMode bool
		}{
			PublicMode: publicMode,
			Set:        set,
		}); err != nil {
		return fmt.Errorf("unable to generate identities_registry code: %w", err)
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("unable to format identities_registry code: %w", err)
	}

	p, err = imports.Process("", p, nil)
	if err != nil {
		fmt.Println(buf.String())
		return fmt.Errorf("unable to goimport relationships_registry code: %w", err)
	}

	if err := writeFile(path.Join(outFolder, "identities_registry.go"), p); err != nil {
		return fmt.Errorf("unable to write file for identities_registry: %w", err)
	}

	return nil
}

func writeRelationshipsRegistry(set spec.SpecificationSet, outFolder string, publicMode bool) error {

	tmpl, err := makeTemplate("templates/relationships_registry.gotpl")
	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err = tmpl.Execute(
		&buf,
		struct {
			Set        spec.SpecificationSet
			PublicMode bool
		}{
			PublicMode: publicMode,
			Set:        set,
		}); err != nil {
		return fmt.Errorf("unable to generate relationships_registry code: %w", err)
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(buf.String())
		return fmt.Errorf("unable to format relationships_registry code: %w", err)
	}

	p, err = imports.Process("", p, nil)
	if err != nil {
		fmt.Println(buf.String())
		return fmt.Errorf("unable to goimport relationships_registry code: %w", err)
	}

	if err := writeFile(path.Join(outFolder, "relationships_registry.go"), p); err != nil {
		return fmt.Errorf("unable to write file for relationships_registry: %w", err)
	}

	return nil
}
