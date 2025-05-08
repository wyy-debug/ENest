package main

import (
	"fmt"
	"time"
	
	"github.com/golang-jwt/jwt/v4"
)

func main() {
	// 创建JWT claims
	claims := jwt.MapClaims{
		"user_id": 1,
		"email": "test@example.com",
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
		"iss": "newenest",
	}
	
	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	// 使用密钥签名
	tokenString, err := token.SignedString([]byte("newenest_secret_key"))
	if err != nil {
		fmt.Println("生成token失败:", err)
		return
	}
	
	fmt.Println("测试Token:")
	fmt.Println(tokenString)
} 