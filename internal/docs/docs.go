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
        "/info": {
            "get": {
                "description": "Возвращает песни с поддержкой пагинации",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Получение песни",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Название группы",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Название песни",
                        "name": "song",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Номер страницы (по умолчанию 1)",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Количество куплетов на странице (по умолчанию 2)",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Song"
                            }
                        }
                    }
                }
            }
        },
        "/songs": {
            "post": {
                "description": "Добавляет новую песню в библиотеку",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Добавление новой песни",
                "parameters": [
                    {
                        "description": "Данные песни",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SongCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/songs/{id}": {
            "delete": {
                "description": "Удаляет песню по её идентификатору",
                "tags": [
                    "songs"
                ],
                "summary": "Удаление песни",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор песни",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/songs/{id}/lyrics": {
            "get": {
                "description": "Возвращает текст песни с поддержкой пагинации по куплетам",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Получение текста песни",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID песни",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Номер страницы (по умолчанию 1)",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Количество куплетов на странице (по умолчанию 2)",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.SongCreateRequest": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                }
            }
        },
        "model.Song": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "releaseDate": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Online Music Library API",
	Description:      "REST API для управления онлайн-библиотекой песен.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
