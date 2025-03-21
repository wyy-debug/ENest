package handlers

import (
	"go-server/models"
	"go-server/proto"
	"time"

	protobuf "google.golang.org/protobuf/proto"
)

// CreateStudyRoom 创建自习室
// GetStudyRoomDetail 获取自习室详细信息
// JoinStudyRoom 加入自习室
// JoinStudyRoomByShareLink 通过分享链接加入自习室
// LeaveStudyRoom 离开自习室
// DestroyStudyRoom 销毁自习室

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
