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
        "/git/login": {
            "get": {
                "tags": [
                    "GitHub Oauth"
                ],
                "responses": {
                    "200": {
                        "description": "UserName\":userResp.Name,\"AvatarURL\":userResp.AvatarURL,}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/index": {
            "get": {
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/adduser": {
            "get": {
                "tags": [
                    "UserModule"
                ],
                "summary": "Add user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Twice Password",
                        "name": "repassword",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/deluser": {
            "delete": {
                "tags": [
                    "UserModule"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/index": {
            "get": {
                "tags": [
                    "User Home"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/updateuser": {
            "post": {
                "tags": [
                    "UserModule"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/userlist": {
            "get": {
                "tags": [
                    "UserModule"
                ],
                "summary": "Find all user",
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
