// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/go-programming-tour-book",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/Todo/{id}": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "更新文章",
                "parameters": [
                    {
                        "description": "标签ID",
                        "name": "tag_id",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "文章标题",
                        "name": "title",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "文章简述",
                        "name": "desc",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "封面图片地址",
                        "name": "cover_image_url",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "文章内容",
                        "name": "content",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "修改者",
                        "name": "modified_by",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.Todo"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Todo": {
            "type": "object",
            "properties": {
                "completed_time": {
                    "description": "完成时间",
                    "type": "string"
                },
                "created_at": {
                    "description": "创建时间（由GORM自动管理）",
                    "type": "string"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "task_type": {
                    "description": "任务类型",
                    "type": "integer"
                },
                "title": {
                    "description": "DeletedAt     gorm.DeletedAt ` + "`" + `gorm:\"index\"` + "`" + `",
                    "type": "string"
                },
                "updated_at": {
                    "description": "最后一次更新时间（由GORM自动管理）",
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
	Title:            "博客系统",
	Description:      "Go 语言编程之旅：一起用 Go 做项目",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
