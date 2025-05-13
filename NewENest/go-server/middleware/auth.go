package middleware

import (
	"NewENest/go-server/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

// RegisterAuthMiddleware 注册认证中间件
func RegisterAuthMiddleware(app *fiber.App) {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("加载配置失败: " + err.Error())
	}

	app.Use(func(c *fiber.Ctx) error {
		// 跳过不需要认证的路由
		if isPublicRoute(c.Path()) {
			return c.Next()
		}

		// 获取Authorization头
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"code":    401,
				"message": "未授权，请先登录",
			})
		}

		// 验证Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(401).JSON(fiber.Map{
				"code":    401,
				"message": "无效的认证格式",
			})
		}

		// 解析JWT token
		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWT.Secret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{
				"code":    401,
				"message": "无效的token或token已过期",
			})
		}

		// 将用户ID存储在上下文中
		claims := token.Claims.(jwt.MapClaims)
		c.Locals("userID", int(claims["user_id"].(float64)))

		return c.Next()
	})
}

// isPublicRoute 判断是否为公开路由
func isPublicRoute(path string) bool {
	// 定义公开路由前缀
	publicPrefixes := []string{
		"/api/v1/auth/",    // v1版本的认证相关
		"/api/auth/",       // 旧版本的认证相关
		"/api/v1/study-rooms", // v1版本的自习室列表
		"/api/study-rooms",    // 旧版本的自习室列表
		"/api/v1/health",      // v1版本的健康检查
		"/api/health",         // 旧版本的健康检查
		"/api/v1/ws",          // v1版本的WebSocket
		"/api/ws",             // 旧版本的WebSocket
	}

	// 检查路径是否以任何公开前缀开头
	for _, prefix := range publicPrefixes {
		if strings.HasPrefix(path, prefix) {
			// 对于自习室列表，只允许GET请求
			if (prefix == "/api/v1/study-rooms" || prefix == "/api/study-rooms") && !strings.Contains(path, "/") {
				return true
			}
			// 对于其他路由，完全匹配前缀
			if prefix != "/api/v1/study-rooms" && prefix != "/api/study-rooms" {
				return true
			}
		}
	}

	return false
} 