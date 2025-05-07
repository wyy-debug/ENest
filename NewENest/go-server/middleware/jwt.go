package middleware

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// 密钥应当存储在环境变量中，这里仅作示例
var jwtSecret = []byte("your-jwt-secret-key")

// JWTClaims 自定义JWT声明结构
type JWTClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

// GenerateToken 生成JWT令牌
func GenerateToken(userID int, email string) (string, error) {
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
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 解析JWT令牌
func ParseToken(tokenString string) (*JWTClaims, error) {
	// 解析令牌
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 确保令牌算法匹配
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	// 解析失败
	if err != nil {
		return nil, err
	}

	// 验证声明
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// GetUserID 从请求上下文获取用户ID
func GetUserID(c *fiber.Ctx) (int, error) {
	userID, ok := c.Locals("userID").(int)
	if !ok {
		return 0, errors.New("用户ID不存在")
	}
	return userID, nil
}

// JWTMiddleware 验证JWT令牌的中间件
func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 从请求头获取令牌
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    fiber.StatusUnauthorized,
				"message": "未提供授权令牌",
			})
		}

		// 验证令牌格式
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    fiber.StatusUnauthorized,
				"message": "授权格式无效",
			})
		}

		// 解析令牌
		claims, err := ParseToken(parts[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    fiber.StatusUnauthorized,
				"message": "无效的令牌",
				"errors":  err.Error(),
			})
		}

		// 将用户信息存储在上下文中
		c.Locals("userID", claims.UserID)
		c.Locals("email", claims.Email)

		// 继续处理请求
		return c.Next()
	}
} 