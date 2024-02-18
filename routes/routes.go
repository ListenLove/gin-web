package routes

import (
	"gin-web/logger"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}
