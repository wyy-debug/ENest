package handlers

import (
	"go-server/models"
	"go-server/proto"
	"log"

	protobuf "google.golang.org/protobuf/proto"
)

// handleFriendMessage 处理好友消息
func handleFriendMessage(conn *Connection, payload []byte) error {
	var friendMsg proto.FriendMessage
	if err := protobuf.Unmarshal(payload, &friendMsg); err != nil {
		log.Printf("Unmarshal friend message error: %v", err)
		return err
	}

	switch friendMsg.Operation {
	case proto.FriendMessage_SEND_REQUEST:
		// 发送好友请求
		err := models.SendFriendRequest(conn.userID, int(friendMsg.FriendId))
		if err != nil {
			sendErrorMessage(conn, err.Error())
			return err
		}
		sendSuccessMessage(conn, "Friend request sent")

	case proto.FriendMessage_HANDLE_REQUEST:
		// 处理好友请求
		err := models.HandleFriendRequest(int(friendMsg.ContractId), conn.userID, friendMsg.Action)
		if err != nil {
			sendErrorMessage(conn, err.Error())
			return err
		}
		sendSuccessMessage(conn, "Friend request handled")

	case proto.FriendMessage_GET_LIST:
		// 获取好友列表
		friends, err := models.GetFriendList(conn.userID)
		if err != nil {
			sendErrorMessage(conn, err.Error())
			return err
		}
		sendJSONMessage(conn, friends)

	case proto.FriendMessage_DELETE:
		// 删除好友
		err := models.DeleteFriend(conn.userID, int(friendMsg.FriendId))
		if err != nil {
			sendErrorMessage(conn, err.Error())
			return err
		}
		sendSuccessMessage(conn, "Friend deleted")

	case proto.FriendMessage_CREATE_CONTRACT:
		// 创建好友契约
		contract := &models.FriendContract{
			UserID:        conn.userID,
			FriendID:      int(friendMsg.FriendId),
			ContractType:  friendMsg.ContractType,
			ContractTerms: friendMsg.ContractTerms,
		}
		err := models.CreateFriendContract(conn.userID, contract)
		if err != nil {
			sendErrorMessage(conn, err.Error())
			return err
		}
		sendJSONMessage(conn, contract)

	case proto.FriendMessage_TERMINATE_CONTRACT:
		// 终止好友契约
		err := models.TerminateFriendContract(int(friendMsg.ContractId), conn.userID)
		if err != nil {
			sendErrorMessage(conn, err.Error())
			return err
		}
		sendSuccessMessage(conn, "Contract terminated")

	case proto.FriendMessage_GET_CONTRACTS:
		// 获取契约列表
		contracts, err := models.GetFriendContracts(conn.userID)
		if err != nil {
			sendErrorMessage(conn, err.Error())
			return err
		}
		sendJSONMessage(conn, contracts)
	}

	return nil
}

func init() {
	// 注册好友消息处理器
	RegisterMessageHandler(proto.MessageType_FRIEND, handleFriendMessage)
}
