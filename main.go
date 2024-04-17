package main

import (
	"fmt"
	"todoList-golang/controller"
	"todoList-golang/model"
)

func main() {
	// 强制日志颜色化
	// gin.ForceConsoleColor()

	// 记录到文件。
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// 连接数据库
	model.ConnectDB()
	r := controller.NewRouter()
	fmt.Println("Server is running on port 8080...")
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
