package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResCode int64

const (
	CodeSuccess ResCode = 10000 + iota
	CodeRequestFailed
	CodeInvalidParams
	CodeUserExist
	CodeUserNotExist
	CodeNeedLogin
	CodeLoginFailed
	CodeServerBusy
	CodeInternalServerError
	CodeNotAllowedOperation
	CodeNoAuthInfo
	CodeInvalidAuthInfo
	CodeNoBearerToken
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:             "请求成功",
	CodeRequestFailed:       "请求失败",
	CodeInvalidParams:       "请求参数错误",
	CodeUserExist:           "用户已存在",
	CodeUserNotExist:        "用户不存在",
	CodeNeedLogin:           "需要登录",
	CodeLoginFailed:         "登录失败",
	CodeServerBusy:          "服务繁忙",
	CodeInternalServerError: "服务器内部错误",
	CodeNotAllowedOperation: "不允许的操作",
	CodeNoAuthInfo:          "没有权限信息",
	CodeNoBearerToken:       "没有Bearer Token",
	CodeInvalidAuthInfo:     "无效的权限信息",
}

func (r ResCode) msg() string {
	msg, ok := codeMsgMap[r]
	if ok {
		return msg
	}
	return codeMsgMap[CodeServerBusy]
}

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, ResponseData{
		Code: code,
		Msg:  code.msg(),
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg string) {
	if msg == "" {
		msg = codeMsgMap[code]
	}
	c.JSON(http.StatusOK, ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.msg(),
		Data: data,
	})
}

func ResponseSuccessWithMsgAndData(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code: CodeSuccess,
		Msg:  msg,
		Data: data,
	})
}
