package models

import (
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
	Friend    *User     `json:"friend,omitempty" db:"-"` // 关联的好友用户信息
}

// FriendRequest 好友请求模型
type FriendRequest struct {
	ID         int       `json:"id" db:"id"`
	UserID     int       `json:"user_id" db:"user_id"`     // 发送者ID
	ReceiverID int       `json:"receiver_id" db:"receiver_id"` // 接收者ID
	Status     string    `json:"status" db:"status"`      // pending, accepted, rejected
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	Sender     *User     `json:"sender,omitempty" db:"-"`    // 发送者信息
}

// FriendContract 好友契约模型
type FriendContract struct {
	ID             int       `json:"id" db:"id"`
	UserID         int       `json:"user_id" db:"user_id"`
	FriendID       int       `json:"friend_id" db:"friend_id"`
	ContractType   string    `json:"contract_type" db:"contract_type"` // study_buddy, accountability_partner
	ContractTerms  string    `json:"contract_terms" db:"contract_terms"`
	StartDate      time.Time `json:"start_date" db:"start_date"`
	EndDate        time.Time `json:"end_date,omitempty" db:"end_date"`
	GoalType       string    `json:"goal_type,omitempty" db:"goal_type"`
	GoalValue      int       `json:"goal_value,omitempty" db:"goal_value"`
	Status         string    `json:"status" db:"status"` // active, completed, terminated
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	Friend         *User     `json:"friend,omitempty" db:"-"` // 关联的好友用户信息
}

// FriendMessage 好友消息模型
type FriendMessage struct {
	ID          int       `json:"id" db:"id"`
	SenderID    int       `json:"sender_id" db:"sender_id"`
	ReceiverID  int       `json:"receiver_id" db:"receiver_id"`
	MessageType string    `json:"message_type" db:"message_type"` // text, image, etc.
	Content     string    `json:"content" db:"content"`
	IsRead      bool      `json:"is_read" db:"is_read"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	Sender      *User     `json:"sender,omitempty" db:"-"`   // 关联的发送者信息
	Receiver    *User     `json:"receiver,omitempty" db:"-"` // 关联的接收者信息
}

// FriendRepository 好友仓库接口
type FriendRepository interface {
	// 好友关系管理
	FindByID(id int) (*Friend, error)
	GetFriendList(userID int) ([]Friend, error)
	GetFriendRequests(userID int) ([]FriendRequest, error)
	SendFriendRequest(userID, friendID int) error
	AcceptFriendRequest(requestID, userID int) error
	RejectFriendRequest(requestID, userID int) error
	DeleteFriend(userID, friendID int) error
	IsFriend(userID, friendID int) (bool, error)
	
	// 好友契约管理
	CreateContract(contract *FriendContract) error
	GetContractByID(contractID int) (*FriendContract, error)
	GetUserContracts(userID int) ([]FriendContract, error)
	UpdateContractStatus(contractID int, status string) error
	UpdateContract(contract *FriendContract) error
	
	// 好友消息管理
	SaveMessage(message *FriendMessage) error
	GetMessage(messageID int) (*FriendMessage, error)
	MarkMessageAsRead(messageID int) error
	GetChatHistory(userID, friendID int, limit, offset int) ([]FriendMessage, int, error)
	GetUnreadMessageCount(userID int) (int, error)
}

// FriendRequestDTO 好友请求DTO
type FriendRequestDTO struct {
	ReceiverID int `json:"receiver_id" validate:"required"`
}

// FriendResponseDTO 好友请求响应DTO
type FriendResponseDTO struct {
	RequestID int    `json:"request_id" validate:"required"`
	Action    string `json:"action" validate:"required,oneof=accept reject"`
}

// FriendContractCreateDTO 创建好友契约DTO
type FriendContractCreateDTO struct {
	FriendID      int       `json:"friend_id" validate:"required"`
	ContractType  string    `json:"contract_type" validate:"required,oneof=study_buddy accountability_partner"`
	ContractTerms string    `json:"contract_terms"`
	StartDate     time.Time `json:"start_date" validate:"required"`
	EndDate       time.Time `json:"end_date,omitempty"`
	GoalType      string    `json:"goal_type,omitempty"`
	GoalValue     int       `json:"goal_value,omitempty"`
}

// FriendMessageCreateDTO 创建好友消息DTO
type FriendMessageCreateDTO struct {
	ReceiverID  int    `json:"receiver_id" validate:"required"`
	MessageType string `json:"message_type" validate:"required,oneof=text image"`
	Content     string `json:"content" validate:"required"`
}

// FriendInfoDTO 好友信息DTO
type FriendInfoDTO struct {
	ID               int       `json:"id"`
	FriendshipID     int       `json:"friendship_id"`
	Username         string    `json:"username"`
	Avatar           string    `json:"avatar,omitempty"`
	Signature        string    `json:"signature,omitempty"`
	StudyDirection   string    `json:"study_direction,omitempty"`
	TotalStudyTime   int       `json:"total_study_time"`
	FriendSince      time.Time `json:"friend_since"`
	HasActiveContract bool     `json:"has_active_contract"`
	UnreadMessages   int       `json:"unread_messages"`
}

// ToFriendInfoDTO 将Friend转换为FriendInfoDTO
func (f *Friend) ToFriendInfoDTO(hasActiveContract bool, unreadMessages int) *FriendInfoDTO {
	if f == nil || f.Friend == nil {
		return nil
	}
	
	return &FriendInfoDTO{
		ID:                f.Friend.ID,
		FriendshipID:      f.ID,
		Username:          f.Friend.Username,
		Avatar:            f.Friend.Avatar,
		Signature:         f.Friend.Signature,
		StudyDirection:    f.Friend.StudyDirection,
		TotalStudyTime:    f.Friend.TotalStudyTime,
		FriendSince:       f.CreatedAt,
		HasActiveContract: hasActiveContract,
		UnreadMessages:    unreadMessages,
	}
}

// ContractDetailDTO 契约详情DTO
type ContractDetailDTO struct {
	ID             int       `json:"id"`
	ContractType   string    `json:"contract_type"`
	ContractTerms  string    `json:"contract_terms"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date,omitempty"`
	GoalType       string    `json:"goal_type,omitempty"`
	GoalValue      int       `json:"goal_value,omitempty"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	Friend         *UserProfileDTO `json:"friend,omitempty"`
} 