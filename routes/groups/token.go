package groups

import (
	"gin-web/controller"
	"gin-web/middleware"

	"github.com/gin-gonic/gin"
)

// TokenRegisterRoutes token路由组
func TokenRegisterRoutes(r *gin.Engine) {
	token := r.Group("/token")
	{
		token.GET("", controller.GenerateTokenHandler)
		token.POST("", controller.ParseTokenHandler)
		// 请求头部需要携带 Authorization: Bearer token
		token.GET("/must-with-token", middleware.JWTTokenParseMiddleware(), controller.ParseTokenFromHeaderHandler)
	}
}
