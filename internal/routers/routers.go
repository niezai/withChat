package routers

import (
	"github.com/gin-gonic/gin"
	"withChat/internal/api/test"
	"withChat/pkg/gin/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(middleware.Cors()) // 配置跨域请求

	r.Use(gin.Recovery()) // 处理错误信息

	r.Use(middleware.LoggerToFile()) // 将请求日志放到log文件中

	r.GET("hello", test.Hello)

	v1 := r.Group("")
	{
		v1.GET("/user/login")
	}
	return r
}
