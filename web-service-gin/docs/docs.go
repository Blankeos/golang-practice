// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/album": {
            "get": {
                "description": "gets one (1) album using a query.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "albums"
                ],
                "summary": "gets one (1) album using a query.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "album search by id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully fetched album.",
                        "schema": {
                            "$ref": "#/definitions/main.album"
                        }
                    }
                }
            }
        },
        "/albums": {
            "get": {
                "description": "gets all albums from the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "albums"
                ],
                "summary": "gets all albums",
                "responses": {
                    "200": {
                        "description": "Successfully retrieved the albums",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.album"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "postAlbums adds a new album from JSON received in the request body.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "albums"
                ],
                "summary": "adds a new album",
                "parameters": [
                    {
                        "description": "Album object",
                        "name": "album",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.album"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully added album.",
                        "schema": {
                            "$ref": "#/definitions/main.album"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.album": {
            "type": "object",
            "properties": {
                "artist": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Carlo's Album API",
	Description:      "This is a generated swagger API from the Album API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
