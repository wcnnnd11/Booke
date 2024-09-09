package middleware

import (
	"GVB_server/models/ctype"
	"GVB_server/models/res"
	"GVB_server/service/redis_ser"
	"GVB_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")

		if token == "" {
			res.FailWithMessage("未携带token登录", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token是错误的", c)
			c.Abort()
			return
		}
		//判断是否在redis中
		if redis_ser.CheckLogout(token) {
			res.FailWithMessage("token已失效", c)
			c.Abort()
			return
		}
		//登录的用户
		c.Set("claims", claims)
	}
}

func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")

		if token == "" {
			res.FailWithMessage("未携带token登录", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token是错误的", c)
			c.Abort()
			return
		}
		//判断是否在redis中
		if redis_ser.CheckLogout(token) {
			res.FailWithMessage("token已失效", c)
			c.Abort()
			return
		}
		//登录的用户
		if claims.Role != int(ctype.PermissionAdmin) {
			res.FailWithMessage("权限错误", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
