// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/v1/property-category/create": {
            "post": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "Create a PropertyCategory",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "PropertyCategory"
                ],
                "summary": "Create a PropertyCategory",
                "parameters": [
                    {
                        "description": "Create a PropertyCategory",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreatePropertyCategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "PropertyCategory response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/dto.PropertyCategoryResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            }
        },
        "/v1/property-category/delete/{id}": {
            "delete": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "Get a PropertyCategory",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "PropertyCategory"
                ],
                "summary": "Get a PropertyCategory",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "PropertyCategory response",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            }
        },
        "/v1/property-category/get/{id}": {
            "get": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "Delete a PropertyCategory",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "PropertyCategory"
                ],
                "summary": "Delete a PropertyCategory",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "response",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            }
        },
        "/v1/property-category/update/{id}": {
            "put": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "Update a PropertyCategory",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "PropertyCategory"
                ],
                "summary": "Update a PropertyCategory",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update a PropertyCategory",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdatePropertyCategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "PropertyCategory response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/dto.PropertyCategoryResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            }
        },
        "/v1/users/login/phone": {
            "post": {
                "description": "RegisterLoginByPhone",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "RegisterLoginByPhone",
                "parameters": [
                    {
                        "description": "Create a RegisterLoginByPhone",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterLoginByPhone"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "RegisterLoginByPhone response",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            }
        },
        "/v1/users/login/username": {
            "post": {
                "description": "LoginByUsername",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "LoginByUsername",
                "parameters": [
                    {
                        "description": "Create a LoginByUsername",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginByUsername"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "LoginByUsername response",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            }
        },
        "/v1/users/register/phone": {
            "post": {
                "description": "RegisterLoginByPhone",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "RegisterLoginByPhone",
                "parameters": [
                    {
                        "description": "Create a RegisterLoginByPhone",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterLoginByPhone"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "RegisterLoginByPhone response",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            }
        },
        "/v1/users/register/username": {
            "post": {
                "description": "RegisterByUsername",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "RegisterByUsername",
                "parameters": [
                    {
                        "description": "Create a RegisterByUsername",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterUserByUsername"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "RegisterByUsername response",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            }
        },
        "/v1/users/send-otp": {
            "post": {
                "description": "SendOtp",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "SendOtp",
                "parameters": [
                    {
                        "description": "Create a GetOtpRequest",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.GetOtpRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "SendOtp response",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreatePropertyCategoryRequest": {
            "type": "object",
            "required": [
                "icon",
                "name"
            ],
            "properties": {
                "icon": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.GetOtpRequest": {
            "type": "object",
            "required": [
                "mobile_number"
            ],
            "properties": {
                "mobile_number": {
                    "type": "string",
                    "maxLength": 11,
                    "minLength": 11
                }
            }
        },
        "dto.LoginByUsername": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 3
                },
                "username": {
                    "type": "string",
                    "minLength": 3
                }
            }
        },
        "dto.PropertyCategoryResponse": {
            "type": "object",
            "properties": {
                "icon": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.RegisterLoginByPhone": {
            "type": "object",
            "required": [
                "otp",
                "phone"
            ],
            "properties": {
                "otp": {
                    "type": "string"
                },
                "phone": {
                    "type": "string",
                    "maxLength": 11,
                    "minLength": 11
                }
            }
        },
        "dto.RegisterUserByUsername": {
            "type": "object",
            "required": [
                "first_name",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.UpdatePropertyCategoryRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "helper.Response": {
            "type": "object",
            "properties": {
                "error": {},
                "result": {},
                "result_code": {
                    "type": "integer"
                },
                "success": {
                    "type": "boolean"
                },
                "validation_error": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/validators.ValidationError"
                    }
                }
            }
        },
        "validators.ValidationError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "property": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
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