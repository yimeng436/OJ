// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/question/add": {
            "post": {
                "description": "添加问题",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "添加问题",
                "parameters": [
                    {
                        "description": "question",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/pb.QuestionAddRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/question/delete/{id}": {
            "get": {
                "description": "删除题目",
                "produces": [
                    "application/json"
                ],
                "summary": "删除题目",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Question ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/question/get": {
            "get": {
                "description": "id获取完整题目内容",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "id获取完整题目内容",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "数据的ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/question/get/vo/{id}": {
            "get": {
                "description": "根据Id查询QuestionVO",
                "produces": [
                    "application/json"
                ],
                "summary": "根据Id查询QuestionVO",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Question ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/question/get/{id}": {
            "get": {
                "description": "根据id获取题目",
                "produces": [
                    "application/json"
                ],
                "summary": "根据id获取题目",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Question ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/question/list": {
            "post": {
                "description": "分页查询问题（管理员）",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "分页查询问题（管理员）",
                "parameters": [
                    {
                        "description": "QuestionVoPage",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/pb.GetQuestionPageRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/question/list/vo": {
            "post": {
                "description": "分页查询问题（用户）",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "分页查询问题（用户）",
                "parameters": [
                    {
                        "description": "QuestionVoPage",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/pb.GetQuestionPageRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/question/submit/do": {
            "post": {
                "description": "提交问题",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "提交问题",
                "parameters": [
                    {
                        "description": "QuestionSubmitAddRequest",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/pb.QuestionSubmitAddRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/question/submit/query": {
            "post": {
                "description": "查询问题提交信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "查询问题提交信息",
                "parameters": [
                    {
                        "description": "QuestionSubmit",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/pb.QuestionSubmitQueryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/question/update": {
            "post": {
                "description": "修改问题",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "修改问题",
                "parameters": [
                    {
                        "description": "question",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/pb.QuestionInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/test/{id}": {
            "get": {
                "description": "描述信息",
                "produces": [
                    "application/json"
                ],
                "summary": "测试接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/user/getLoginUser": {
            "get": {
                "description": "获取登录用户",
                "produces": [
                    "application/json"
                ],
                "summary": "获取登录用户",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "用户登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/pb.UserLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "pb.Context": {
            "type": "object",
            "properties": {
                "context": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "pb.GetQuestionPageRequest": {
            "type": "object",
            "properties": {
                "page": {
                    "$ref": "#/definitions/pb.Page"
                },
                "question": {
                    "$ref": "#/definitions/pb.QuestionInfo"
                }
            }
        },
        "pb.JudgeCase": {
            "type": "object",
            "properties": {
                "inputs": {
                    "description": "用例输入",
                    "type": "string"
                },
                "outputs": {
                    "description": "用例输出",
                    "type": "string"
                }
            }
        },
        "pb.JudgeConfig": {
            "type": "object",
            "properties": {
                "memoryLimit": {
                    "description": "内存限制",
                    "type": "integer"
                },
                "stackLimit": {
                    "description": "堆栈限制",
                    "type": "integer"
                },
                "timeLimit": {
                    "description": "时间限制",
                    "type": "integer"
                }
            }
        },
        "pb.Page": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                }
            }
        },
        "pb.QuestionAddRequest": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "judgeCase": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pb.JudgeCase"
                    }
                },
                "judgeConfig": {
                    "$ref": "#/definitions/pb.JudgeConfig"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "pb.QuestionInfo": {
            "type": "object",
            "properties": {
                "acceptedNum": {
                    "description": "题目通过数",
                    "type": "integer"
                },
                "answer": {
                    "description": "题目答案",
                    "type": "string"
                },
                "content": {
                    "description": "内容",
                    "type": "string"
                },
                "createTime": {
                    "description": "创建时间",
                    "type": "string"
                },
                "favourNum": {
                    "description": "收藏数",
                    "type": "integer"
                },
                "id": {
                    "description": "id",
                    "type": "integer"
                },
                "isDelete": {
                    "description": "是否删除",
                    "type": "integer"
                },
                "judgeCase": {
                    "description": "判题用例（json 数组）",
                    "type": "string"
                },
                "judgeConfig": {
                    "description": "判题配置（json 对象）",
                    "type": "string"
                },
                "submitNum": {
                    "description": "题目提交数",
                    "type": "integer"
                },
                "tags": {
                    "description": "标签列表（json 数组）",
                    "type": "string"
                },
                "thumbNum": {
                    "description": "点赞数",
                    "type": "integer"
                },
                "title": {
                    "description": "标题",
                    "type": "string"
                },
                "updateTime": {
                    "description": "更新时间",
                    "type": "string"
                },
                "userId": {
                    "description": "创建用户 id",
                    "type": "integer"
                }
            }
        },
        "pb.QuestionSubmitAddRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "ctx": {
                    "$ref": "#/definitions/pb.Context"
                },
                "language": {
                    "type": "string"
                },
                "questionId": {
                    "type": "integer"
                }
            }
        },
        "pb.QuestionSubmitQueryRequest": {
            "type": "object",
            "properties": {
                "ctx": {
                    "$ref": "#/definitions/pb.Context"
                },
                "language": {
                    "type": "string"
                },
                "page": {
                    "$ref": "#/definitions/pb.Page"
                },
                "questionId": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "pb.UserLoginRequest": {
            "type": "object",
            "properties": {
                "userName": {
                    "description": "用户昵称",
                    "type": "string"
                },
                "userPassword": {
                    "description": "密码",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "swaggo测试",
	Description:      "swaggo测试API文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
