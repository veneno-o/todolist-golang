package controller

import (
	"fmt"
	"todoList-golang/model"

	"github.com/gin-gonic/gin"
)

// 获取todo-task详情
func GetTodoTask(c *gin.Context) {
	fmt.Println("GetTodoTask")
}

// 获取todo-task列表
func GetTodoTaskList(c *gin.Context) {
	fmt.Println("GetTodoTaskList")
}

// 创建todo-task
func CreateTodoTask(c *gin.Context) {
	todo := model.Todo{}
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(400, gin.H{
			"code": 4001,
			"msg":  "请求参数错误",
			"data": nil,
		})
	}
	err := model.Create(&todo)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 5000,
			"msg":  "failed",
			"data": nil,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": nil,
		})
	}
}

// 更新todo-task
func UpdateTodoTask(c *gin.Context) {
	fmt.Println("update todo task")
}

// 删除todo-task
func DeleteTodoTask(c *gin.Context) {
	id := c.Param("id")
	err := model.Delete(id)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 5000,
			"msg":  "delete failed",
			"data": nil,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": nil,
		})
	}
}
