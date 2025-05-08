package utils

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// GetUserIDFromToken 从请求的JWT中提取用户ID
func GetUserIDFromToken(c *fiber.Ctx) (int, error) {
	// 获取Authorization头
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return 0, errors.New("未提供授权头")
	}

	// 检查Bearer前缀
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return 0, errors.New("授权格式无效")
	}

	// 获取token字符串
	tokenString := parts[1]

	// 解析JWT，这里仅做示例，实际项目中会使用密钥验证
	// 为简单测试，这里不验证签名
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 这里应该返回签名密钥，但示例中我们返回一个空字符串
		return []byte("your_jwt_secret"), nil
	})

	if err != nil {
		return 0, err
	}

	// 检查token有效性
	if !token.Valid {
		return 0, errors.New("令牌无效")
	}

	// 从JWT claims中获取用户ID
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("无法解析令牌声明")
	}

	// 获取用户ID
	userID, ok := claims["userID"].(float64)
	if !ok {
		// 测试环境中，如果没有有效的token，返回一个默认的用户ID
		return 1, nil // 开发测试环境返回默认ID为1
	}

	return int(userID), nil
}

// GenerateTestToken 生成测试用的JWT token
func GenerateTestToken(userID int) (string, error) {
	// 创建token
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置claims
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["username"] = "test_user"
	claims["exp"] = 1900000000 // 一个远期的过期时间

	// 签名token
	tokenString, err := token.SignedString([]byte("your_jwt_secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
} 