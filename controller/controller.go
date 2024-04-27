package controller

import "github.com/gin-gonic/gin"

// NewRouter 路由初始化
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 简单的路由组: todo-task
	baseGroup := r.Group("/")
	v1 := baseGroup.Group("todo-task")
	{
		// 获取列表
		v1.GET("/list", GetTodoTaskList)
		// 获取详情
		v1.GET("/:id", GetTodoTask)
		// 修改某个任务
		v1.PUT("/:id", UpdateTodoTask)
		// 新建任务
		v1.POST("/", CreateTodoTask)
		// 删除任务
		v1.DELETE("/:id", DeleteTodoTask)
	}

	v2 := baseGroup.Group("user")
	{
		v2.GET("/login", Login)
		v2.GET("/logout", Logout)
		v2.POST("/register", Register)
		v2.GET("/list", GetUserList)
		v2.GET("/:id", GetUser)
		v2.PUT("/:id", UpdateUser)
		v2.DELETE("/:id", DeleteUser)
	}
	return r
}
