// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
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
        "/api/v1/urls": {
            "post": {
                "description": "generate a shorten URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "url"
                ],
                "summary": "Create Shorten URL",
                "parameters": [
                    {
                        "description": "Shorten URL information",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ShortenUrlResquest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/model.ShortenUrlResponse"
                        }
                    },
                    "400": {
                        "description": "wrong type or missing value",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    },
                    "500": {
                        "description": "server error",
                        "schema": {
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        },
        "/{url_id}": {
            "get": {
                "description": "redirect to origin url if {url_id} exist and without expired",
                "tags": [
                    "url"
                ],
                "summary": "Redirect to URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The id which response by /api/v1/urls",
                        "name": "url_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "redirect"
                    },
                    "404": {
                        "description": "{url_id} not found"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.HTTPError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "error message"
                }
            }
        },
        "model.ShortenUrlResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "\u003curl_id\u003e"
                },
                "shortUrl": {
                    "type": "string",
                    "example": "http://localhost/\u003curl_id\u003e"
                }
            }
        },
        "model.ShortenUrlResquest": {
            "type": "object",
            "required": [
                "expireAt",
                "url"
            ],
            "properties": {
                "expireAt": {
                    "type": "string",
                    "example": "2021-02-08T09:20:41Z"
                },
                "url": {
                    "type": "string",
                    "example": "\u003coriginal_url\u003e"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "URL Shortener API",
	Description:      "URL Shortener API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
