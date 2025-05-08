package middleware

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// 密钥应当存储在环境变量中，这里仅作示例
var jwtSecret = []byte("newenest_secret_key")

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
	userID, ok := c.Locals("user_id").(int)
	if !ok {
		return 0, errors.New("unauthorized")
	}
	return userID, nil
}

// JWTMiddleware JWT认证中间件
func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 获取请求头中的Authorization信息
		authHeader := c.Get("Authorization")
		log.Printf("JWT中间件 - 接收到的Authorization头: %s", authHeader)
		
		if authHeader == "" {
			log.Println("JWT中间件 - 未提供认证令牌")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    fiber.StatusUnauthorized,
				"message": "未提供认证令牌",
			})
		}
		
		// 检查格式是否为"Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Printf("JWT中间件 - 认证令牌格式无效: %s", authHeader)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    fiber.StatusUnauthorized,
				"message": "认证令牌格式无效",
			})
		}
		
		// 获取令牌
		tokenString := parts[1]
		log.Printf("JWT中间件 - 提取的令牌: %s", tokenString)
		
		// 试图使用自定义Claims解析
		claims := &JWTClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("签名方法无效: %v", t.Header["alg"])
			}
			return []byte(GetJWTSecret()), nil
		})
		
		// 如果自定义Claims解析失败，尝试使用MapClaims
		if err != nil {
			log.Printf("JWT中间件 - 使用自定义Claims解析失败: %v，尝试使用MapClaims", err)
			
			// 解析令牌
			token, err = jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("签名方法无效: %v", t.Header["alg"])
				}
				return []byte(GetJWTSecret()), nil
			})
			
			if err != nil {
				log.Printf("JWT中间件 - 令牌解析失败: %v", err)
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"code":    fiber.StatusUnauthorized,
					"message": "无效的认证令牌",
					"error":   err.Error(),
				})
			}
		}
		
		// 验证令牌有效性
		if !token.Valid {
			log.Println("JWT中间件 - 令牌无效")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    fiber.StatusUnauthorized,
				"message": "认证令牌已过期或无效",
			})
		}
		
		var userID float64
		
		// 获取令牌中的用户ID，尝试不同的方式
		if claims, ok := token.Claims.(*JWTClaims); ok {
			log.Printf("JWT中间件 - 成功解析自定义Claims，用户ID: %d", claims.UserID)
			c.Locals("user_id", claims.UserID)
			return c.Next()
		} else if mapClaims, mapOk := token.Claims.(jwt.MapClaims); mapOk {
			// 尝试获取user_id
			if userIDValue, exists := mapClaims["user_id"]; exists {
				userID, ok = userIDValue.(float64)
				if !ok {
					log.Printf("JWT中间件 - 无法将user_id转换为数字: %v", userIDValue)
				}
			} 
			
			// 如果user_id不存在或类型不对，尝试获取userID
			if !ok {
				if userIDValue, exists := mapClaims["userID"]; exists {
					userID, ok = userIDValue.(float64)
					if !ok {
						log.Printf("JWT中间件 - 无法将userID转换为数字: %v", userIDValue)
					}
				}
			}
			
			if !ok {
				log.Println("JWT中间件 - 令牌中缺少有效的用户ID")
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"code":    fiber.StatusUnauthorized,
					"message": "令牌中缺少有效的用户ID",
				})
			}
			
			log.Printf("JWT中间件 - 成功解析MapClaims，用户ID: %f", userID)
		} else {
			log.Println("JWT中间件 - 无法解析令牌Claims")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    fiber.StatusUnauthorized,
				"message": "无法获取令牌信息",
			})
		}
		
		// 将用户ID存储在上下文中
		c.Locals("user_id", int(userID))
		log.Printf("JWT中间件 - 验证成功，用户ID: %d", int(userID))
		
		return c.Next()
	}
}

// GetJWTSecret 获取JWT密钥
func GetJWTSecret() string {
	return "newenest_secret_key" // 在实际应用中应该从环境变量或配置中获取
} 