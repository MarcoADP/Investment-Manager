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
        "/api/v1/acoes": {
            "get": {
                "description": "Get All Ação BR",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "acoes"
                ],
                "summary": "Get All Acão BR",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1_response.AcaoBrResponse"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new acao",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "acoes"
                ],
                "summary": "Create a acao",
                "parameters": [
                    {
                        "description": "AcaoBrRequest info",
                        "name": "acao",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1_request.AcaoBrRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1_response.AcaoBrResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/acoes/{id}": {
            "get": {
                "description": "Get a Ação BR by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "acoes"
                ],
                "summary": "Get a Acão BR",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "AcaoBr ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1_response.AcaoBrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a acao by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "acoes"
                ],
                "summary": "Update a acao",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "AcaoBr ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "AcaoBrRequest info",
                        "name": "acao",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1_request.AcaoBrRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1_response.AcaoBrResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a acao by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "acoes"
                ],
                "summary": "Delete a acao",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "AcaoBr ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/fundos-imobiliarios": {
            "get": {
                "description": "Get All Fundos Imobiliarios",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "fundos-imobiliarios"
                ],
                "summary": "Get All Fundos Imobiliarios",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1_response.FundoImobiliarioResponse"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new Fundo Imobiliario",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "fundos-imobiliarios"
                ],
                "summary": "Create a Fundo Imobiliario",
                "parameters": [
                    {
                        "description": "FundoImobiliarioRequest info",
                        "name": "fundoImobiliario",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1_request.FundoImobiliarioRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1_response.FundoImobiliarioResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/fundos-imobiliarios/{id}": {
            "get": {
                "description": "Get a Fundo Imobiliario by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "fundos-imobiliarios"
                ],
                "summary": "Get a Fundo Imobiliario",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "FundoImobiliario ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1_response.FundoImobiliarioResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a Fundo Imobiliario by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "fundos-imobiliarios"
                ],
                "summary": "Update a Fundo Imobiliario",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "FundoImobiliario ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "FundoImobiliarioRequest info",
                        "name": "fundoImobiliario",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1_request.FundoImobiliarioRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1_response.FundoImobiliarioResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a Fundo Imobiliario by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "fundos-imobiliarios"
                ],
                "summary": "Delete a Fundo Imobiliario",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "FundoImobiliario ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_MarcoADP_Investment-Manager_pkg_api_v1.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_MarcoADP_Investment-Manager_pkg_api_v1.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Bad Request"
                }
            }
        },
        "github_com_MarcoADP_Investment-Manager_pkg_api_v1_request.AcaoBrRequest": {
            "type": "object",
            "properties": {
                "cnpj": {
                    "type": "string"
                },
                "codigo": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                },
                "setor": {
                    "type": "string"
                }
            }
        },
        "github_com_MarcoADP_Investment-Manager_pkg_api_v1_request.FundoImobiliarioRequest": {
            "type": "object",
            "properties": {
                "cnpj": {
                    "type": "string"
                },
                "codigo": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                },
                "segmento": {
                    "type": "string"
                },
                "tipo": {
                    "type": "string"
                }
            }
        },
        "github_com_MarcoADP_Investment-Manager_pkg_api_v1_response.AcaoBrResponse": {
            "type": "object",
            "properties": {
                "cnpj": {
                    "type": "string"
                },
                "codigo": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                },
                "setor": {
                    "type": "string"
                }
            }
        },
        "github_com_MarcoADP_Investment-Manager_pkg_api_v1_response.FundoImobiliarioResponse": {
            "type": "object",
            "properties": {
                "cnpj": {
                    "type": "string"
                },
                "codigo": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                },
                "segmento": {
                    "type": "string"
                },
                "tipo": {
                    "type": "string"
                }
            }
        },
        "pkg_api_v1.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Bad Request"
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
