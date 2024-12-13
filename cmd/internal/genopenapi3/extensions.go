package genopenapi3

import "fmt"

const (
	pathExtensionName              = "openapi_path"
	responseMapExtensionName       = "openapi_response_map"
	componentRegistryExtensionName = "openapi_component_registry"
)

type responseSpec struct {
	Spec        string
	IsArray     bool
	Description string
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
