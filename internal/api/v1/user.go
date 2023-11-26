package v1

import (
	"github.com/gin-gonic/gin"
	"withChat/internal/model"
	"withChat/internal/msg"
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
	var user model.LoginUser
	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithError(msg.ERR_PARAMETER, err)
		return
	}

	err := userService.Login(&user, c.Request.Context())
	if err != nil {
		response.Response(c, msg.FAIL, nil, msg.GetMsg(msg.FAIL))
	} else {
		response.Response(c, msg.SUCCESS, nil, msg.GetMsg(msg.SUCCESS))
	}
}

func Register(c *gin.Context) {

}
