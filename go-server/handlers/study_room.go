package handlers

import (
	"go-server/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// CreateStudyRoom 创建自习室
func CreateStudyRoom(c *fiber.Ctx) error {
	// 获取当前用户ID
	userID := c.Locals("userID").(int)

	// 解析请求参数
	type CreateRoomRequest struct {
		Name       string        `json:"name"`
		MaxMembers int           `json:"max_members"`
		IsPrivate  bool          `json:"is_private"`
		Duration   time.Duration `json:"duration"`
	}

	var req CreateRoomRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// 创建自习室
	room, err := models.CreateStudyRoom(userID, req.Name, req.MaxMembers, req.IsPrivate, req.Duration)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(room)
}

// GetStudyRoomDetail 获取自习室详细信息
func GetStudyRoomDetail(c *fiber.Ctx) error {
	// 解析请求参数
	type GetRoomDetailRequest struct {
		RoomID int `json:"room_id"`
	}

	var req GetRoomDetailRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// 获取自习室详细信息
	detail, err := models.GetStudyRoomDetail(req.RoomID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(detail)
}

// JoinStudyRoom 加入自习室
func JoinStudyRoom(c *fiber.Ctx) error {
	// 获取当前用户ID
	userID := c.Locals("userID").(int)

	// 解析请求参数
	type JoinRoomRequest struct {
		RoomID int `json:"room_id"`
	}

	var req JoinRoomRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// 加入自习室
	member, err := models.JoinStudyRoom(req.RoomID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(member)
}

// JoinStudyRoomByShareLink 通过分享链接加入自习室
func JoinStudyRoomByShareLink(c *fiber.Ctx) error {
	// 获取当前用户ID
	userID := c.Locals("userID").(int)

	// 获取分享链接
	shareLink := c.Params("shareLink")

	// 通过分享链接获取自习室信息
	room, err := models.GetStudyRoomByShareLink(shareLink)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// 加入自习室
	member, err := models.JoinStudyRoom(room.ID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(member)
}

// LeaveStudyRoom 离开自习室
func LeaveStudyRoom(c *fiber.Ctx) error {
	// 获取当前用户ID
	userID := c.Locals("userID").(int)

	// 解析请求参数
	type LeaveRoomRequest struct {
		RoomID int `json:"room_id"`
	}

	var req LeaveRoomRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// 离开自习室
	if err := models.LeaveStudyRoom(req.RoomID, userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully left the study room",
	})
}

// DestroyStudyRoom 销毁自习室
func DestroyStudyRoom(c *fiber.Ctx) error {
	// 获取当前用户ID
	userID := c.Locals("userID").(int)

	// 解析请求参数
	type DestroyRoomRequest struct {
		RoomID int `json:"room_id"`
	}

	var req DestroyRoomRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// 销毁自习室
	if err := models.DestroyStudyRoom(req.RoomID, userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully destroyed the study room",
	})
}
