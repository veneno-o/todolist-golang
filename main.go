package main

import Controller "todoList-golang/controller"

func main() {
	r := Controller.NewRouter()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
