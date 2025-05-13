package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	
	"NewENest/go-server/middleware"
)

// LoginV2 用户登录 - 更新版本
func LoginV2(c *fiber.Ctx) error {
	// 解析请求数据
	var loginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	if err := c.BodyParser(&loginRequest); err != nil {
		return BadRequest(c, "请求格式无效", nil)
	}
	
	// 检查邮箱和密码是否为空
	if loginRequest.Email == "" || loginRequest.Password == "" {
		return BadRequest(c, "邮箱和密码不能为空", nil)
	}

	// 检查邮箱是否存在（在实际应用中应查询数据库）
	if loginRequest.Email != "test@example.com" && loginRequest.Email != "buddy@example.com" {
		return Unauthorized(c, "邮箱或密码错误")
	}
	
	// 检查密码是否正确（在实际应用中应验证哈希值）
	if loginRequest.Password != "password123" {
		return Unauthorized(c, "邮箱或密码错误")
	}
		
	var userID int
	var username string
	
	if loginRequest.Email == "test@example.com" {
		userID = 1
		username = "test_user"
	} else if loginRequest.Email == "buddy@example.com" {
		userID = 2
		username = "study_buddy"
	}
	
	// 生成JWT令牌
	token, err := middleware.GenerateToken(userID, loginRequest.Email)
	if err != nil {
		return ServerError(c, "生成令牌失败", err)
	}
	
	// 构建模拟用户数据
	user := fiber.Map{
		"id":              userID,
		"username":        username,
		"email":           loginRequest.Email,
		"avatar":          fmt.Sprintf("https://api.dicebear.com/7.x/avataaars/svg?seed=%s", username),
		"signature":       "这是我的个性签名",
		"study_direction": "计算机科学",
	}
	
	// 返回成功响应
	return Success(c, fiber.Map{
		"token": token,
		"user":  user,
	}, "登录成功")
} 