package router

import (
	"fmt"

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
	fmt.Println("CreateTodoTask")

}

// 更新todo-task
func UpdateTodoTask(c *gin.Context) {
	fmt.Println("update todo task")
}
