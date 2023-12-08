package router

import (
	"github.com/gin-gonic/gin"
	"withChat/internal/api/test"
	v1 "withChat/internal/api/v1"
	"withChat/internal/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(middleware.Cors()) // 配置跨域请求

	r.Use(gin.Recovery()) // 处理错误信息
	r.Use(middleware.CheckLogin())
	r.Use(middleware.LoggerToFile()) // 将请求日志放到log文件中

	r.GET("hello", test.Hello)

	user := r.Group("/user")
	{
		user.GET("/login", v1.Login)
		user.GET("/register", v1.Register)
	}
	return r
}
