// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplatelive = `{
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
        "/live/user/create": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "创建用户通话",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "liveUser"
                ],
                "summary": "创建用户通话",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserCallRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "url=webRtc服务器地址 token=加入通话的token room_name=房间名称 room_id=房间id",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object",
                                            "additionalProperties": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/live/user/leave": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "结束通话",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "liveUser"
                ],
                "summary": "结束通话",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserLeaveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            }
        },
        "/live/user/show": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "获取通话房间信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "liveUser"
                ],
                "summary": "获取通话房间信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "房间名",
                        "name": "room",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "sid=房间id identity=用户id state=用户状态 (0=加入中 1=已加入 2=已连接 3=断开连接 ) joined_at=加入时间 name=用户名称 is_publisher=是否是创建者",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.UserShowResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CallOption": {
            "type": "object",
            "properties": {
                "audio_enabled": {
                    "description": "是否启用音频",
                    "type": "boolean"
                },
                "codec": {
                    "description": "编解码器",
                    "type": "string"
                },
                "frame_rate": {
                    "description": "帧率",
                    "type": "integer"
                },
                "resolution": {
                    "description": "分辨率",
                    "type": "string"
                },
                "video_enabled": {
                    "description": "是否启用视频",
                    "type": "boolean"
                }
            }
        },
        "dto.ParticipantState": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3
            ],
            "x-enum-varnames": [
                "ParticipantInfo_JOINING",
                "ParticipantInfo_JOINED",
                "ParticipantInfo_ACTIVE",
                "ParticipantInfo_DISCONNECTED"
            ]
        },
        "dto.Response": {
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
        "dto.UserCallRequest": {
            "type": "object",
            "required": [
                "user_id"
            ],
            "properties": {
                "option": {
                    "$ref": "#/definitions/dto.CallOption"
                },
                "user_id": {
                    "description": "接收视频通话的用户ID",
                    "type": "string"
                }
            }
        },
        "dto.UserLeaveRequest": {
            "type": "object",
            "required": [
                "room",
                "user_id"
            ],
            "properties": {
                "room": {
                    "description": "房间名称",
                    "type": "string"
                },
                "user_id": {
                    "description": "用户ID",
                    "type": "string"
                }
            }
        },
        "dto.UserShowResponse": {
            "type": "object",
            "properties": {
                "identity": {
                    "type": "string"
                },
                "is_publisher": {
                    "type": "boolean"
                },
                "joined_at": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "sid": {
                    "type": "string"
                },
                "state": {
                    "$ref": "#/definitions/dto.ParticipantState"
                }
            }
        }
    }
}`

// SwaggerInfolive holds exported Swagger Info so clients can modify it
var SwaggerInfolive = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "",
	InfoInstanceName: "live",
	SwaggerTemplate:  docTemplatelive,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfolive.InstanceName(), SwaggerInfolive)
}