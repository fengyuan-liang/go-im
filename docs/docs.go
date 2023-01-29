// Code generated by swaggo/swag. DO NOT EDIT
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
        "/user/delOne": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "删除一个用户"
                ],
                "summary": "参数类型：{\"userId\":123456,\"isLogicDel\":true}",
                "parameters": [
                    {
                        "description": "上传的JSON",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserBasic"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "UserBasic"
                        }
                    }
                }
            }
        },
        "/user/pageQuery": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关"
                ],
                "summary": "分页查询",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "第几页",
                        "name": "pageNo",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "每页多少条",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "UserBasic"
                        }
                    }
                }
            }
        },
        "/user/pageQueryByFilter": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关"
                ],
                "summary": "前端请求参数应为：http://xx:xx/pageQueryByFilter?pageSize=1\u0026pageNo=1\u0026name=1\u0026age=2\u0026email=xxx@xxx",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "UserBasic"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "创建一个用户"
                ],
                "summary": "用于用户注册",
                "parameters": [
                    {
                        "description": "上传的JSON",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserBasic"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "UserBasic"
                        }
                    }
                }
            }
        },
        "/user/update": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "更新用户信息"
                ],
                "parameters": [
                    {
                        "description": "上传的JSON",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserBasic"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "models.UserBasic": {
            "type": "object"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}