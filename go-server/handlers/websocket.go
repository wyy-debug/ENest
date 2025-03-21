package handlers

import (
	"encoding/json"
	"go-server/proto"
	"go-server/utils"
	"log"
	"sync"
	"time"

	"github.com/gofiber/websocket/v2"
	protobuf "google.golang.org/protobuf/proto"
)

const (
	heartbeatInterval = 30 * time.Second
	heartbeatTimeout  = 90 * time.Second
)

// Connection WebSocket连接结构
type Connection struct {
	conn      *websocket.Conn
	userID    int
	deviceID  string
	lastPing  time.Time
	crypto    *utils.CryptoManager
	sendMutex sync.Mutex
}

// ConnectionManager WebSocket连接管理器
type ConnectionManager struct {
	connections sync.Map // map[int]*Connection，用户ID到连接的映射
	crypto      *utils.CryptoManager
}

// 全局连接管理器实例
var globalManager = &ConnectionManager{}

func init() {
	// 初始化加密管理器
	crypto, err := utils.NewCryptoManager()
	if err != nil {
		log.Fatal("Failed to create crypto manager:", err)
	}
	globalManager.crypto = crypto
}

// HandleWebSocket 处理WebSocket连接
func HandleWebSocket(c *websocket.Conn) {
	// 创建新连接
	conn := &Connection{
		conn:     c,
		lastPing: time.Now(),
	}

	// 启动心跳检测
	go handleHeartbeat(conn)

	// 处理消息
	for {
		messageType, data, err := c.ReadMessage()
		if err != nil {
			log.Printf("Read error: %v", err)
			break
		}

		if messageType == websocket.BinaryMessage {
			handleMessage(conn, data)
		}

		// 发送心跳响应
		if messageType == websocket.PingMessage {
			c.WriteMessage(websocket.PongMessage, []byte{})
		}
	}

	// 清理连接
	cleanupConnection(conn)
}

// handleHeartbeat 处理心跳
func handleHeartbeat(conn *Connection) {
	ticker := time.NewTicker(heartbeatInterval)
	defer ticker.Stop()

	for {
		<-ticker.C

		// 检查最后一次心跳时间
		if time.Since(conn.lastPing) > heartbeatTimeout {
			log.Printf("Client %d heartbeat timeout", conn.userID)
			conn.conn.Close()
			break
		}

		// 发送心跳包
		heartbeat := &proto.HeartbeatMessage{
			Timestamp: time.Now().Unix(),
		}
		sendProtoMessage(conn, proto.MessageType_HEARTBEAT, heartbeat)
	}
}

// MessageHandler 消息处理函数类型
type MessageHandler func(*Connection, []byte) error

// messageHandlers 消息处理器映射
var messageHandlers = make(map[proto.MessageType]MessageHandler)

// RegisterMessageHandler 注册消息处理器
func RegisterMessageHandler(msgType proto.MessageType, handler MessageHandler) {
	messageHandlers[msgType] = handler
}

// handleMessage 处理接收到的消息
func handleMessage(conn *Connection, data []byte) {
	// 解密消息
	plaintext, err := conn.crypto.Decrypt(data)
	if err != nil {
		log.Printf("Decrypt error: %v", err)
		return
	}

	// 解析消息
	var message proto.Message
	if err := protobuf.Unmarshal(plaintext, &message); err != nil {
		log.Printf("Unmarshal error: %v", err)
		return
	}

	// 更新心跳时间
	conn.lastPing = time.Now()

	// 处理不同类型的消息
	if handler, ok := messageHandlers[message.Type]; ok {
		if err := handler(conn, message.Payload); err != nil {
			log.Printf("Handle message error: %v", err)
		}
	}
}

// cleanupConnection 清理连接资源
func cleanupConnection(conn *Connection) {
	// 从连接管理器中移除连接
	globalManager.connections.Delete(conn.userID)
	// 关闭连接
	conn.conn.Close()
	log.Printf("Connection cleaned up for user %d", conn.userID)
}

// sendSystemMessage 发送系统消息
func sendSystemMessage(conn *Connection, message string) {
	// 创建系统消息
	sysMsg := &proto.SystemMessage{
		Content: message,
	}
	// 使用protobuf消息发送
	sendProtoMessage(conn, proto.MessageType_SYSTEM, sysMsg)
}

// sendProtoMessage 发送protobuf消息
func sendProtoMessage(conn *Connection, msgType proto.MessageType, msg protobuf.Message) {
	// 使用互斥锁保护发送过程
	conn.sendMutex.Lock()
	defer conn.sendMutex.Unlock()

	// 序列化消息内容
	payload, err := protobuf.Marshal(msg)
	if err != nil {
		log.Printf("Marshal message error: %v", err)
		return
	}

	// 创建并序列化完整消息
	message := &proto.Message{
		Type:    msgType,
		Payload: payload,
	}
	data, err := protobuf.Marshal(message)
	if err != nil {
		log.Printf("Marshal wrapper message error: %v", err)
		return
	}

	// 加密消息
	encrypted, err := conn.crypto.Encrypt(data)
	if err != nil {
		log.Printf("Encrypt message error: %v", err)
		return
	}

	// 发送消息
	if err := conn.conn.WriteMessage(websocket.BinaryMessage, encrypted); err != nil {
		log.Printf("Write message error: %v", err)
	}
}

// sendErrorMessage 发送错误消息
func sendErrorMessage(conn *Connection, errMsg string) {
	// 创建错误消息
	errorMsg := &proto.ErrorMessage{
		Message: errMsg,
	}
	// 使用protobuf消息发送
	sendProtoMessage(conn, proto.MessageType_ERROR, errorMsg)
}

// sendJSONMessage 发送JSON格式的消息
func sendJSONMessage(conn *Connection, data interface{}) {
	// 序列化数据为JSON字符串
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Marshal JSON error: %v", err)
		return
	}

	// 创建系统消息
	sysMsg := &proto.SystemMessage{
		Type:    "json",
		Content: string(jsonData),
	}

	// 使用protobuf消息发送
	sendProtoMessage(conn, proto.MessageType_SYSTEM, sysMsg)
}

// sendSuccessMessage 发送成功消息
func sendSuccessMessage(conn *Connection, data interface{}) {
	// 创建系统消息
	sysMsg := &proto.SystemMessage{
		Type:    "success",
		Content: "success",
	}
	// 使用protobuf消息发送
	sendProtoMessage(conn, proto.MessageType_SYSTEM, sysMsg)

	// 发送具体数据
	sendJSONMessage(conn, data)
}
