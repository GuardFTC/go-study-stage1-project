// Package utils @Author:冯铁城 [17615007230@163.com] 2023-06-20 15:19:31
package utils

import (
	"math/rand"
)

// 定义字符集常量
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString 随机长度字符串
func RandomString(length int) string {

	//1.创建一个字节切片,长度为指定的长度
	randomBytes := make([]byte, length)

	//2.生成随机字符串
	for i := 0; i < length; i++ {
		randomBytes[i] = charset[rand.Intn(len(charset))]
	}

	//3.将字节切片转换为字符串并返回
	return string(randomBytes)
}
