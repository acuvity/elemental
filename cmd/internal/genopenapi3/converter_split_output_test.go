package genopenapi3

import "testing"

func TestConverter_Do__splitOutput_emptyRootModel(t *testing.T) {
	t.Parallel()

	inSpec := `
		model:
			root: true
			rest_name: root
			resource_name: root
			entity_name: Root
			package: root
			group: core
			description: root object.
	`
	outDocs := `
		{
			"openapi": "3.1.0",
			"servers": [{ "url": "https://api.acuvity.ai"}],
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
	`

	testCaseWrapper := map[string]testCase{
		"empty-toplevel-single-output": {
			inSplitOutput: true,
			inSpec:        inSpec,
			outDocs:       map[string]string{"toplevel": outDocs},
		},
	}
	runAllTestCases(t, testCaseWrapper)
}

func TestConverter_Do__split_output_complex(t *testing.T) {
	t.Parallel()

	inSpec := `
		model:
			root: true
			rest_name: root
			resource_name: root
			entity_name: Root
			package: root
			group: core
			description: root object.

		relations:
		- rest_name: minesite
			get:
				description: Retrieves all minesites.
			create:
				description: Creates a new minesite.
	`

	supportingSpecs := []string{
		`
		model:
			rest_name: minesite
			resource_name: minesites
			entity_name: MineSites
			package: usefulPackageName
			group: useful/thing
			description: Represents a resource mine site.
			get:
				description: Retrieves a mine site by ID.
			update:
				description: Updates a mine site by ID.
			delete:
				description: Delete a minesite by ID.

		relations:
		- rest_name: resource
			get:
				description: Retrieves a list of resources for a given mine site.
			create:
				description: assign a new resource for a given mine site.
		`,
		`
		model:
			rest_name: resource
			resource_name: resources
			entity_name: Resources
			package: naturalResources
			group: oil/gas
			description: Represents a natural resource.
		attributes:
			v1:
			- name: supervisor
				description: The supervisor of this natural resource.
				exposed: true
				type: ref
				subtype: employee
		`,
		`
		model:
			rest_name: employee
			resource_name: employees
			entity_name: Employees
			package: people
			group: employee/affairs
			description: Represents a full-time employee.
		`,
	}

	outDocs := map[string]string{
		"minesite": `
			{
				"openapi": "3.1.0",
				"servers": [{ "url": "https://api.acuvity.ai"}],
				"tags":[
					{
						"name": "useful/thing",
						"description": "This tag is for group 'useful/thing'"
					}
				],
				"info": {
					"contact": {
						"email": "dev@aporeto.com",
						"name":  "Aporeto Inc.",
						"url":   "go.acuvity.ai/api"
					},
					"version": "1.0",
					"title": "minesite"
				},
				"components": {
					"schemas": {
						"minesite": {
							"description": "Represents a resource mine site.",
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
				"paths": {
					"/minesites": {
						"get": {
							"description": "Retrieves all minesites.",
							"operationId": "get-all-MineSites",
							"responses": {
								"200": {
									"content": {
										"application/json": {
											"schema": {
												"items": {
													"$ref": "#/components/schemas/minesite"
												},
												"type": "array"
											}
										}
									},
									"description": "n/a"
								}
							},
							"tags": [
								"useful/thing"
							]
						},
						"post": {
							"description": "Creates a new minesite.",
							"operationId": "create-MineSites",
							"requestBody": {
								"content": {
									"application/json": {
										"schema": {
											"$ref": "#/components/schemas/minesite"
										}
									}
								},
								"required": true
							},
							"responses": {
								"200": {
									"content": {
										"application/json": {
											"schema": {
												"$ref": "#/components/schemas/minesite"
											}
										}
									},
									"description": "n/a"
								}
							},
							"tags": [
								"useful/thing"
							]
						}
					},
					"/minesites/{id}": {
						"delete": {
							"description": "Delete a minesite by ID.",
							"operationId": "delete-MineSites",
							"responses": {
								"200": {
									"content": {
										"application/json": {
											"schema": {
												"$ref": "#/components/schemas/minesite"
											}
										}
									},
									"description": "n/a"
								}
							},
							"tags": [
								"useful/thing"
							]
						},
						"get": {
							"description": "Retrieves a mine site by ID.",
							"operationId": "get-MineSites",
							"responses": {
								"200": {
									"content": {
										"application/json": {
											"schema": {
												"$ref": "#/components/schemas/minesite"
											}
										}
									},
									"description": "n/a"
								}
							},
							"tags": [
								"useful/thing"
							]
						},
						"parameters": [
							{
								"in": "path",
								"name": "id",
								"required": true,
								"schema": {
									"type": "string"
								}
							}
						],
						"put": {
							"description": "Updates a mine site by ID.",
							"operationId": "update-MineSites",
							"requestBody": {
								"content": {
									"application/json": {
										"schema": {
											"$ref": "#/components/schemas/minesite"
										}
									}
								},
								"required": true
							},
							"responses": {
								"200": {
									"content": {
										"application/json": {
											"schema": {
												"$ref": "#/components/schemas/minesite"
											}
										}
									},
									"description": "n/a"
								}
							},
							"tags": [
								"useful/thing"
							]
						}
					},
					"/minesites/{id}/resources": {
						"get": {
							"description": "Retrieves a list of resources for a given mine site.",
							"operationId": "get-all-Resources-in-minesite",
							"responses": {
								"200": {
									"content": {
										"application/json": {
											"schema": {
												"items": {
													"$ref": "./resource#/components/schemas/resource"
												},
												"type": "array"
											}
										}
									},
									"description": "n/a"
								}
							},
							"tags": [
								"oil/gas"
							]
						},
						"parameters": [
							{
								"in": "path",
								"name": "id",
								"required": true,
								"schema": {
									"type": "string"
								}
							}
						],
						"post": {
							"description": "assign a new resource for a given mine site.",
							"operationId": "create-Resources-in-minesite",
							"requestBody": {
								"content": {
									"application/json": {
										"schema": {
											"$ref": "./resource#/components/schemas/resource"
										}
									}
								},
								"required": true
							},
							"responses": {
								"200": {
									"content": {
										"application/json": {
											"schema": {
												"$ref": "./resource#/components/schemas/resource"
											}
										}
									},
									"description": "n/a"
								}
							},
							"tags": [
								"oil/gas"
							]
						}
					}
				},
				"security": [
					{
						"BearerAuth": [],
						"NamespaceHeader": []
					}
				]
			}
		`,

		"resource": `
			{
				"servers": [{ "url": "https://api.acuvity.ai"}],
				"openapi": "3.1.0",
				"tags":[
					{
						"name": "oil/gas",
						"description": "This tag is for group 'oil/gas'"
					}
				],
				"info": {
					"contact": {
						"email": "dev@aporeto.com",
						"name":  "Aporeto Inc.",
						"url":   "go.acuvity.ai/api"
					},
					"version": "1.0",
					"title": "resource"
				},
				"components": {
					"schemas": {
						"resource": {
							"description": "Represents a natural resource.",
							"properties": {
								"supervisor": {
									"$ref": "./employee#/components/schemas/employee"
								}
							},
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

		"employee": `
			{
				"servers": [{ "url": "https://api.acuvity.ai"}],
				"openapi": "3.1.0",
				"tags":[
					{
						"name": "employee/affairs",
						"description": "This tag is for group 'employee/affairs'"
					}
				],
				"info": {
					"contact": {
						"email": "dev@aporeto.com",
						"name":  "Aporeto Inc.",
						"url":   "go.acuvity.ai/api"
					},
					"version": "1.0",
					"title": "employee"
				},
				"components": {
					"schemas": {
						"employee": {
							"description": "Represents a full-time employee.",
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
	}

	testCaseWrapper := map[string]testCase{
		"multiple-models-and-relations": {
			inSplitOutput:   true,
			inSpec:          inSpec,
			supportingSpecs: supportingSpecs,
			outDocs:         outDocs,
		},
	}
	runAllTestCases(t, testCaseWrapper)
}

func TestConverter_Do__split_output_withPrivateModel(t *testing.T) {
	t.Parallel()

	inSpec := `
		model:
			root: true
			rest_name: root
			resource_name: root
			entity_name: Root
			package: root
			group: core
			description: root object.

		relations:
		- rest_name: minesite
			get:
				description: Retrieves all minesites.
		- rest_name: hidden
			get:
				description: Retrieves all hidden secrets.
	`

	supportingSpecs := []string{
		`
		model:
			rest_name: minesite
			resource_name: minesites
			entity_name: MineSites
			package: usefulPackageName
			group: useful/thing
			description: Represents a resource mine site.
		`,
		`
		model:
			rest_name: hidden
			resource_name: hiddens
			entity_name: Hiddens
			package: secret
			group: secret/affairs
			description: Represents a private model.
			private: true
		`,
	}

	outDocs := map[string]string{
		"minesite": `
			{
				"servers": [{ "url": "https://api.acuvity.ai"}],
				"openapi": "3.1.0",
				"tags":[
					{
						"name": "useful/thing",
						"description": "This tag is for group 'useful/thing'"
					}
				],
				"info": {
					"contact": {
						"email": "dev@aporeto.com",
						"name":  "Aporeto Inc.",
						"url":   "go.acuvity.ai/api"
					},
					"version": "1.0",
					"title": "minesite"
				},
				"components": {
					"schemas": {
						"minesite": {
							"description": "Represents a resource mine site.",
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
				"paths": {
					"/minesites": {
						"get": {
							"description": "Retrieves all minesites.",
							"operationId": "get-all-MineSites",
							"responses": {
								"200": {
									"content": {
										"application/json": {
											"schema": {
												"items": {
													"$ref": "#/components/schemas/minesite"
												},
												"type": "array"
											}
										}
									},
									"description": "n/a"
								}
							},
							"tags": [
								"useful/thing"
							]
						}
					}
				},
				"security": [
					{
						"BearerAuth": [],
						"NamespaceHeader": []
					}
				]
			}
		`,
	}

	testCaseWrapper := map[string]testCase{
		"root-relation-has-private-model": {
			inSplitOutput:       true,
			inSkipPrivateModels: true,
			inSpec:              inSpec,
			supportingSpecs:     supportingSpecs,
			outDocs:             outDocs,
		},
	}
	runAllTestCases(t, testCaseWrapper)
}
