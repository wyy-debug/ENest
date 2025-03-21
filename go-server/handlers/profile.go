package handlers

import (
	"go-server/models"
	"go-server/proto"
	"log"

	protobuf "google.golang.org/protobuf/proto"
)

// handleProfileMessage 处理个人信息消息
func handleProfileMessage(conn *Connection, payload []byte) error {
	var profileMsg proto.ProfileMessage
	if err := protobuf.Unmarshal(payload, &profileMsg); err != nil {
		log.Printf("Unmarshal profile message error: %v", err)
		return err
	}

	switch profileMsg.Operation {
	case proto.ProfileMessage_UPDATE:
		// 更新个人信息
		updates := make(map[string]string)
		if profileMsg.Username != "" {
			updates["username"] = profileMsg.Username
		}
		if profileMsg.Signature != "" {
			updates["signature"] = profileMsg.Signature
		}
		if profileMsg.StudyDirection != "" {
			updates["study_direction"] = profileMsg.StudyDirection
		}

		err := models.UpdateUserProfile(conn.userID, updates)
		if err != nil {
			sendErrorMessage(conn, err.Error())
			return err
		}
		sendSuccessMessage(conn, "Profile updated")

	case proto.ProfileMessage_GET:
		// 获取个人信息
		user, err := models.GetUserByID(conn.userID)
		if err != nil {
			sendErrorMessage(conn, err.Error())
			return err
		}
		sendJSONMessage(conn, user)
	}

	return nil
}

func init() {
	// 注册个人信息消息处理器
	RegisterMessageHandler(proto.MessageType_PROFILE, handleProfileMessage)
}
