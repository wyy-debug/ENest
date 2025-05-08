package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"NewENest/go-server/models"
)

// FriendRepository 好友数据库操作实现
type FriendRepository struct {
	db *sqlx.DB
}

// NewFriendRepository 创建好友仓库实例
func NewFriendRepository(db *sqlx.DB) models.FriendRepository {
	return &FriendRepository{
		db: db,
	}
}

// FindByID 根据ID查找好友关系
func (r *FriendRepository) FindByID(id int) (*models.Friend, error) {
	query := `
		SELECT f.id, f.user_id, f.friend_id, f.status, f.created_at, f.updated_at,
		       u.id, u.username, u.email, u.avatar, u.signature, u.study_direction, u.total_study_time
		FROM friends f
		JOIN users u ON f.friend_id = u.id
		WHERE f.id = $1
	`
	
	friend := &models.Friend{}
	friendUser := &models.User{}
	
	err := r.db.QueryRowx(query, id).Scan(
		&friend.ID, &friend.UserID, &friend.FriendID, &friend.Status, &friend.CreatedAt, &friend.UpdatedAt,
		&friendUser.ID, &friendUser.Username, &friendUser.Email, &friendUser.Avatar, &friendUser.Signature, 
		&friendUser.StudyDirection, &friendUser.TotalStudyTime,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("好友关系不存在")
		}
		return nil, err
	}
	
	friend.Friend = friendUser
	return friend, nil
}

// GetFriendList 获取用户的好友列表
func (r *FriendRepository) GetFriendList(userID int) ([]models.Friend, error) {
	query := `
		SELECT f.id, f.user_id, f.friend_id, f.status, f.created_at, f.updated_at,
		       u.id, u.username, u.email, u.avatar, u.signature, u.study_direction, u.total_study_time
		FROM friends f
		JOIN users u ON CASE
			WHEN f.user_id = $1 THEN f.friend_id = u.id
			WHEN f.friend_id = $1 THEN f.user_id = u.id
		END
		WHERE (f.user_id = $1 OR f.friend_id = $1) AND f.status = 'accepted'
	`
	
	rows, err := r.db.Queryx(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var friends []models.Friend
	
	for rows.Next() {
		friend := models.Friend{}
		friendUser := models.User{}
		
		err := rows.Scan(
			&friend.ID, &friend.UserID, &friend.FriendID, &friend.Status, &friend.CreatedAt, &friend.UpdatedAt,
			&friendUser.ID, &friendUser.Username, &friendUser.Email, &friendUser.Avatar, &friendUser.Signature, 
			&friendUser.StudyDirection, &friendUser.TotalStudyTime,
		)
		
		if err != nil {
			return nil, err
		}
		
		// 确保 FriendID 是好友的ID，而不是当前用户的ID
		if friend.UserID == userID {
			friend.Friend = &friendUser
		} else {
			friend.Friend = &friendUser
			// 在UI展示中，我们希望UserID始终是当前用户，FriendID始终是好友
			// 但是不修改数据库中的记录，只在返回时调整
			friend.FriendID = friend.UserID
			friend.UserID = userID
		}
		
		friends = append(friends, friend)
	}
	
	return friends, nil
}

// GetFriendRequests 获取发送给用户的好友请求
func (r *FriendRepository) GetFriendRequests(userID int) ([]models.Friend, error) {
	query := `
		SELECT f.id, f.user_id, f.friend_id, f.status, f.created_at, f.updated_at,
		       u.id, u.username, u.email, u.avatar, u.signature, u.study_direction, u.total_study_time
		FROM friends f
		JOIN users u ON f.user_id = u.id
		WHERE f.friend_id = $1 AND f.status = 'pending'
	`
	
	rows, err := r.db.Queryx(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var requests []models.Friend
	
	for rows.Next() {
		request := models.Friend{}
		requestUser := models.User{}
		
		err := rows.Scan(
			&request.ID, &request.UserID, &request.FriendID, &request.Status, &request.CreatedAt, &request.UpdatedAt,
			&requestUser.ID, &requestUser.Username, &requestUser.Email, &requestUser.Avatar, &requestUser.Signature, 
			&requestUser.StudyDirection, &requestUser.TotalStudyTime,
		)
		
		if err != nil {
			return nil, err
		}
		
		request.Friend = &requestUser
		requests = append(requests, request)
	}
	
	return requests, nil
}

// SendFriendRequest 发送好友请求
func (r *FriendRepository) SendFriendRequest(userID, friendID int) error {
	// 检查是否已存在请求
	var count int
	err := r.db.QueryRowx(
		"SELECT COUNT(*) FROM friends WHERE (user_id = $1 AND friend_id = $2) OR (user_id = $2 AND friend_id = $1)",
		userID, friendID,
	).Scan(&count)
	
	if err != nil {
		return err
	}
	
	if count > 0 {
		return errors.New("已存在好友关系或请求")
	}
	
	// 插入好友请求
	_, err = r.db.Exec(
		"INSERT INTO friends (user_id, friend_id, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $4)",
		userID, friendID, "pending", time.Now(),
	)
	
	return err
}

// AcceptFriendRequest 接受好友请求
func (r *FriendRepository) AcceptFriendRequest(requestID, userID int) error {
	// 验证请求是否发送给该用户
	var count int
	err := r.db.QueryRowx(
		"SELECT COUNT(*) FROM friends WHERE id = $1 AND friend_id = $2 AND status = 'pending'",
		requestID, userID,
	).Scan(&count)
	
	if err != nil {
		return err
	}
	
	if count == 0 {
		return errors.New("好友请求不存在或不是发送给你的")
	}
	
	// 更新请求状态
	_, err = r.db.Exec(
		"UPDATE friends SET status = 'accepted', updated_at = $1 WHERE id = $2",
		time.Now(), requestID,
	)
	
	return err
}

// RejectFriendRequest 拒绝好友请求
func (r *FriendRepository) RejectFriendRequest(requestID, userID int) error {
	// 验证请求是否发送给该用户
	var count int
	err := r.db.QueryRowx(
		"SELECT COUNT(*) FROM friends WHERE id = $1 AND friend_id = $2 AND status = 'pending'",
		requestID, userID,
	).Scan(&count)
	
	if err != nil {
		return err
	}
	
	if count == 0 {
		return errors.New("好友请求不存在或不是发送给你的")
	}
	
	// 更新请求状态
	_, err = r.db.Exec(
		"UPDATE friends SET status = 'rejected', updated_at = $1 WHERE id = $2",
		time.Now(), requestID,
	)
	
	return err
}

// DeleteFriend 删除好友关系
func (r *FriendRepository) DeleteFriend(userID, friendID int) error {
	_, err := r.db.Exec(
		"DELETE FROM friends WHERE (user_id = $1 AND friend_id = $2) OR (user_id = $2 AND friend_id = $1)",
		userID, friendID,
	)
	
	return err
}

// IsFriend 检查两个用户是否为好友
func (r *FriendRepository) IsFriend(userID, friendID int) (bool, error) {
	var count int
	err := r.db.QueryRowx(
		"SELECT COUNT(*) FROM friends WHERE ((user_id = $1 AND friend_id = $2) OR (user_id = $2 AND friend_id = $1)) AND status = 'accepted'",
		userID, friendID,
	).Scan(&count)
	
	if err != nil {
		return false, err
	}
	
	return count > 0, nil
}

// SaveMessage 保存消息
func (r *FriendRepository) SaveMessage(message *models.FriendMessage) error {
	query := `
		INSERT INTO friend_messages (sender_id, receiver_id, message_type, content, is_read, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	
	err := r.db.QueryRowx(
		query,
		message.SenderID, message.ReceiverID, message.MessageType, message.Content, message.IsRead, message.CreatedAt,
	).Scan(&message.ID)
	
	return err
}

// GetMessage 根据ID获取消息
func (r *FriendRepository) GetMessage(messageID int) (*models.FriendMessage, error) {
	query := `
		SELECT id, sender_id, receiver_id, message_type, content, is_read, created_at
		FROM friend_messages
		WHERE id = $1
	`
	
	message := &models.FriendMessage{}
	err := r.db.QueryRowx(query, messageID).Scan(
		&message.ID, &message.SenderID, &message.ReceiverID, &message.MessageType,
		&message.Content, &message.IsRead, &message.CreatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("消息不存在")
		}
		return nil, err
	}
	
	return message, nil
}

// MarkMessageAsRead 将消息标记为已读
func (r *FriendRepository) MarkMessageAsRead(messageID int) error {
	_, err := r.db.Exec(
		"UPDATE friend_messages SET is_read = true WHERE id = $1",
		messageID,
	)
	
	return err
}

// GetChatHistory 获取两个用户之间的聊天记录
func (r *FriendRepository) GetChatHistory(userID, friendID int, limit, offset int) ([]models.FriendMessage, int, error) {
	// 获取总消息数
	var total int
	err := r.db.QueryRowx(
		"SELECT COUNT(*) FROM friend_messages WHERE (sender_id = $1 AND receiver_id = $2) OR (sender_id = $2 AND receiver_id = $1)",
		userID, friendID,
	).Scan(&total)
	
	if err != nil {
		return nil, 0, err
	}
	
	// 获取消息列表
	query := `
		SELECT id, sender_id, receiver_id, message_type, content, is_read, created_at
		FROM friend_messages
		WHERE (sender_id = $1 AND receiver_id = $2) OR (sender_id = $2 AND receiver_id = $1)
		ORDER BY created_at DESC
		LIMIT $3 OFFSET $4
	`
	
	rows, err := r.db.Queryx(query, userID, friendID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	
	var messages []models.FriendMessage
	
	for rows.Next() {
		message := models.FriendMessage{}
		
		err := rows.Scan(
			&message.ID, &message.SenderID, &message.ReceiverID, &message.MessageType,
			&message.Content, &message.IsRead, &message.CreatedAt,
		)
		
		if err != nil {
			return nil, 0, err
		}
		
		messages = append(messages, message)
	}
	
	// 标记收到的消息为已读
	_, err = r.db.Exec(
		"UPDATE friend_messages SET is_read = true WHERE sender_id = $1 AND receiver_id = $2 AND is_read = false",
		friendID, userID,
	)
	
	if err != nil {
		fmt.Printf("标记消息为已读失败: %v\n", err)
		// 不返回错误，继续处理
	}
	
	return messages, total, nil
}

// GetUnreadMessageCount 获取用户未读消息数量
func (r *FriendRepository) GetUnreadMessageCount(userID int) (int, error) {
	var count int
	err := r.db.QueryRowx(
		"SELECT COUNT(*) FROM friend_messages WHERE receiver_id = $1 AND is_read = false",
		userID,
	).Scan(&count)
	
	if err != nil {
		return 0, err
	}
	
	return count, nil
}

// CreateContract 创建好友契约
func (r *FriendRepository) CreateContract(contract *models.FriendContract) error {
	query := `INSERT INTO friend_contracts 
		(user_id, friend_id, contract_type, contract_terms, start_date, end_date, goal_type, goal_value, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`
		
	var id int
	err := r.db.QueryRowx(
		query,
		contract.UserID,
		contract.FriendID,
		contract.ContractType,
		contract.ContractTerms,
		contract.StartDate,
		contract.EndDate,
		contract.GoalType,
		contract.GoalValue,
		contract.Status,
	).Scan(&id)
	
	if err != nil {
		return err
	}
	
	contract.ID = id
	return nil
}

// GetContractByID 通过ID获取契约
func (r *FriendRepository) GetContractByID(contractID int) (*models.FriendContract, error) {
	query := `
		SELECT id, user_id, friend_id, contract_type, contract_terms, start_date, end_date, 
		       goal_type, goal_value, status, created_at, updated_at
		FROM friend_contracts
		WHERE id = $1
	`
	
	contract := &models.FriendContract{}
	err := r.db.QueryRowx(query, contractID).Scan(
		&contract.ID, &contract.UserID, &contract.FriendID, &contract.ContractType,
		&contract.ContractTerms, &contract.StartDate, &contract.EndDate,
		&contract.GoalType, &contract.GoalValue, &contract.Status,
		&contract.CreatedAt, &contract.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}
	
	return contract, nil
}

// GetUserContracts 获取用户所有契约
func (r *FriendRepository) GetUserContracts(userID int) ([]models.FriendContract, error) {
	query := `
		SELECT id, user_id, friend_id, contract_type, contract_terms, start_date, end_date, 
		       goal_type, goal_value, status, created_at, updated_at
		FROM friend_contracts
		WHERE user_id = $1 OR friend_id = $1
	`
	
	rows, err := r.db.Queryx(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var contracts []models.FriendContract
	
	for rows.Next() {
		contract := models.FriendContract{}
		
		err := rows.Scan(
			&contract.ID, &contract.UserID, &contract.FriendID, &contract.ContractType,
			&contract.ContractTerms, &contract.StartDate, &contract.EndDate,
			&contract.GoalType, &contract.GoalValue, &contract.Status,
			&contract.CreatedAt, &contract.UpdatedAt,
		)
		
		if err != nil {
			return nil, err
		}
		
		contracts = append(contracts, contract)
	}
	
	return contracts, nil
}

// UpdateContractStatus 更新契约状态
func (r *FriendRepository) UpdateContractStatus(contractID int, status string) error {
	_, err := r.db.Exec(
		"UPDATE friend_contracts SET status = $1, updated_at = $2 WHERE id = $3",
		status, time.Now(), contractID,
	)
	
	return err
}

// UpdateContract 更新契约信息
func (r *FriendRepository) UpdateContract(contract *models.FriendContract) error {
	_, err := r.db.Exec(
		`UPDATE friend_contracts 
		 SET contract_terms = $1, end_date = $2, goal_type = $3, goal_value = $4, 
		     status = $5, updated_at = $6
		 WHERE id = $7`,
		contract.ContractTerms, contract.EndDate, contract.GoalType, contract.GoalValue,
		contract.Status, time.Now(), contract.ID,
	)
	
	return err
} 