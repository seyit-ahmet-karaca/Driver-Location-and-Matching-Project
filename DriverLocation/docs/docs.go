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
        "/api/driver-locations": {
            "post": {
                "description": "Creates driver location by longitude and latitude or bulk file.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Driver"
                ],
                "summary": "Create Driver Location.",
                "parameters": [
                    {
                        "type": "number",
                        "description": "latitude",
                        "name": "latitude",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "description": "longitude",
                        "name": "longitude",
                        "in": "query"
                    },
                    {
                        "type": "file",
                        "description": "bulk location file",
                        "name": "file",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/search/location": {
            "get": {
                "description": "gets nearest driver location by given point and radius",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Driver"
                ],
                "summary": "Gets nearest driver location",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "radius",
                        "name": "radius",
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
                        "type": "number",
                        "description": "latitude",
                        "name": "latitude",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.NearestDriverResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.NearestDriverResponse": {
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
	Host:             "localhost:3000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Driver Location API",
	Description:      "Driver Location API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
