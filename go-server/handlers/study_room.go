package handlers

import (
	"go-server/models"
	"go-server/proto"
	"go-server/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	protobuf "google.golang.org/protobuf/proto"
)

// CreateStudyRoom 创建自习室
// GetStudyRoomDetail 获取自习室详细信息
// JoinStudyRoom 加入自习室
// JoinStudyRoomByShareLink 通过分享链接加入自习室
// LeaveStudyRoom 离开自习室
// DestroyStudyRoom 销毁自习室

// HTTP API 处理函数

// CreateStudyRoom 创建自习室
func CreateStudyRoom(c *fiber.Ctx) error {
	// 验证用户身份
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized access",
		})
	}

	// 解析请求体
	type CreateRoomRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		MaxMembers  int    `json:"max_members"`
		IsPrivate   bool   `json:"is_private"`
		Duration    string `json:"duration"` // 例如 "24h", "3d" 等
	}

	var req CreateRoomRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// 验证请求参数
	if req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Room name is required",
		})
	}

	// 解析持续时间
	duration, err := time.ParseDuration(req.Duration)
	if err != nil {
		duration = 24 * time.Hour // 默认24小时
	}

	// 创建自习室
	room, err := models.CreateStudyRoom(userID, req.Name, req.MaxMembers, req.IsPrivate, duration)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(room)
}

// GetStudyRooms 获取自习室列表
func GetStudyRooms(c *fiber.Ctx) error {
	// 验证用户身份
	_, err := utils.GetUserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized access",
		})
	}

	// 获取查询参数
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	
	// 获取自习室列表
	rooms, err := models.GetStudyRooms(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(rooms)
}

// GetStudyRoom 获取自习室详情
func GetStudyRoom(c *fiber.Ctx) error {
	// 验证用户身份
	_, err := utils.GetUserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized access",
		})
	}

	// 获取自习室ID
	roomID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid room ID",
		})
	}

	// 获取自习室详情
	room, err := models.GetStudyRoomDetail(roomID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(room)
}

// UpdateStudyRoom 更新自习室信息
func UpdateStudyRoom(c *fiber.Ctx) error {
	// 验证用户身份
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized access",
		})
	}

	// 获取自习室ID
	roomID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid room ID",
		})
	}

	// 解析请求体
	type UpdateRoomRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		MaxMembers  int    `json:"max_members"`
		IsPrivate   bool   `json:"is_private"`
	}

	var req UpdateRoomRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// 更新自习室
	room, err := models.UpdateStudyRoom(roomID, userID, req.Name, req.Description, req.MaxMembers, req.IsPrivate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(room)
}

// DeleteStudyRoom 删除自习室
func DeleteStudyRoom(c *fiber.Ctx) error {
	// 验证用户身份
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized access",
		})
	}

	// 获取自习室ID
	roomID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid room ID",
		})
	}

	// 删除自习室
	if err := models.DestroyStudyRoom(roomID, userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Study room deleted successfully",
	})
}

// JoinStudyRoom 加入自习室
func JoinStudyRoom(c *fiber.Ctx) error {
	// 验证用户身份
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized access",
		})
	}

	// 获取自习室ID
	roomID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid room ID",
		})
	}

	// 解析请求体
	type JoinRoomRequest struct {
		IsAnonymous bool `json:"is_anonymous"`
	}

	var req JoinRoomRequest
	if err := c.BodyParser(&req); err != nil {
		// 默认不匿名
		req.IsAnonymous = false
	}

	// 加入自习室
	member, err := models.JoinStudyRoom(roomID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(member)
}

// LeaveStudyRoom 离开自习室
func LeaveStudyRoom(c *fiber.Ctx) error {
	// 验证用户身份
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized access",
		})
	}

	// 获取自习室ID
	roomID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid room ID",
		})
	}

	// 离开自习室
	if err := models.LeaveStudyRoom(roomID, userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Left study room successfully",
	})
}

// GetStudyRoomMembers 获取自习室成员列表
func GetStudyRoomMembers(c *fiber.Ctx) error {
	// 验证用户身份
	_, err := utils.GetUserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized access",
		})
	}

	// 获取自习室ID
	roomID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid room ID",
		})
	}

	// 获取成员列表
	members, err := models.GetStudyRoomMembers(roomID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(members)
}

// WebSocket消息处理函数

// handleStudyRoomMessage 处理自习室消息
func handleStudyRoomMessage(conn *Connection, payload []byte) error {
	// 解析自习室消息
	var studyRoomMsg proto.StudyRoomMessage
	if err := protobuf.Unmarshal(payload, &studyRoomMsg); err != nil {
		return err
	}

	// 根据操作类型处理消息
	switch studyRoomMsg.Operation {
	case proto.StudyRoomMessage_CREATE:
		// 创建自习室
		duration, _ := time.ParseDuration(studyRoomMsg.Duration)
		room, err := models.CreateStudyRoom(conn.userID, studyRoomMsg.Name, int(studyRoomMsg.MaxMembers), studyRoomMsg.IsPrivate, duration)
		if err != nil {
			sendSystemMessage(conn, err.Error())
			return err
		}
		sendSuccessMessage(conn, room)

	case proto.StudyRoomMessage_JOIN:
		// 加入自习室
		member, err := models.JoinStudyRoom(int(studyRoomMsg.RoomId), conn.userID)
		if err != nil {
			sendSystemMessage(conn, err.Error())
			return err
		}
		sendSuccessMessage(conn, member)

	case proto.StudyRoomMessage_LEAVE:
		// 离开自习室
		if err := models.LeaveStudyRoom(int(studyRoomMsg.RoomId), conn.userID); err != nil {
			sendSystemMessage(conn, err.Error())
			return err
		}
		sendSuccessMessage(conn, map[string]string{"message": "Successfully left the study room"})

	case proto.StudyRoomMessage_DESTROY:
		// 销毁自习室
		if err := models.DestroyStudyRoom(int(studyRoomMsg.RoomId), conn.userID); err != nil {
			sendSystemMessage(conn, err.Error())
			return err
		}
		sendSuccessMessage(conn, map[string]string{"message": "Successfully destroyed the study room"})

	case proto.StudyRoomMessage_GET_DETAIL:
		// 获取自习室详情
		detail, err := models.GetStudyRoomDetail(int(studyRoomMsg.RoomId))
		if err != nil {
			sendSystemMessage(conn, err.Error())
			return err
		}
		sendSuccessMessage(conn, detail)
	}

	return nil
}

func init() {
	// 注册个人信息消息处理器
	RegisterMessageHandler(proto.MessageType_STUDY_ROOM, handleStudyRoomMessage)
}
