package middleware

import (
	"gin-web/controller"
	jwtToken "gin-web/pkg/jwt-token"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTTokenParseMiddleware 解析JWT Token中间件
func JWTTokenParseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization 头部信息
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			controller.ResponseError(c, controller.CodeNoAuthInfo)
			c.Abort()
			return
		} else {
			// 获取token
			parts := strings.SplitN(auth, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				controller.ResponseError(c, controller.CodeNoBearerToken)
				c.Abort()
				return
			} else {
				token := parts[1]
				tokenInfo, err := jwtToken.ParseUserToken(token)
				if err != nil {
					controller.ResponseError(c, controller.CodeInvalidAuthInfo)
					c.Abort()
					return
				} else {
					// 将解析出来的token信息存入上下文
					c.Set("userID", tokenInfo.UserID)
					c.Set("username", tokenInfo.Username)
					c.Next()
				}
			}
		}
	}
}
