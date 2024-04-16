package main

import (
	"io"
	"os"
	"todoList-golang/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	// 强制日志颜色化
	gin.ForceConsoleColor()

	// 记录到文件。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := controller.NewRouter()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
