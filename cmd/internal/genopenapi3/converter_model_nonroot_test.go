package genopenapi3

import (
	"testing"
)

func TestConverter_Do__modelsAndAttributes_nonRoot(t *testing.T) {
	// t.Parallel()

	cases := map[string]testCase{

		"no-attributes": {
			inSpec: `
				model:
					rest_name: void
					resource_name: voids
					entity_name: Void
					package: None
					group: N/A
					description: empty model.
			`,
			outDocs: map[string]string{
				"toplevel": `
					{
						"servers": [{ "url": "https://api.acuvity.ai"}],
						"openapi": "3.1.0",
						"tags": [
							{
								"name": "N/A",
								"description": "This tag is for group 'N/A'"
							}
						],
						"info": {
							"contact": {
								"email": "dev@aporeto.com",
								"name":  "Aporeto Inc.",
								"url":   "go.acuvity.ai/api"
							},
							"version": "1.0",
							"title": "dummy"
						},
						"components": {
							"schemas": {
								"void": {
									"description": "empty model.",
									"type": "object"
								}
							},
							"securitySchemes": {
								"BearerAuth": {
									"scheme": "bearer",
									"type": "http"
								},
								"NamespaceHeader": {
									"in": "header",
									"name": "X-Namespace",
									"type": "apiKey"
								}
							}
						},
						"paths": {},
						"security": [
							{
								"BearerAuth": [],
								"NamespaceHeader": []
							}
						]
					}
				`,
			},
		},

		"attribute-ignored-if-unexposed": {
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
				attributes:
					v1:
					- name: somefield
						description: useful description.
						type: integer
						exposed: false
			`,
			outDocs: map[string]string{
				"toplevel": `
					{
						"servers": [{ "url": "https://api.acuvity.ai"}],
						"openapi": "3.1.0",
						"tags": [
							{
								"name": "N/A",
								"description": "This tag is for group 'N/A'"
							}
						],
						"info": {
							"contact": {
								"email": "dev@aporeto.com",
								"name":  "Aporeto Inc.",
								"url":   "go.acuvity.ai/api"
							},
							"version": "1.0",
							"title": "dummy"
						},
						"components": {
							"schemas": {
								"test": {
									"description": "dummy.",
									"type": "object"
								}
							},
							"securitySchemes": {
								"BearerAuth": {
									"scheme": "bearer",
									"type": "http"
								},
								"NamespaceHeader": {
									"in": "header",
									"name": "X-Namespace",
									"type": "apiKey"
								}
							}
						},
						"paths": {},
						"security": [
							{
								"BearerAuth": [],
								"NamespaceHeader": []
							}
						]
					}
				`,
			},
		},

		"model-is-ignored-if-private-and-skip-flag-is-set": {
			inSkipPrivateModels: true,
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
					private: true
				attributes:
					v1:
					- name: somefield
						description: useful description.
						type: integer
						exposed: true
			`,
			outDocs: map[string]string{
				"toplevel": `
					{
						"servers": [{ "url": "https://api.acuvity.ai"}],
						"openapi": "3.1.0",
						"info": {
							"contact": {
								"email": "dev@aporeto.com",
								"name":  "Aporeto Inc.",
								"url":   "go.acuvity.ai/api"
							},
							"version": "1.0",
							"title": "dummy"
						},
						"components": {
							"securitySchemes": {
								"BearerAuth": {
									"scheme": "bearer",
									"type": "http"
								},
								"NamespaceHeader": {
									"in": "header",
									"name": "X-Namespace",
									"type": "apiKey"
								}
							}
						},
						"paths": {},
						"security": [
							{
								"BearerAuth": [],
								"NamespaceHeader": []
							}
						]
					}
				`,
			},
		},

		"primitive-attributes": {
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
				attributes:
					v1:
					- name: stringField
						description: useful description for string.
						type: string
						exposed: true
					- name: intField
						description: useful description for integer.
						type: integer
						exposed: true
					- name: floatField
						description: useful description for float.
						type: float
						exposed: true
					- name: booleanField
						description: useful description for boolean.
						type: boolean
						exposed: true
					- name: timeField
						description: useful description for time.
						type: time
						exposed: true
			`,
			outDocs: map[string]string{
				"toplevel": `
					{
						"servers": [{ "url": "https://api.acuvity.ai"}],
						"openapi": "3.1.0",
						"tags": [
							{
								"name": "N/A",
								"description": "This tag is for group 'N/A'"
							}
						],
						"info": {
							"contact": {
								"email": "dev@aporeto.com",
								"name":  "Aporeto Inc.",
								"url":   "go.acuvity.ai/api"
							},
							"version": "1.0",
							"title": "dummy"
						},
						"components": {
							"schemas": {
								"test": {
									"description": "dummy.",
									"type": "object",
									"properties": {
										"stringField": {
											"description": "useful description for string.",
											"type": "string"
										},
										"intField": {
											"description": "useful description for integer.",
											"type": "integer"
										},
										"floatField": {
											"description": "useful description for float.",
											"type": "number"
										},
										"booleanField": {
											"description": "useful description for boolean.",
											"type": "boolean"
										},
										"timeField": {
											"description": "useful description for time.",
											"type": "string",
											"format": "date-time"
										}
									}
								}
							},
							"securitySchemes": {
								"BearerAuth": {
									"scheme": "bearer",
									"type": "http"
								},
								"NamespaceHeader": {
									"in": "header",
									"name": "X-Namespace",
									"type": "apiKey"
								}
							}
						},
						"paths": {},
						"security": [
							{
								"BearerAuth": [],
								"NamespaceHeader": []
							}
						]
					}
				`,
			},
		},

		"enum-attribute": {
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
				attributes:
					v1:
					- name: someField
						description: useful description.
						type: enum
						allowed_choices:
							- Choice1
							- Choice2
						exposed: true
			`,
			outDocs: map[string]string{
				"toplevel": `
					{
						"servers": [{ "url": "https://api.acuvity.ai"}],
						"openapi": "3.1.0",
						"tags": [
							{
								"name": "N/A",
								"description": "This tag is for group 'N/A'"
							}
						],
						"info": {
							"contact": {
								"email": "dev@aporeto.com",
								"name":  "Aporeto Inc.",
								"url":   "go.acuvity.ai/api"
							},
							"version": "1.0",
							"title": "dummy"
						},
						"components": {
							"schemas": {
								"test": {
									"description": "dummy.",
									"type": "object",
									"properties": {
										"someField": {
											"description": "useful description.",
											"enum": ["Choice1", "Choice2"]
										}
									}
								}
							},
							"securitySchemes": {
								"BearerAuth": {
									"scheme": "bearer",
									"type": "http"
								},
								"NamespaceHeader": {
									"in": "header",
									"name": "X-Namespace",
									"type": "apiKey"
								}
							}
						},
						"paths": {},
						"security": [
							{
								"BearerAuth": [],
								"NamespaceHeader": []
							}
						]
					}
				`,
			},
		},

		"object-attribute": {
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
				attributes:
					v1:
					- name: someField
						description: useful description.
						type: object
						exposed: true
			`,
			outDocs: map[string]string{
				"toplevel": `
					{
						"servers": [{ "url": "https://api.acuvity.ai"}],
						"openapi": "3.1.0",
						"tags": [
							{
								"name": "N/A",
								"description": "This tag is for group 'N/A'"
							}
						],
						"info": {
							"contact": {
								"email": "dev@aporeto.com",
								"name":  "Aporeto Inc.",
								"url":   "go.acuvity.ai/api"
							},
							"version": "1.0",
							"title": "dummy"
						},
						"components": {
							"schemas": {
								"test": {
									"description": "dummy.",
									"type": "object",
									"properties": {
										"someField": {
											"description": "useful description.",
											"type": "object"
										}
									}
								}
							},
							"securitySchemes": {
								"BearerAuth": {
									"scheme": "bearer",
									"type": "http"
								},
								"NamespaceHeader": {
									"in": "header",
									"name": "X-Namespace",
									"type": "apiKey"
								}
							}
						},
						"paths": {},
						"security": [
							{
								"BearerAuth": [],
								"NamespaceHeader": []
							}
						]
					}
				`,
			},
		},

		"list-of-primitive-attributes": {
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
				attributes:
					v1:
					- name: stringListField
						description: useful stringListField description.
						type: list
						subtype: string
						exposed: true
					- name: integerListField
						description: useful integerListField description.
						type: list
						subtype: integer
						exposed: true
					- name: floatListField
						description: useful floatListField description.
						type: list
						subtype: float
						exposed: true
					- name: booleanListField
						description: useful booleanListField description.
						type: list
						subtype: boolean
						exposed: true
					- name: timeListField
						description: useful timeListField description.
						type: list
						subtype: time
						exposed: true
				`,
			outDocs: map[string]string{
				"toplevel": `
					{
						"servers": [{ "url": "https://api.acuvity.ai"}],
						"openapi": "3.1.0",
						"tags": [
							{
								"name": "N/A",
								"description": "This tag is for group 'N/A'"
							}
						],
						"info": {
							"contact": {
								"email": "dev@aporeto.com",
								"name":  "Aporeto Inc.",
								"url":   "go.acuvity.ai/api"
							},
							"version": "1.0",
							"title": "dummy"
						},
						"components": {
							"schemas": {
								"test": {
									"description": "dummy.",
									"type": "object",
									"properties": {
										"stringListField": {
											"description": "useful stringListField description.",
											"type": "array",
											"items": {
												"type": "string"
											}
										},
										"integerListField": {
											"description": "useful integerListField description.",
											"type": "array",
											"items": {
												"type": "integer"
											}
										},
										"floatListField": {
											"description": "useful floatListField description.",
											"type": "array",
											"items": {
												"type": "number"
											}
										},
										"booleanListField": {
											"description": "useful booleanListField description.",
											"type": "array",
											"items": {
												"type": "boolean"
											}
										},
										"timeListField": {
											"description": "useful timeListField description.",
											"type": "array",
											"items": {
												"type": "string",
												"format": "date-time"
											}
										}
									}
								}
							},
							"securitySchemes": {
								"BearerAuth": {
									"scheme": "bearer",
									"type": "http"
								},
								"NamespaceHeader": {
									"in": "header",
									"name": "X-Namespace",
									"type": "apiKey"
								}
							}
						},
						"paths": {},
						"security": [
							{
								"BearerAuth": [],
								"NamespaceHeader": []
							}
						]
					}
				`,
			},
		},

		// we assume any referenced type is already defined in 'components.schemas'
		"attribute-with-ref-type": {
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
				attributes:
					v1:
					- name: someField
						description: this should be ignored per openapi3 specs.
						type: ref
						subtype: imaginary
						exposed: true
			`,
			outDocs: map[string]string{
				"toplevel": `
					{
						"servers": [{ "url": "https://api.acuvity.ai"}],
						"openapi": "3.1.0",
						"tags": [
							{
								"name": "N/A",
								"description": "This tag is for group 'N/A'"
							}
						],
						"info": {
							"contact": {
								"email": "dev@aporeto.com",
								"name":  "Aporeto Inc.",
								"url":   "go.acuvity.ai/api"
							},
							"version": "1.0",
							"title": "dummy"
						},
						"components": {
							"schemas": {
								"test": {
									"description": "dummy.",
									"type": "object",
									"properties": {
										"someField": {
											"$ref": "#/components/schemas/imaginary"
										}
									}
								}
							},
							"securitySchemes": {
								"BearerAuth": {
									"scheme": "bearer",
									"type": "http"
								},
								"NamespaceHeader": {
									"in": "header",
									"name": "X-Namespace",
									"type": "apiKey"
								}
							}
						},
						"paths": {},
						"security": [
							{
								"BearerAuth": [],
								"NamespaceHeader": []
							}
						]
					}
				`,
			},
		},

		// we assume any referenced type is already defined in 'components.schemas'
		"attributes-with-refList-type": {
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
				attributes:
					v1:
					- name: someField1
						description: useful someField1 description.
						type: refList
						subtype: imaginary1
						exposed: true
					- name: someField2
						description: useful someField2 description.
						type: refList
						subtype: imaginary2
						exposed: true
			`,
			outDocs: map[string]string{
				"toplevel": `
					{
						"servers": [{ "url": "https://api.acuvity.ai"}],
						"openapi": "3.1.0",
						"tags": [
							{
								"name": "N/A",
								"description": "This tag is for group 'N/A'"
							}
						],
						"info": {
							"contact": {
								"email": "dev@aporeto.com",
								"name":  "Aporeto Inc.",
								"url":   "go.acuvity.ai/api"
							},
							"version": "1.0",
							"title": "dummy"
						},
						"components": {
							"schemas": {
								"test": {
									"description": "dummy.",
									"type": "object",
									"properties": {
										"someField1": {
											"description": "useful someField1 description.",
											"type": "array",
											"items": {
												"$ref": "#/components/schemas/imaginary1"
											}
										},
										"someField2": {
											"description": "useful someField2 description.",
											"type": "array",
											"items": {
												"$ref": "#/components/schemas/imaginary2"
											}
										}
									}
								}
							},
							"securitySchemes": {
								"BearerAuth": {
									"scheme": "bearer",
									"type": "http"
								},
								"NamespaceHeader": {
									"in": "header",
									"name": "X-Namespace",
									"type": "apiKey"
								}
							}
						},
						"paths": {},
						"security": [
							{
								"BearerAuth": [],
								"NamespaceHeader": []
							}
						]
					}
				`,
			},
		},

		// we assume any referenced type is already defined in 'components.schemas'
		"attributes-with-refMap-type": {
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
				attributes:
					v1:
					- name: someField1
						description: useful someField1 description.
						type: refMap
						subtype: imaginary1
						exposed: true
					- name: someField2
						description: useful someField2 description.
						type: refMap
						subtype: imaginary2
						exposed: true
			`,
			outDocs: map[string]string{
				"toplevel": `
					{
						"servers": [{ "url": "https://api.acuvity.ai"}],
						"openapi": "3.1.0",
						"tags": [
							{
								"name": "N/A",
								"description": "This tag is for group 'N/A'"
							}
						],
						"info": {
							"contact": {
								"email": "dev@aporeto.com",
								"name":  "Aporeto Inc.",
								"url":   "go.acuvity.ai/api"
							},
							"version": "1.0",
							"title": "dummy"
						},
						"components": {
							"schemas": {
								"test": {
									"description": "dummy.",
									"type": "object",
									"properties": {
										"someField1": {
											"description": "useful someField1 description.",
											"type": "object",
											"additionalProperties": {
												"$ref": "#/components/schemas/imaginary1"
											}
										},
										"someField2": {
											"description": "useful someField2 description.",
											"type": "object",
											"additionalProperties": {
												"$ref": "#/components/schemas/imaginary2"
											}
										}
									}
								}
							},
							"securitySchemes": {
								"BearerAuth": {
									"scheme": "bearer",
									"type": "http"
								},
								"NamespaceHeader": {
									"in": "header",
									"name": "X-Namespace",
									"type": "apiKey"
								}
							}
						},
						"paths": {},
						"security": [
							{
								"BearerAuth": [],
								"NamespaceHeader": []
							}
						]
					}
				`,
			},
		},

		"attributes-with-external-type--[]byte-turns-into-string": {
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
				attributes:
					v1:
					- name: someField
						description: useful description.
						type: external
						subtype: '[]byte'
						exposed: true
			`,
			outDocs: map[string]string{
				"toplevel": `
					{
						"servers": [{ "url": "https://api.acuvity.ai"}],
						"openapi": "3.1.0",
						"tags": [
							{
								"name": "N/A",
								"description": "This tag is for group 'N/A'"
							}
						],
						"info": {
							"contact": {
								"email": "dev@aporeto.com",
								"name":  "Aporeto Inc.",
								"url":   "go.acuvity.ai/api"
							},
							"version": "1.0",
							"title": "dummy"
						},
						"components": {
							"schemas": {
								"test": {
									"description": "dummy.",
									"type": "object",
									"properties": {
										"someField": {
											"description": "useful description.",
											"type": "string"
										}
									}
								}
							},
							"securitySchemes": {
								"BearerAuth": {
									"scheme": "bearer",
									"type": "http"
								},
								"NamespaceHeader": {
									"in": "header",
									"name": "X-Namespace",
									"type": "apiKey"
								}
							}
						},
						"paths": {},
						"security": [
							{
								"BearerAuth": [],
								"NamespaceHeader": []
							}
						]
					}
				`,
			},
		},

		"required-attributes": {
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
				attributes:
					v1:
					- name: stringField
						description: useful description for string.
						type: string
						exposed: true
						required: true
						default_value: hello-world
					- name: intField
						description: useful description for integer.
						type: integer
						exposed: true
			`,
			outDocs: map[string]string{
				"toplevel": `
					{
						"servers": [{ "url": "https://api.acuvity.ai"}],
						"openapi": "3.1.0",
						"tags": [
							{
								"name": "N/A",
								"description": "This tag is for group 'N/A'"
							}
						],
						"info": {
							"contact": {
								"email": "dev@aporeto.com",
								"name":  "Aporeto Inc.",
								"url":   "go.acuvity.ai/api"
							},
							"version": "1.0",
							"title": "dummy"
						},
						"components": {
							"schemas": {
								"test": {
									"description": "dummy.",
									"type": "object",
									"required": ["stringField"],
									"properties": {
										"stringField": {
											"description": "useful description for string.",
											"default": "hello-world",
											"type": "string"
										},
										"intField": {
											"description": "useful description for integer.",
											"type": "integer"
										}
									}
								}
							},
							"securitySchemes": {
								"BearerAuth": {
									"scheme": "bearer",
									"type": "http"
								},
								"NamespaceHeader": {
									"in": "header",
									"name": "X-Namespace",
									"type": "apiKey"
								}
							}
						},
						"paths": {},
						"security": [
							{
								"BearerAuth": [],
								"NamespaceHeader": []
							}
						]
					}
				`,
			},
		},
	}
	runAllTestCases(t, cases)
}
