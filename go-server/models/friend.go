package models

import (
	"errors"
	"go-server/database"
	"time"
)

// Friend 好友关系模型
type Friend struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	FriendID  int       `json:"friend_id" db:"friend_id"`
	Status    string    `json:"status" db:"status"` // pending, accepted, rejected
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Friend    *User     `json:"friend,omitempty"` // 关联的好友用户信息
}

// FriendContract 好友契约模型
type FriendContract struct {
	ID            int       `json:"id" db:"id"`
	UserID        int       `json:"user_id" db:"user_id"`
	FriendID      int       `json:"friend_id" db:"friend_id"`
	ContractType  string    `json:"contract_type" db:"contract_type"` // study_buddy, accountability_partner
	ContractTerms string    `json:"contract_terms" db:"contract_terms"`
	Status        string    `json:"status" db:"status"` // active, terminated
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	Friend        *User     `json:"friend,omitempty"` // 关联的好友用户信息
}

// FriendMessage 好友消息模型
type FriendMessage struct {
	ID          int       `json:"id" db:"id"`
	SenderID    int       `json:"sender_id" db:"sender_id"`
	ReceiverID  int       `json:"receiver_id" db:"receiver_id"`
	MessageType string    `json:"message_type" db:"message_type"` // text, image
	Content     string    `json:"content" db:"content"`
	IsRead      bool      `json:"is_read" db:"is_read"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	Sender      *User     `json:"sender,omitempty"`   // 关联的发送者信息
	Receiver    *User     `json:"receiver,omitempty"` // 关联的接收者信息
}

// GetFriendList 获取用户的好友列表
func GetFriendList(userID int) ([]Friend, error) {
	db := database.GetDB()
	rows, err := db.Query(`
		SELECT f.id, f.user_id, f.friend_id, f.status, f.created_at, f.updated_at,
			u.id as friend_user_id, u.username as friend_username, u.email as friend_email
		FROM friends f
		JOIN users u ON (CASE 
			WHEN f.user_id = $1 THEN f.friend_id = u.id
			WHEN f.friend_id = $1 THEN f.user_id = u.id
		END)
		WHERE (f.user_id = $1 OR f.friend_id = $1) AND f.status = 'accepted'
	`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var friends []Friend
	for rows.Next() {
		var f Friend
		f.Friend = &User{}
		err := rows.Scan(
			&f.ID, &f.UserID, &f.FriendID, &f.Status, &f.CreatedAt, &f.UpdatedAt,
			&f.Friend.ID, &f.Friend.Username, &f.Friend.Email,
		)
		if err != nil {
			return nil, err
		}
		friends = append(friends, f)
	}

	return friends, nil
}

// SendFriendRequest 发送好友请求
func SendFriendRequest(userID, friendID int) error {
	db := database.GetDB()

	// 检查是否已经是好友
	var exists bool
	err := db.QueryRow(`
		SELECT EXISTS (
			SELECT 1 FROM friends 
			WHERE (user_id = $1 AND friend_id = $2) 
			OR (user_id = $2 AND friend_id = $1)
		)`, userID, friendID).Scan(&exists)

	if err != nil {
		return err
	}

	if exists {
		return errors.New("friend request already exists")
	}

	// 创建好友请求
	_, err = db.Exec(`
		INSERT INTO friends (user_id, friend_id, status)
		VALUES ($1, $2, 'pending')
	`, userID, friendID)

	return err
}

// HandleFriendRequest 处理好友请求
func HandleFriendRequest(requestID, userID int, action string) error {
	if action != "accept" && action != "reject" {
		return errors.New("invalid action")
	}

	db := database.GetDB()
	result, err := db.Exec(`
		UPDATE friends
		SET status = $1, updated_at = CURRENT_TIMESTAMP
		WHERE id = $2 AND friend_id = $3 AND status = 'pending'
	`, action, requestID, userID)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("friend request not found")
	}

	return nil
}

// DeleteFriend 删除好友关系
func DeleteFriend(userID, friendID int) error {
	db := database.GetDB()
	result, err := db.Exec(`
		DELETE FROM friends
		WHERE (user_id = $1 AND friend_id = $2)
		OR (user_id = $2 AND friend_id = $1)
	`, userID, friendID)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("friend relationship not found")
	}

	return nil
}

// CreateFriendContract 创建好友契约
func CreateFriendContract(userID int, contract *FriendContract) error {
	db := database.GetDB()

	// 检查是否是好友关系
	var exists bool
	err := db.QueryRow(`
		SELECT EXISTS (
			SELECT 1 FROM friends
			WHERE ((user_id = $1 AND friend_id = $2) OR (user_id = $2 AND friend_id = $1))
			AND status = 'accepted'
		)`, userID, contract.FriendID).Scan(&exists)

	if err != nil {
		return err
	}

	if !exists {
		return errors.New("friend relationship not found")
	}

	// 创建契约
	err = db.QueryRow(`
		INSERT INTO friend_contracts (user_id, friend_id, contract_type, contract_terms, status)
		VALUES ($1, $2, $3, $4, 'active')
		RETURNING id, created_at, updated_at
	`, userID, contract.FriendID, contract.ContractType, contract.ContractTerms).Scan(
		&contract.ID, &contract.CreatedAt, &contract.UpdatedAt)

	return err
}

// GetFriendContracts 获取好友契约列表
func GetFriendContracts(userID int) ([]FriendContract, error) {
	db := database.GetDB()
	rows, err := db.Query(`
		SELECT c.*, u.id as friend_user_id, u.username as friend_username, u.email as friend_email
		FROM friend_contracts c
		JOIN users u ON (CASE 
			WHEN c.user_id = $1 THEN c.friend_id = u.id
			WHEN c.friend_id = $1 THEN c.user_id = u.id
		END)
		WHERE c.user_id = $1 OR c.friend_id = $1
	`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contracts []FriendContract
	for rows.Next() {
		var c FriendContract
		c.Friend = &User{}
		err := rows.Scan(
			&c.ID, &c.UserID, &c.FriendID, &c.ContractType, &c.ContractTerms,
			&c.Status, &c.CreatedAt, &c.UpdatedAt,
			&c.Friend.ID, &c.Friend.Username, &c.Friend.Email,
		)
		if err != nil {
			return nil, err
		}
		contracts = append(contracts, c)
	}

	return contracts, nil
}

// TerminateFriendContract 终止好友契约
func TerminateFriendContract(contractID, userID int) error {
	db := database.GetDB()
	result, err := db.Exec(`
		UPDATE friend_contracts
		SET status = 'terminated', updated_at = CURRENT_TIMESTAMP
		WHERE id = $1 AND (user_id = $2 OR friend_id = $2) AND status = 'active'
	`, contractID, userID)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("contract not found or already terminated")
	}

	return nil
}

// SaveFriendMessage 保存好友消息
func SaveFriendMessage(message *FriendMessage) error {
	db := database.GetDB()
	err := db.QueryRow(`
		INSERT INTO friend_messages (sender_id, receiver_id, message_type, content, is_read)
		VALUES ($1, $2, $3, $4, false)
		RETURNING id, created_at
	`, message.SenderID, message.ReceiverID, message.MessageType, message.Content).Scan(
		&message.ID, &message.CreatedAt)

	return err
}

// GetChatHistory 获取聊天历史记录
func GetChatHistory(userID, friendID int, limit int) ([]FriendMessage, error) {
	db := database.GetDB()
	rows, err := db.Query(`
		SELECT m.*, 
			s.id as sender_id, s.username as sender_username, s.email as sender_email,
			r.id as receiver_id, r.username as receiver_username, r.email as receiver_email
		FROM friend_messages m
		JOIN users s ON m.sender_id = s.id
		JOIN users r ON m.receiver_id = r.id
		WHERE (m.sender_id = $1 AND m.receiver_id = $2)
		OR (m.sender_id = $2 AND m.receiver_id = $1)
		ORDER BY m.created_at DESC
		LIMIT $3
	`, userID, friendID, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []FriendMessage
	for rows.Next() {
		var m FriendMessage
		m.Sender = &User{}
		m.Receiver = &User{}
		err := rows.Scan(
			&m.ID, &m.SenderID, &m.ReceiverID, &m.MessageType, &m.Content, &m.IsRead, &m.CreatedAt,
			&m.Sender.ID, &m.Sender.Username, &m.Sender.Email,
			&m.Receiver.ID, &m.Receiver.Username, &m.Receiver.Email,
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}

	return messages, nil
}

// CheckFriendRelationship 检查两个用户是否为好友关系
func CheckFriendRelationship(userID, friendID int) (bool, error) {
    db := database.GetDB()
    var exists bool
    err := db.QueryRow(`
        SELECT EXISTS (
            SELECT 1 FROM friends
            WHERE ((user_id = $1 AND friend_id = $2) OR (user_id = $2 AND friend_id = $1))
            AND status = 'accepted'
        )`, userID, friendID).Scan(&exists)

    return exists, err
}
