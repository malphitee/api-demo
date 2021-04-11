package main

import (
	"api/controller"
	"api/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8888")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.ApiCheckSignature())
	// Ping 测试路由
	r.GET("/ping", controller.Resp)

	r.POST("/ping", controller.Resp)
	return r
}