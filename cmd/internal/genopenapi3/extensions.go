package genopenapi3

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

const (
	licenseExtensionName           = "openapi_license"
	serversExtensionName           = "openapi_servers"
	pathExtensionName              = "openapi_path"
	responseMapExtensionName       = "openapi_response_map"
	componentRegistryExtensionName = "openapi_component_registry"
)

type responseSpec struct {
	Spec        string
	IsArray     bool
	Description string
}

func getLicenseExtension(extensions map[string]any) (*openapi3.License, bool) {

	if extensions == nil {
		return nil, false
	}

	entry, ok := extensions[licenseExtensionName]
	if !ok {
		return nil, false
	}

	licenseMap, ok := entry.(map[string]any)
	if !ok {
		panic(fmt.Sprintf("invalid %s extension: expected type map[string]any, got %T", licenseExtensionName, entry))
	}

	license := &openapi3.License{}
	if name, ok := licenseMap["name"].(string); ok {
		license.Name = name
	}
	if url, ok := licenseMap["url"].(string); ok {
		license.URL = url
	}

	return license, true
}

func getServersExtension(extensions map[string]any) ([]*openapi3.Server, bool) {

	if extensions == nil {
		return nil, false
	}

	entry, ok := extensions[serversExtensionName]
	if !ok {
		return nil, false
	}

	serversUntyped, ok := entry.([]any)
	if !ok {
		panic(fmt.Sprintf("invalid %s extension: expected type []any, got %T", serversExtensionName, entry))
	}

	servers := make([]*openapi3.Server, len(serversUntyped))
	for i, serverUntyped := range serversUntyped {
		serverMap, ok := serverUntyped.(map[string]any)
		if !ok {
			panic(fmt.Sprintf("invalid %s extension entry: expected type map[string]any, got %T", serversExtensionName, serverUntyped))
		}

		server := &openapi3.Server{}
		if url, ok := serverMap["url"].(string); ok {
			server.URL = url
		}
		if description, ok := serverMap["description"].(string); ok {
			server.Description = description
		}
		if variablesUntyped, ok := serverMap["variables"].(map[string]any); ok {
			variables := make(map[string]*openapi3.ServerVariable, len(variablesUntyped))
			for variableName, variableUntyped := range variablesUntyped {
				variableMap, ok := variableUntyped.(map[string]any)
				if !ok {
					panic(fmt.Sprintf("invalid %s extension entry: server variable: expected type map[string]any, got %T", serversExtensionName, variableUntyped))
				}

				variable := &openapi3.ServerVariable{}
				if defaultValue, ok := variableMap["default"].(string); ok {
					variable.Default = defaultValue
				}
				if description, ok := variableMap["description"].(string); ok {
					variable.Description = description
				}
				if enumUntyped, ok := variableMap["enum"].([]any); ok {
					enum := make([]string, len(enumUntyped))
					for i, enumEntry := range enumUntyped {
						enum[i] = enumEntry.(string)
					}
					variable.Enum = enum
				}
				variables[variableName] = variable
			}
			server.Variables = variables
		}

		servers[i] = server
	}

	return servers, true
}

func getPathExtension(extensions map[string]any) (string, bool) {

	if extensions == nil {
		return "", false
	}

	entry, ok := extensions[pathExtensionName]
	if !ok {
		return "", false
	}

	path, ok := entry.(string)
	if !ok {
		panic(fmt.Sprintf("invalid %s extension: expected type string, got %T", pathExtensionName, path))
	}
	return path, true
}

func getResponseMapExtension(extensions map[string]any) (map[string]responseSpec, bool) {

	if extensions == nil {
		return nil, false
	}

	entry, ok := extensions[responseMapExtensionName]
	if !ok {
		return nil, false
	}

	responseMap, ok := entry.(map[string]any)
	if !ok {
		panic(fmt.Sprintf("invalid %s extension: status code: expected type map[string]any, got %T", responseMapExtensionName, entry))
	}
	ret := make(map[string]responseSpec, len(responseMap))
	for statusCode, responseEntry := range responseMap {
		responseEntryMap, ok := responseEntry.(map[string]any)
		if !ok {
			panic(fmt.Sprintf("invalid %s extension entry: response spec: expected type map[string]any, got %T", responseMapExtensionName, responseEntry))
		}
		specEntry, ok := responseEntryMap["spec"]
		var spec string
		if ok {
			spec, ok = specEntry.(string)
			if !ok {
				panic(fmt.Sprintf("invalid %s extension entry: response spec: 'spec' key: expected type string, got %T", responseMapExtensionName, specEntry))
			}
		}
		description := ""
		if descEntry, ok := responseEntryMap["description"]; ok {
			desc, ok := descEntry.(string)
			if !ok {
				panic(fmt.Sprintf("invalid %s extension entry: response spec: 'description' key: expected type string, got %T", responseMapExtensionName, descEntry))
			}
			description = desc
		}
		isArray := false
		if isArrayEntry, ok := responseEntryMap["is_array"]; ok {
			isArray, ok = isArrayEntry.(bool)
			if !ok {
				panic(fmt.Sprintf("invalid %s extension entry: response spec: 'is_array' key: expected type bool, got %T", responseMapExtensionName, isArrayEntry))
			}
		}
		ret[statusCode] = responseSpec{
			Spec:        spec,
			IsArray:     isArray,
			Description: description,
		}
	}
	return ret, true
}

func getComponentRegistryExtension(extensions map[string]any) ([]string, bool) {

	if extensions == nil {
		return nil, false
	}

	entry, ok := extensions[componentRegistryExtensionName]
	if !ok {
		return nil, false
	}

	typesUntyped, ok := entry.([]any)
	if !ok {
		panic(fmt.Sprintf("invalid %s extension: expected type []any, got %T", componentRegistryExtensionName, entry))
	}
	types := make([]string, len(typesUntyped))
	for i, typ := range typesUntyped {
		typStr, ok := typ.(string)
		if !ok {
			panic(fmt.Sprintf("invalid %s extension entry: expected type string, got %T", componentRegistryExtensionName, typ))
		}
		types[i] = typStr
	}

	return types, true
}
