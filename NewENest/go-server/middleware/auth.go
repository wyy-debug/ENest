package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
)

// UserClaims JWT令牌声明
type UserClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// 从环境变量或配置获取JWT密钥，这里简化为静态字符串
// 实际生产环境应从配置获取
var jwtSecret = []byte("your-jwt-secret")

// Auth 认证中间件
func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 获取Authorization头
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "缺少认证令牌")
		}

		// 提取JWT令牌
		tokenString := ""
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			tokenString = authHeader
		}

		// 解析JWT令牌
		token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			// 验证签名算法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("意外的签名方法: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil {
			log.Error().Err(err).Str("token", tokenString).Msg("解析JWT令牌失败")
			return fiber.NewError(fiber.StatusUnauthorized, "无效的认证令牌")
		}

		// 验证令牌有效
		if !token.Valid {
			return fiber.NewError(fiber.StatusUnauthorized, "无效的认证令牌")
		}

		// 获取JWT声明
		claims, ok := token.Claims.(*UserClaims)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "无效的令牌声明")
		}

		// 检查令牌是否过期
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
			return fiber.NewError(fiber.StatusUnauthorized, "令牌已过期")
		}

		// 将用户信息存储在上下文中
		c.Locals("userId", claims.UserID)
		c.Locals("userEmail", claims.Email)

		return c.Next()
	}
}

// GenerateToken 生成JWT令牌
func GenerateToken(userID int, email string) (string, error) {
	// 设置令牌过期时间（例如24小时）
	expirationTime := time.Now().Add(24 * time.Hour)

	// 创建JWT声明
	claims := &UserClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "newenest",
			Subject:   fmt.Sprintf("%d", userID),
		},
	}

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名令牌
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Error().Err(err).Int("userID", userID).Msg("生成JWT令牌失败")
		return "", err
	}

	return tokenString, nil
}

// GetUserID 从请求上下文获取当前用户ID
func GetUserID(c *fiber.Ctx) (int, error) {
	userID, ok := c.Locals("userId").(int)
	if !ok {
		return 0, fmt.Errorf("未找到用户ID")
	}
	return userID, nil
}

// GetUserEmail 从请求上下文获取当前用户邮箱
func GetUserEmail(c *fiber.Ctx) (string, error) {
	userEmail, ok := c.Locals("userEmail").(string)
	if !ok {
		return "", fmt.Errorf("未找到用户邮箱")
	}
	return userEmail, nil
} 