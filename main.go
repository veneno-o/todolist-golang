package main

import (
	"fmt"
	"todoList-golang/controller"
	"todoList-golang/model"
)

// @title			博客系统
// @version		1.0
// @description	Go 语言编程之旅：一起用 Go 做项目
// @termsOfService	https://github.com/go-programming-tour-book
func main() {
	// 强制日志颜色化
	// gin.ForceConsoleColor()

	// 记录到文件。
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// 连接数据库
	model.ConnectDB()
	r := controller.NewRouter()
	fmt.Println("Server is running on port 6666...")
	r.Run(":6666") // 监听并在 0.0.0.0:8080 上启动服务
}
