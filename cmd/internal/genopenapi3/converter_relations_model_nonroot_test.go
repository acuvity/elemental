package genopenapi3

import "testing"

func TestConverter_Do__modelRelations_nonRoot(t *testing.T) {
	t.Parallel()

	cases := map[string]testCase{

		"get-by-ID": {
			inSpec: `
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Resource
					package: usefulPackageName
					group: useful/thing
					description: useful description.
					get:
						description: Retrieves the resource with the given ID.
						parameters:
							entries:
							- name: duplicateParam
								description: This is a fancy parameter that should appear only once.
								type: time
							- name: duplicateParam
								description: This is a fancy parameter that should appear only once.
								type: time
			`,
			outDocs: map[string]string{
				"toplevel": `
					{
						"servers": [{ "url": "https://api.acuvity.ai"}],
						"openapi": "3.1.0",
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
									"description": "useful description.",
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
							"/resources/{id}": {
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
								"get": {
									"operationId": "get-Resource",
									"tags": ["useful/thing"],
									"description": "Retrieves the resource with the given ID.",
									"parameters": [
										{
											"description": "This is a fancy parameter that should appear only once.",
											"in": "query",
											"name": "duplicateParam",
											"schema": {
												"type": "string",
												"format": "date-time"
											}
										}
									],
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
		},

		"delete-by-ID": {
			inSpec: `
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Resource
					package: usefulPackageName
					group: useful/thing
					description: useful description.
					delete:
						description: Deletes the resource with the given ID.
						parameters:
							entries:
							- name: fancyParam
								description: This is a fancy parameter.
								type: duration
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
									"description": "useful description.",
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
							"/resources/{id}": {
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
								"delete": {
									"operationId": "delete-Resource",
									"tags": ["useful/thing"],
									"description": "Deletes the resource with the given ID.",
									"parameters": [
										{
											"description": "This is a fancy parameter.",
											"in": "query",
											"name": "fancyParam",
											"schema": {
												"type": "string"
											}
										}
									],
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
		},

		"put-by-ID": {
			inSpec: `
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Resource
					package: usefulPackageName
					group: useful/thing
					description: useful description.
					update:
						description: Updates the resource with the given ID.
						parameters:
							entries:
							- name: fancyParam
								description: This is a fancy parameter.
								type: enum
								allowed_choices:
								- Choice1
								- Choice2
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
									"description": "useful description.",
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
							"/resources/{id}": {
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
									"operationId": "update-Resource",
									"tags": ["useful/thing"],
									"description": "Updates the resource with the given ID.",
									"parameters": [
										{
											"description": "This is a fancy parameter.",
											"in": "query",
											"name": "fancyParam",
											"schema": {
												"enum": ["Choice1", "Choice2"]
											}
										}
									],
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
		},

		"get-put-delete-by-ID--do-not-duplicate-param-ID": {
			inSpec: `
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Resource
					package: usefulPackageName
					group: useful/thing
					description: useful description.
					get:
						description: Retrieves the resource with the given ID.
					delete:
						description: Deletes the resource with the given ID.
					update:
						description: Updates the resource with the given ID.
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
									"description": "useful description.",
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
							"/resources/{id}": {
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
								"get": {
									"operationId": "get-Resource",
									"tags": ["useful/thing"],
									"description": "Retrieves the resource with the given ID.",
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
								},
								"delete": {
									"operationId": "delete-Resource",
									"tags": ["useful/thing"],
									"description": "Deletes the resource with the given ID.",
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
								},
								"put": {
									"operationId": "update-Resource",
									"tags": ["useful/thing"],
									"description": "Updates the resource with the given ID.",
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
		},
	}
	runAllTestCases(t, cases)
}
