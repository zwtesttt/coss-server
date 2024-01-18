// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplatemsg = `{
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
        "/msg/dialog/list": {
            "get": {
                "description": "获取用户对话列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取用户对话列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/msg/list/user": {
            "get": {
                "description": "获取私聊消息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取私聊消息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "类型",
                        "name": "type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "消息",
                        "name": "content",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page_num",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页大小",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/msg/send/group": {
            "post": {
                "description": "发送群聊消息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "发送群聊消息",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SendGroupMsgRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/msg/send/user": {
            "post": {
                "description": "发送私聊消息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "发送私聊消息",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SendUserMsgRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/msg/ws": {
            "get": {
                "description": "websocket请求",
                "summary": "websocket请求",
                "responses": {}
            }
        }
    },
    "definitions": {
        "model.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "model.SendGroupMsgRequest": {
            "type": "object",
            "required": [
                "content",
                "dialog_id",
                "group_id",
                "type"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "dialog_id": {
                    "type": "integer"
                },
                "group_id": {
                    "type": "integer"
                },
                "replay_id": {
                    "type": "integer"
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "model.SendUserMsgRequest": {
            "type": "object",
            "required": [
                "content",
                "dialog_id",
                "receiver_id",
                "type"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "dialog_id": {
                    "type": "integer"
                },
                "receiver_id": {
                    "type": "string"
                },
                "replay_id": {
                    "type": "integer"
                },
                "type": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfomsg holds exported Swagger Info so clients can modify it
var SwaggerInfomsg = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "",
	InfoInstanceName: "msg",
	SwaggerTemplate:  docTemplatemsg,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfomsg.InstanceName(), SwaggerInfomsg)
}
