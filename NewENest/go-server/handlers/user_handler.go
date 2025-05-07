package handlers

import (
	"database/sql"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"

	"NewENest/go-server/middleware"
	"NewENest/go-server/models"
	"NewENest/go-server/services"
)

// UserHandler 用户控制器
type UserHandler struct {
	userService services.UserService
	validator   *validator.Validate
}

// NewUserHandler 创建用户控制器实例
func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		validator:   validator.New(),
	}
}

// Response API响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// Register 用户注册
func Register(c *fiber.Ctx) error {
	// 解析请求体
	var input models.UserCreate
	if err := c.BodyParser(&input); err != nil {
		return BadRequest(c, "请求格式无效", err)
	}

	// 验证邮箱是否已存在（示例，实际应当在数据库中检查）
	// 此处简化处理，假设邮箱不存在
	if input.Email == "test@example.com" {
		return BadRequest(c, "邮箱已被注册", nil)
	}

	// 对密码进行哈希
	hashedPassword, err := models.HashPassword(input.Password)
	if err != nil {
		return ServerError(c, "密码哈希失败", err)
	}

	// 创建用户（示例，实际应当保存到数据库）
	user := &models.User{
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: hashedPassword,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// 返回响应
	return Success(c, user.ToResponse(), "注册成功")
}

// Login 用户登录
func Login(c *fiber.Ctx) error {
	// 解析请求体
	var input models.UserLogin
	if err := c.BodyParser(&input); err != nil {
		return BadRequest(c, "请求格式无效", err)
	}

	// 查找用户（示例，实际应当从数据库获取）
	// 此处简化处理，直接创建一个示例用户
	user := &models.User{
		ID:               1,
		Username:         "testuser",
		Email:            "test@example.com",
		PasswordHash:     "$2a$10$1qAz2wSx3eDc4rFv5tGb5edZT/WsmrytDkOAFoZiNBe0X0kKai5r.", // "password"
		Avatar:           "https://example.com/avatar.jpg",
		Signature:        "这是我的个性签名",
		StudyDirection:   "programming",
		TotalStudyTime:   3600,
		AchievementPoints: 120,
		CreatedAt:        time.Now().Add(-24 * time.Hour),
		UpdatedAt:        time.Now(),
	}

	// 检查密码
	if !models.CheckPasswordHash(input.Password, user.PasswordHash) {
		return Unauthorized(c, "邮箱或密码错误")
	}

	// 生成JWT令牌
	token, err := middleware.GenerateToken(user.ID, user.Email)
	if err != nil {
		return ServerError(c, "生成令牌失败", err)
	}

	// 返回用户信息和令牌
	return Success(c, fiber.Map{
		"user":  user.ToResponse(),
		"token": token,
	}, "登录成功")
}

// GetCurrentUser 获取当前用户信息
func GetCurrentUser(c *fiber.Ctx) error {
	// 从请求上下文获取用户ID
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return Unauthorized(c, "未授权")
	}

	// 获取用户信息（示例，实际应当从数据库获取）
	// 此处简化处理，直接创建一个示例用户
	user := &models.User{
		ID:               userID,
		Username:         "testuser",
		Email:            "test@example.com",
		Avatar:           "https://example.com/avatar.jpg",
		Signature:        "这是我的个性签名",
		StudyDirection:   "programming",
		TotalStudyTime:   3600,
		AchievementPoints: 120,
		CreatedAt:        time.Now().Add(-24 * time.Hour),
		UpdatedAt:        time.Now(),
	}

	return Success(c, user.ToResponse(), "获取用户信息成功")
}

// UpdateCurrentUser 更新当前用户信息
func UpdateCurrentUser(c *fiber.Ctx) error {
	// 从请求上下文获取用户ID
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return Unauthorized(c, "未授权")
	}

	// 解析请求体
	var input models.UserUpdate
	if err := c.BodyParser(&input); err != nil {
		return BadRequest(c, "请求格式无效", err)
	}

	// 获取用户信息（示例，实际应当从数据库获取）
	// 此处简化处理，直接创建一个示例用户
	user := &models.User{
		ID:               userID,
		Username:         "testuser",
		Email:            "test@example.com",
		Avatar:           "https://example.com/avatar.jpg",
		Signature:        "这是我的个性签名",
		StudyDirection:   "programming",
		TotalStudyTime:   3600,
		AchievementPoints: 120,
		CreatedAt:        time.Now().Add(-24 * time.Hour),
		UpdatedAt:        time.Now(),
	}

	// 更新用户信息
	if input.Username != "" {
		user.Username = input.Username
	}
	if input.Avatar != "" {
		user.Avatar = input.Avatar
	}
	if input.Signature != "" {
		user.Signature = input.Signature
	}
	if input.StudyDirection != "" {
		user.StudyDirection = input.StudyDirection
	}
	user.UpdatedAt = time.Now()

	// 此处应当将更新后的用户信息保存到数据库

	return Success(c, user.ToResponse(), "用户信息更新成功")
}

// RegisterRoutes 注册路由
func (h *UserHandler) RegisterRoutes(router fiber.Router) {
	userRouter := router.Group("/users")
	
	userRouter.Post("/register", Register)
	userRouter.Post("/login", Login)
	userRouter.Post("/logout", h.Logout)
	
	// 需要认证的路由
	userRouter.Get("/profile", AuthMiddleware(h.userService), h.GetProfile)
	userRouter.Put("/profile", AuthMiddleware(h.userService), h.UpdateProfile)
	userRouter.Put("/password", AuthMiddleware(h.userService), h.ChangePassword)
}

// AuthMiddleware 认证中间件
func AuthMiddleware(userService services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 从请求头获取令牌
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(Response{
				Code:    fiber.StatusUnauthorized,
				Message: "Unauthorized: No token provided",
			})
		}
		
		// 验证令牌
		user, err := userService.ValidateToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(Response{
				Code:    fiber.StatusUnauthorized,
				Message: "Unauthorized: Invalid token",
				Errors:  []string{err.Error()},
			})
		}
		
		// 将用户ID存储在上下文中
		c.Locals("userID", user.ID)
		c.Locals("user", user)
		
		return c.Next()
	}
} 