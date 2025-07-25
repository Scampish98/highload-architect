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
        "/user/get/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Получение информации о пользователе по его идентификатору",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user info by ID",
                "parameters": [
                    {
                        "type": "integer",
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
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/user/search/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Получения списка пользователей, удовлетворяющих фильтрам",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Search users by filter",
                "parameters": [
                    {
                        "type": "string",
                        "example": "van",
                        "name": "first_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "vano",
                        "name": "last_name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/user/signin": {
            "post": {
                "description": "Вход для существующего пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "User sign in",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.SignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.SignInResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/user/signup": {
            "post": {
                "description": "Регистрация нового пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "User sign up",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Error": {
            "type": "object",
            "required": [
                "code",
                "message"
            ],
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 1
                },
                "message": {
                    "type": "string",
                    "example": "internal server error"
                }
            }
        },
        "models.Sex": {
            "type": "string",
            "enum": [
                "any",
                "male",
                "female"
            ],
            "x-enum-varnames": [
                "SexAny",
                "SexMale",
                "SexFemale"
            ]
        },
        "models.User": {
            "type": "object",
            "required": [
                "first_name",
                "id",
                "last_name",
                "sex",
                "username"
            ],
            "properties": {
                "biography": {
                    "type": "string",
                    "example": "London is the capital of Great Britain"
                },
                "birthdate": {
                    "type": "string",
                    "format": "date",
                    "example": "1990-01-01"
                },
                "city": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 1,
                    "example": "Saint-Petersburg"
                },
                "first_name": {
                    "type": "string",
                    "example": "Ivan"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "last_name": {
                    "type": "string",
                    "example": "Ivanov"
                },
                "sex": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Sex"
                        }
                    ],
                    "example": "male"
                },
                "username": {
                    "type": "string",
                    "example": "my_username"
                }
            }
        },
        "requests.SignInRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 1,
                    "example": "1234567"
                },
                "username": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 1,
                    "example": "myusername"
                }
            }
        },
        "requests.SignUpRequest": {
            "type": "object",
            "required": [
                "first_name",
                "last_name",
                "password",
                "username"
            ],
            "properties": {
                "biography": {
                    "type": "string",
                    "example": "London is the capital of Great Britain"
                },
                "birthdate": {
                    "type": "string",
                    "format": "date",
                    "example": "2006-01-02"
                },
                "city": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 1,
                    "example": "Saint-Petersburg"
                },
                "first_name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 1,
                    "example": "Ivan"
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 1,
                    "example": "Ivanov"
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 1,
                    "example": "123456"
                },
                "sex": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Sex"
                        }
                    ],
                    "example": "male"
                },
                "username": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 1,
                    "example": "myusername"
                }
            }
        },
        "responses.SignInResponse": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "X-Api-Key",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Highload-Architect Homework Backend API",
	Description:      "Swagger API for Highload-Architect Homework API Gateway backend http-server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
