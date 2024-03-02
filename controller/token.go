package controller

import (
	bindvalidator "gin-web/pkg/bind-validator"
	jwtToken "gin-web/pkg/jwt-token"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ParamsToken 定义Token参数结构体, 只是简单测验一下参数校验，一般我会新建一个文件夹专门放置参数
type ParamsToken struct {
	Token string `json:"token" binding:"required"`
}

// GenerateTokenHandler 生成Token
func GenerateTokenHandler(c *gin.Context) {
	// 生成Token
	token, err := jwtToken.GenerateUserToken("admin", 1)
	if err != nil {
		zap.L().Error("GenerateTokenHandler 失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, token)
}

// ParseTokenHandler 解析Token, 同时测试一下参数校验
func ParseTokenHandler(c *gin.Context) {
	// 绑定参数
	var params ParamsToken
	if _, err := bindvalidator.BindAndValid(c, &params); err != nil {
		zap.L().Error("ParseTokenHandler 失败", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	// 解析Token
	tokenInfo, err := jwtToken.ParseUserToken(params.Token)
	if err != nil {
		zap.L().Error("ParseTokenHandler 失败", zap.Error(err))
		ResponseError(c, CodeInvalidAuthInfo)
		return
	}
	ResponseSuccess(c, tokenInfo)
}

// ParseTokenFromHeaderHandler 从Header中解析Token
func ParseTokenFromHeaderHandler(c *gin.Context) {
	// 走到这里，说明已经通过了JWTTokenParseMiddleware中间件，所以可以直接从上下文中获取token信息
	userID, _ := c.Get("userID")
	username, _ := c.Get("username")
	ResponseSuccess(c, gin.H{
		"userID":   userID,
		"username": username,
	})
}
