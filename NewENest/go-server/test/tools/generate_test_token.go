package main

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func main() {
	// 加载环境变量
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println("未找到.env文件，使用默认配置")
	}

	// 获取JWT密钥
	jwtSecret := os.Getenv("NEST_JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "newenest_secret_key" // 默认密钥
	}

	// 测试用户1 - 使用JWTClaims结构
	generateStructToken(1, "test@example.com", jwtSecret)

	// 测试用户1 - 使用MapClaims结构
	generateMapToken(1, "test@example.com", jwtSecret)

	// 测试用户2
	generateStructToken(2, "buddy@example.com", jwtSecret)
	generateMapToken(2, "buddy@example.com", jwtSecret)
}

// 使用结构化Claims生成令牌
func generateStructToken(userID int, email string, secret string) {
	// 自定义Claims结构
	type JWTClaims struct {
		UserID int    `json:"user_id"`
		Email  string `json:"email"`
		jwt.StandardClaims
	}

	// 设置JWT声明
	claims := JWTClaims{
		UserID: userID,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // 24小时后过期
			IssuedAt:  time.Now().Unix(),
			Issuer:    "newenest",
		},
	}

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名字符串
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Printf("结构化令牌生成失败: %v\n", err)
		return
	}

	fmt.Printf("用户ID: %d, 邮箱: %s\n", userID, email)
	fmt.Printf("结构化令牌: %s\n\n", tokenString)
}

// 使用MapClaims生成令牌
func generateMapToken(userID int, email string, secret string) {
	// 创建MapClaims
	claims := jwt.MapClaims{
		"user_id":  userID,
		"email":    email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
		"issuer":   "newenest",
	}

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名令牌
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Printf("Map令牌生成失败: %v\n", err)
		return
	}

	fmt.Printf("用户ID: %d, 邮箱: %s\n", userID, email)
	fmt.Printf("Map令牌: %s\n\n", signedToken)
} 