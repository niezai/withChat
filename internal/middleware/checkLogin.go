package middleware

import (
	"github.com/gin-gonic/gin"
	"withChat/internal/dao"
	"withChat/internal/db/redis"
	"withChat/internal/e"
	"withChat/internal/response"
	"withChat/internal/util"
)

var whiteUri = []string{
	"",
}

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取路径 如果是登录路径,注册路径,白名单路径都不学需要进行验证
		uri := c.Request.RequestURI
		auth := c.GetHeader("Authentication")
		if uri != "/user/login" && uri != "/user/register" && !util.CheckExistsArr(uri, whiteUri) {
			if auth == "" { //没有token 直接返回意味着没有登录 返回 没有登录的信息
				response.LoginResponse(c, e.UnLogin, "", e.GetMsg(e.UnLogin))
				c.Abort()
			} else {
				cliam, err := util.ParseToken(auth)
				if err != nil {
					response.LoginResponse(c, e.UnLogin, "", e.GetMsg(e.UnLogin))
					c.Abort()
				}
				// 1验证token是否过期
				token := redis.Redis.Get(cliam.UserUuid)
				if token == auth { // 说明没过期
					userDao := dao.NewUserDao(c.Request.Context())
					_, err = userDao.FindUserByUuid(cliam.UserUuid) // 验证身份 是否存在该用户
					if err != nil {
						response.LoginResponse(c, e.UnLogin, "", e.GetMsg(e.UnLogin))
						c.Abort()
					}
				} else { // 说明过期了
					response.LoginResponse(c, e.CertificationFail, "", e.GetMsg(e.CertificationFail))
					c.Abort()
				}
			}
		}
		c.Next()
	}
}
