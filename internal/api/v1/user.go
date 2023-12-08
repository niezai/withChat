package v1

import (
	"github.com/gin-gonic/gin"
	"withChat/internal/e"
	"withChat/internal/model"
	"withChat/internal/response"
	"withChat/internal/service"
)

/**
1. 登录逻辑设计
在数据库中验证用户用户名和密码，生成token返回给用户
将用户上线的信息保存在redis中，考虑将用户登录信息保存在mysql中

*/

var userService *service.UserService

func Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithError(e.FAIL, err)
		return
	}
	token, err := userService.Login(&user, c.Request.Context())
	if err != nil {
		response.LoginResponse(c, e.LoginFail, "", e.GetMsg(e.LoginFail))
	} else {
		response.LoginResponse(c, e.SUCCESS, token, e.GetMsg(e.SUCCESS))
	}
}

func Register(c *gin.Context) {
	var user *model.User
	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithError(e.FAIL, err)
		return
	}
	err := userService.Register(user, c.Request.Context())
	if err != nil {
		response.Response(c, e.RegisterFail, "", e.GetMsg(e.RegisterFail))
	} else {
		response.Response(c, e.SUCCESS, "", "注册成功,即将跳转到登录页面")
	}
}
