package utils

import (
	"errors"
	"strings"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// GetUserIDFromToken 从请求的JWT中提取用户ID
func GetUserIDFromToken(c *fiber.Ctx) (int, error) {
	// 首先检查上下文中是否已经有user_id（由JWT中间件设置）
	if userID, ok := c.Locals("user_id").(int); ok && userID > 0 {
		log.Printf("GetUserIDFromToken - 从context locals获取到用户ID: %d", userID)
		return userID, nil
	}
	
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
	log.Printf("GetUserIDFromToken - 解析token: %s", tokenString)

	// 解析JWT，这里仅做示例，实际项目中会使用密钥验证
	// 为简单测试，这里不验证签名
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 这里应该返回签名密钥，但示例中我们返回一个空字符串
		return []byte("your_jwt_secret"), nil
	})

	if err != nil {
		log.Printf("GetUserIDFromToken - 解析token失败: %v", err)
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

	log.Printf("GetUserIDFromToken - 解析到的claims: %v", claims)

	// 首先尝试获取user_id（与middleware/jwt.go中保持一致）
	if userIDValue, exists := claims["user_id"]; exists {
		if userID, ok := userIDValue.(float64); ok {
			log.Printf("GetUserIDFromToken - 从claims['user_id']获取到用户ID: %f", userID)
			return int(userID), nil
		}
	}
	
	// 然后尝试获取userID（小写，兼容其他可能的实现）
	if userIDValue, exists := claims["userID"]; exists {
		if userID, ok := userIDValue.(float64); ok {
			log.Printf("GetUserIDFromToken - 从claims['userID']获取到用户ID: %f", userID)
			return int(userID), nil
		}
	}
	
	// 测试环境中，如果没有有效的token，返回一个默认的用户ID
	log.Println("GetUserIDFromToken - 未找到用户ID，返回默认用户ID 1")
	return 1, nil // 开发测试环境返回默认ID为1
}

// GenerateTestToken 生成测试用的JWT token
func GenerateTestToken(userID int) (string, error) {
	// 创建token
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置claims
	claims := token.Claims.(jwt.MapClaims)
	// 使用与middleware/jwt.go相同的字段名
	claims["user_id"] = userID
	claims["username"] = "test_user"
	claims["exp"] = 1900000000 // 一个远期的过期时间

	// 签名token
	tokenString, err := token.SignedString([]byte("your_jwt_secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
} 