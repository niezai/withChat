package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"withChat/configs"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "server port:" + configs.Config.Server.IP + configs.Config.Server.Port})
}

func main() {
	configs.Init()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", Hello)
	r.Run(":8877") // 监听并在 0.0.0.0:8080 上启动服务
}
