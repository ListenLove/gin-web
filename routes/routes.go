package routes

import (
	"gin-web/logger"
	"gin-web/middleware"
	"gin-web/routes/groups"

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
	// 注册 token 路由
	groups.TokenRegisterRoutes(r)
	return r
}
