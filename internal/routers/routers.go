package routers

import (
	"github.com/gin-gonic/gin"
	"withChat/internal/api/test"
	"withChat/pkg/gin/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	r.Use(middleware.LoggerToFile())

	r.GET("hello", test.Hello)
	return r
}
