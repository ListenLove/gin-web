package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// EncryptPassword 加密密码
func EncryptPassword(password string) string {
	hash := md5.New()                 // 创建一个新的MD5实例
	hash.Write([]byte(password))      // 将密码写入到hash实例中
	md5sum := hash.Sum(nil)           // 计算MD5校验和
	return hex.EncodeToString(md5sum) // 将MD5校验和转换为16进制字符串
}
