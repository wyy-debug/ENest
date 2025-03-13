package handlers

import (
	"go-server/models"

	"github.com/gofiber/fiber/v2"
)

type UpdateProfileRequest struct {
	Username       string `json:"username,omitempty"`
	Signature      string `json:"signature,omitempty"`
	StudyDirection string `json:"study_direction,omitempty"`
}

// UpdateUserProfile 更新用户个人信息的处理器
func UpdateUserProfile(c *fiber.Ctx) error {
	// 从上下文中获取用户ID
	userID := c.Locals("user_id").(int)

	// 解析请求体
	var req UpdateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// 构建更新字段映射
	updates := make(map[string]string)
	if req.Username != "" {
		updates["username"] = req.Username
	}
	if req.Signature != "" {
		updates["signature"] = req.Signature
	}
	if req.StudyDirection != "" {
		updates["study_direction"] = req.StudyDirection
	}

	// 更新用户信息
	if err := models.UpdateUserProfile(userID, updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Profile updated successfully",
	})
}

// GetUserProfile 获取用户个人信息的处理器
func GetUserProfile(c *fiber.Ctx) error {
	// 从上下文中获取用户ID
	userID := c.Locals("user_id").(int)

	// 获取用户信息
	user, err := models.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// 返回用户信息
	return c.JSON(fiber.Map{
		"id":              user.ID,
		"username":        user.Username,
		"signature":       user.Signature,
		"study_direction": user.StudyDirection,
		"created_at":      user.CreatedAt,
		"updated_at":      user.UpdatedAt,
	})
}
