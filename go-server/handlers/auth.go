package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"go-server/models"
	"go-server/proto"
	"time"

	protobuf "google.golang.org/protobuf/proto"
)

// generateToken 生成随机会话令牌
func generateToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// handleAuthMessage 处理认证消息
func handleAuthMessage(conn *Connection, payload []byte) error {
	var authMsg proto.AuthMessage
	if err := protobuf.Unmarshal(payload, &authMsg); err != nil {
		return err
	}

	// 根据Token和DeviceId处理认证
	if authMsg.Token != "" {
		// 处理登出请求
		if err := models.DeleteSession(authMsg.Token); err != nil {
			sendErrorMessage(conn, "Failed to delete session")
			return err
		}

		// 清除连接信息
		conn.userID = 0
		conn.deviceID = ""
		conn.crypto = nil

		sendSuccessMessage(conn, "Logout successful")
	} else {
		// 处理登录请求
		user, err := models.AuthenticateUser(authMsg.DeviceId, "")
		if err != nil {
			sendErrorMessage(conn, "Invalid credentials")
			return err
		}

		// 生成会话令牌
		token, err := generateToken()
		if err != nil {
			sendErrorMessage(conn, "Failed to create session")
			return err
		}

		// 创建会话
		expiresAt := time.Now().Add(24 * time.Hour)
		_, err = models.CreateSession(user.ID, token, expiresAt)
		if err != nil {
			sendErrorMessage(conn, "Failed to create session")
			return err
		}

		// 设置连接信息
		conn.userID = user.ID
		conn.deviceID = authMsg.DeviceId
		conn.crypto = globalManager.crypto

		// 发送登录成功响应
		sendSuccessMessage(conn, "Login successful")
		sendJSONMessage(conn, map[string]interface{}{
			"token": token,
			"user":  user,
		})
	}

	return nil
}

func init() {
	// 注册认证消息处理器
	RegisterMessageHandler(proto.MessageType_AUTH, handleAuthMessage)
}
