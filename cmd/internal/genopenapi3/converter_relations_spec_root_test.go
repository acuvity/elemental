package genopenapi3

import (
	"testing"
)

func TestConverter_Do__specRelations_root(t *testing.T) {
	t.Parallel()

	cases := map[string]testCase{

		"relation-create": {
			inSpec: `
				model:
					root: true
					rest_name: root
					resource_name: root
					entity_name: Root
					package: root
					group: core
					description: root object.

				relations:
				- rest_name: resource
					create:
						description: Creates some resource.
						parameters:
							entries:
							- name: fancyParam
								description: This is a fancy parameter.
								type: integer
			`,
			outDocs: map[string]string{
				"toplevel": `
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
							"title": "dummy"
						},
						"components": {
							"schemas": {
								"resource": {
									"description": "Represents a resource.",
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
							"/resources": {
								"post": {
									"operationId": "create-Resource",
									"tags": ["useful/thing"],
									"parameters": [
										{
											"description": "This is a fancy parameter.",
											"in": "query",
											"name": "fancyParam",
											"schema": {
												"type": "integer"
											}
										}
									],
									"description": "Creates some resource.",
									"requestBody": {
										"content": {
											"application/json": {
												"schema": {
													"$ref": "#/components/schemas/resource"
												}
											}
										},
										"required": true
									},
									"responses": {
										"200": {
											"description": "n/a",
											"content": {
												"application/json": {
													"schema": {
														"$ref": "#/components/schemas/resource"
													}
												}
											}
										}
									}
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
			},
			supportingSpecs: []string{`
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Resource
					package: usefulPackageName
					group: useful/thing
					description: Represents a resource.
			`},
		},

		"relation-get": {
			inSpec: `
				model:
					root: true
					rest_name: root
					resource_name: root
					entity_name: Root
					package: root
					group: core
					description: root object.

				relations:
				- rest_name: resource
					get:
						description: Retrieve all resources.
						parameters:
						  entries:
						  - name: fancyParam
						    description: This is a fancy parameter.
						    type: boolean
			`,
			outDocs: map[string]string{
				"toplevel": `
					{
						"openapi": "3.1.0",
						"servers": [{ "url": "https://api.acuvity.ai"}],
						"tags": [
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
							"title": "dummy"
						},
						"components": {
							"schemas": {
								"resource": {
									"description": "Represents a resource.",
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
							"/resources": {
								"get": {
									"operationId": "get-all-Resources",
									"tags": ["useful/thing"],
									"description": "Retrieve all resources.",
									"parameters": [
									  {
									    "description": "This is a fancy parameter.",
									    "in": "query",
									    "name": "fancyParam",
									    "schema": {
									      "type": "boolean"
									    }
									  }
									],
									"responses": {
										"200": {
											"description": "n/a",
											"content": {
												"application/json": {
													"schema": {
														"type": "array",
														"items": {
															"$ref": "#/components/schemas/resource"
														}
													}
												}
											}
										}
									}
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
			},
			supportingSpecs: []string{`
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Resource
					package: usefulPackageName
					group: useful/thing
					description: Represents a resource.
			`},
		},

		"relation-without-get-or-create": {
			inSpec: `
				model:
					root: true
					rest_name: root
					resource_name: root
					entity_name: Root
					package: root
					group: core
					description: root object.

				relations:
				- rest_name: resource
			`,
			outDocs: map[string]string{
				"toplevel": `
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
							"title": "dummy"
						},
						"components": {
							"schemas": {
								"resource": {
									"description": "Represents a resource.",
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
			supportingSpecs: []string{`
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Resource
					package: usefulPackageName
					group: useful/thing
					description: Represents a resource.
			`},
		},
	}
	runAllTestCases(t, cases)
}

func TestConverter_Do__specRelations_root_withPrivateModel(t *testing.T) {
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
		- rest_name: resource
			create:
				description: Creates some resource.
		- rest_name: hidden
			create:
				description: Creates some hidden secrets.
	`

	outDoc := map[string]string{
		"toplevel": `
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
					"title": "dummy"
				},
				"components": {
					"schemas": {
						"resource": {
							"description": "Represents a resource.",
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
					"/resources": {
						"post": {
							"operationId": "create-Resource",
							"tags": ["useful/thing"],
							"description": "Creates some resource.",
							"requestBody": {
								"content": {
									"application/json": {
										"schema": {
											"$ref": "#/components/schemas/resource"
										}
									}
								},
								"required": true
							},
							"responses": {
								"200": {
									"description": "n/a",
									"content": {
										"application/json": {
											"schema": {
												"$ref": "#/components/schemas/resource"
											}
										}
									}
								}
							}
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

	supportingSpecs := []string{
		`
		model:
			rest_name: resource
			resource_name: resources
			entity_name: Resource
			package: usefulPackageName
			group: useful/thing
			description: Represents a resource.
		`,
		`
			model:
				rest_name: hidden
				resource_name: hiddens
				entity_name: Hidden
				package: secrets
				group: gossip/talk
				description: Represents a hidden secret.
				private: true
			`,
	}

	testCaseWrapper := map[string]testCase{
		"root-relation-has-private-model": {
			inSkipPrivateModels: true,
			inSpec:              inSpec,
			supportingSpecs:     supportingSpecs,
			outDocs:             outDoc,
		},
	}
	runAllTestCases(t, testCaseWrapper)
}
