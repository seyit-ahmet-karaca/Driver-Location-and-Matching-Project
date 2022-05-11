// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/search/driver": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Find nearest driver to user from given location point",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Driver Matching"
                ],
                "summary": "Find nearest driver to user",
                "parameters": [
                    {
                        "type": "number",
                        "description": "latitude",
                        "name": "latitude",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "longitude",
                        "name": "longitude",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "radius",
                        "name": "radius",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/testDto.MatchResponse"
                        }
                    }
                }
            }
        },
        "/token": {
            "get": {
                "description": "Gets valid jwt token to send request matching API",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "JWT"
                ],
                "summary": "Get valid jwt token",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/tokenNotAuthenticated": {
            "get": {
                "description": "Gets unauthanticated jwt token to send request matching API",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "JWT"
                ],
                "summary": "Get unauthanticated jwt token",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "testDto.MatchResponse": {
            "type": "object",
            "properties": {
                "Title": {
                    "type": "string"
                },
                "distance": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3001",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Matching API",
	Description:      "Matching API to send request Driver Location API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}