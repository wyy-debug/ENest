package utils

import (
	"math/rand"
	"time"
)

// 初始化随机数生成器
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// GetCurrentTime 获取当前时间
func GetCurrentTime() time.Time {
	return time.Now()
}

// FormatTime 格式化时间为ISO8601标准格式
func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

// GenerateRandomString 生成指定长度的随机字符串
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
} 