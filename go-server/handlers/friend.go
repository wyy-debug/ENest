package handlers

import (
	"go-server/models"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// 发送好友请求
func SendFriendRequest(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	var req struct {
		FriendID int `json:"friend_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := models.SendFriendRequest(userID, req.FriendID)
	if err != nil {
		if err.Error() == "friend request already exists" {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

// 处理好友请求
func HandleFriendRequest(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	var req struct {
		RequestID int    `json:"request_id"`
		Action    string `json:"action"` // accept or reject
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := models.HandleFriendRequest(req.RequestID, userID, req.Action)
	if err != nil {
		if err.Error() == "friend request not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

// 获取好友列表
func GetFriendList(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)

	friends, err := models.GetFriendList(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(friends)
}

// 删除好友
func DeleteFriend(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	var req struct {
		FriendID int `json:"friend_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := models.DeleteFriend(userID, req.FriendID)
	if err != nil {
		if err.Error() == "friend relationship not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

// 创建好友契约
func CreateFriendContract(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	var req struct {
		FriendID      int    `json:"friend_id"`
		ContractType  string `json:"contract_type"`
		ContractTerms string `json:"contract_terms"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	contract := &models.FriendContract{
		UserID:        userID,
		FriendID:      req.FriendID,
		ContractType:  req.ContractType,
		ContractTerms: req.ContractTerms,
	}

	err := models.CreateFriendContract(userID, contract)
	if err != nil {
		if err.Error() == "friend relationship not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(contract)
}

// 终止好友契约
func TerminateFriendContract(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	var req struct {
		ContractID int `json:"contract_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := models.TerminateFriendContract(req.ContractID, userID)
	if err != nil {
		if err.Error() == "contract not found or already terminated" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

// 获取好友契约列表
func GetFriendContracts(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)

	contracts, err := models.GetFriendContracts(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(contracts)
}

// WebSocket连接处理
func HandleChat(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

// WebSocket连接管理器
// 需要导入 "sync" 包
// 导入sync包

type WSManager struct {
	connections map[int]*websocket.Conn
	mu          sync.Mutex
}

var wsManager = &WSManager{
	connections: make(map[int]*websocket.Conn),
}

func (m *WSManager) AddConnection(userID int, conn *websocket.Conn) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.connections[userID] = conn
}

func (m *WSManager) RemoveConnection(userID int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.connections, userID)
}

func (m *WSManager) GetConnection(userID int) (*websocket.Conn, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	conn, ok := m.connections[userID]
	return conn, ok
}

// WebSocket连接处理中间件
func HandleChatWS(c *websocket.Conn) {
	userID := c.Locals("user_id").(int)
	defer func() {
		c.Close()
		wsManager.RemoveConnection(userID)
	}()

	// 添加连接到管理器
	wsManager.AddConnection(userID, c)

	for {
		// 读取消息
		var message models.FriendMessage
		if err := c.ReadJSON(&message); err != nil {
			break
		}

		// 验证消息接收者是否为好友
		exists, err := models.CheckFriendRelationship(userID, message.ReceiverID)
		if err != nil || !exists {
			c.WriteJSON(fiber.Map{"error": "Not friend relationship"})
			continue
		}

		// 保存消息到数据库
		message.SenderID = userID
		err = models.SaveFriendMessage(&message)
		if err != nil {
			c.WriteJSON(fiber.Map{"error": "Failed to save message"})
			continue
		}

		// 推送消息给接收者（如果在线）
		if conn, ok := wsManager.GetConnection(message.ReceiverID); ok {
			err = conn.WriteJSON(message)
			if err != nil {
				c.WriteJSON(fiber.Map{"error": "Failed to send message"})
				continue
			}
		}

		// 发送成功响应
		c.WriteJSON(message)
	}
}
