package controller

import "github.com/gin-gonic/gin"

// NewRouter 路由初始化
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 简单的路由组: todo-task
	v1 := r.Group("/todo-task")
	{
		// 获取详情
		v1.GET("/:id", GetTodoTask)
		// 获取列表
		v1.GET("/list", GetTodoTaskList)
		// 修改某个任务
		v1.PUT("/:id", UpdateTodoTask)
		// 新建任务
		v1.POST("/", CreateTodoTask)
		// 删除任务
		v1.DELETE("/:id", DeleteTodoTask)
	}
	return r
}
