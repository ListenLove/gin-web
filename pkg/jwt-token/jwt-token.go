package jwtToken

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = "gin-web-jwt-secret"
var jwtExpiration = 24 * time.Hour
var jwtIssue = "gin-web"

// ErrorParedInvalidToken 定义解析token失败的错误
var ErrorParedInvalidToken = errors.New("解析token失败")

// UserToken 定义用户JWT信息结构体
type UserToken struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateUserToken 生成用户token
func GenerateUserToken(username string, userID int64) (string, error) {
	// 创建一个新的声明
	claims := UserToken{
		userID,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtExpiration)),
			Issuer:    jwtIssue,
		},
	}
	// 创建一个新的token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签名token并返回
	return token.SignedString([]byte(jwtSecret))
}

// ParseUserToken 解析用户token
func ParseUserToken(tokenString string) (*UserToken, error) {
	// 解析token
	claims := new(UserToken)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if token.Valid {
		return claims, nil
	}
	return nil, ErrorParedInvalidToken
}
