package main

import (
	"todoList-golang/controller"
)

func main() {
	r := controller.NewRouter()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
