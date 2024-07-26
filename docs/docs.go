// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/all": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get all users",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number (default: 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of results per page (default: 10)",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.User"
                            }
                        }
                    }
                }
            }
        },
        "/user/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a user",
                "parameters": [
                    {
                        "description": "User object",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get a user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Address": {
            "type": "object",
            "required": [
                "city",
                "district",
                "pincode",
                "state"
            ],
            "properties": {
                "city": {
                    "type": "string",
                    "example": "Metropolis"
                },
                "district": {
                    "type": "string",
                    "example": "Central"
                },
                "lane": {
                    "type": "string",
                    "example": "1234 Elm St"
                },
                "pincode": {
                    "type": "integer",
                    "example": 123456
                },
                "state": {
                    "type": "string",
                    "example": "NY"
                },
                "village": {
                    "type": "string",
                    "example": "Springfield"
                }
            }
        },
        "model.Name": {
            "type": "object",
            "required": [
                "first_name",
                "last_name",
                "middle_name"
            ],
            "properties": {
                "first_name": {
                    "type": "string",
                    "example": "John"
                },
                "last_name": {
                    "type": "string",
                    "example": "Smith"
                },
                "middle_name": {
                    "type": "string",
                    "example": "Doe"
                }
            }
        },
        "model.User": {
            "type": "object",
            "required": [
                "address",
                "created_by",
                "email",
                "name",
                "password"
            ],
            "properties": {
                "active": {
                    "type": "boolean",
                    "example": true
                },
                "address": {
                    "$ref": "#/definitions/model.Address"
                },
                "created_at": {
                    "type": "string",
                    "example": "2024-07-27T00:00:00Z"
                },
                "created_by": {
                    "type": "string",
                    "example": "admin"
                },
                "deleted_at": {
                    "type": "string",
                    "example": "2024-07-27T00:00:00Z"
                },
                "deleted_by": {
                    "type": "string",
                    "example": "admin"
                },
                "email": {
                    "type": "string",
                    "example": "vishal"
                },
                "id": {
                    "type": "string",
                    "example": "123e4567-e89b-12d3-a456-426614174000"
                },
                "name": {
                    "$ref": "#/definitions/model.Name"
                },
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2024-07-27T00:00:00Z"
                },
                "updated_by": {
                    "type": "string",
                    "example": "admin"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "X-Token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "User API",
	Description:      "API for managing school",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
