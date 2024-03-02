package routes

import (
	"gin-web/controller"
	"gin-web/logger"
	"gin-web/middleware"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.New()
	r.Use(
		logger.GinLogger(),
		logger.GinRecovery(true),
		middleware.Translations(),
	)
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/token", controller.GenerateTokenHandler)
	r.POST("/parse-token", controller.ParseTokenHandler)
	// 请求头部需要携带 Authorization: Bearer token
	r.GET("/parse-token", middleware.JWTTokenParseMiddleware(), controller.ParseTokenFromHeaderHandler)
	return r
}
