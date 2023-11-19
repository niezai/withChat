package v1

import (
	"github.com/gin-gonic/gin"
	"withChat/internal/model"
)

/**
1. 登录逻辑设计
在数据库中验证用户用户名和密码，生成token返回给用户
将用户上线的信息保存在redis中，考虑将用户登录信息保存在mysql中

*/

func Login(c *gin.Context) {
	var user model.User

}
