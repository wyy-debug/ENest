package models

import (
	"time"
)

// StudyRoom 自习室模型
type StudyRoom struct {
	ID              int       `json:"id" db:"id"`
	OwnerID         int       `json:"owner_id" db:"owner_id"`
	Name            string    `json:"name" db:"name"`
	Description     string    `json:"description,omitempty" db:"description"`
	ShareLink       string    `json:"share_link,omitempty" db:"share_link"`
	MaxMembers      int       `json:"max_members" db:"max_members"`
	IsPrivate       bool      `json:"is_private" db:"is_private"`
	Theme           string    `json:"theme,omitempty" db:"theme"`
	BackgroundImage string    `json:"background_image,omitempty" db:"background_image"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	ExpiresAt       time.Time `json:"expires_at" db:"expires_at"`
	
	// 关联信息
	Owner           *User     `json:"owner,omitempty" db:"-"`
	Members         []RoomMember `json:"members,omitempty" db:"-"`
	MemberCount     int       `json:"member_count,omitempty" db:"-"`
	OnlineCount     int       `json:"online_count,omitempty" db:"-"`
}

// RoomMember 自习室成员模型
type RoomMember struct {
	ID          int       `json:"id" db:"id"`
	RoomID      int       `json:"room_id" db:"room_id"`
	UserID      int       `json:"user_id" db:"user_id"`
	IsAnonymous bool      `json:"is_anonymous" db:"is_anonymous"`
	Role        string    `json:"role" db:"role"` // 角色：owner, admin, member
	Status      string    `json:"status" db:"status"` // 状态：online, offline, away
	JoinedAt    time.Time `json:"joined_at" db:"joined_at"`
	
	// 关联信息
	User        *User     `json:"user,omitempty" db:"-"`
}

// StudyRoomRepository 自习室仓库接口
type StudyRoomRepository interface {
	FindByID(id int) (*StudyRoom, error)
	FindByShareLink(shareLink string) (*StudyRoom, error)
	GetUserRooms(userID int) ([]StudyRoom, error)
	GetPublicRooms(limit, offset int) ([]StudyRoom, int, error)
	Create(room *StudyRoom) error
	Update(room *StudyRoom) error
	Delete(id int) error
	
	// 成员管理
	GetRoomMembers(roomID int) ([]RoomMember, error)
	GetRoomMember(roomID, userID int) (*RoomMember, error)
	AddMember(member *RoomMember) error
	UpdateMemberStatus(roomID, userID int, status string) error
	RemoveMember(roomID, userID int) error
	
	// 统计信息
	GetRoomMemberCount(roomID int) (int, error)
	GetRoomOnlineCount(roomID int) (int, error)
}

// StudyRoomCreateDTO 创建自习室DTO
type StudyRoomCreateDTO struct {
	Name            string    `json:"name" validate:"required,min=3,max=100"`
	Description     string    `json:"description,omitempty"`
	MaxMembers      int       `json:"max_members" validate:"required,min=1,max=100"`
	IsPrivate       bool      `json:"is_private"`
	Theme           string    `json:"theme,omitempty"`
	BackgroundImage string    `json:"background_image,omitempty"`
	Duration        int       `json:"duration" validate:"required,min=1,max=1440"` // 持续时间（分钟）
}

// StudyRoomUpdateDTO 更新自习室DTO
type StudyRoomUpdateDTO struct {
	Name            string    `json:"name,omitempty" validate:"omitempty,min=3,max=100"`
	Description     string    `json:"description,omitempty"`
	MaxMembers      int       `json:"max_members,omitempty" validate:"omitempty,min=1,max=100"`
	IsPrivate       *bool     `json:"is_private,omitempty"`
	Theme           string    `json:"theme,omitempty"`
	BackgroundImage string    `json:"background_image,omitempty"`
}

// StudyRoomDTO 自习室信息DTO
type StudyRoomDTO struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description,omitempty"`
	ShareLink       string    `json:"share_link,omitempty"`
	MaxMembers      int       `json:"max_members"`
	IsPrivate       bool      `json:"is_private"`
	Theme           string    `json:"theme,omitempty"`
	BackgroundImage string    `json:"background_image,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	ExpiresAt       time.Time `json:"expires_at"`
	Owner           *UserProfileDTO `json:"owner,omitempty"`
	MemberCount     int       `json:"member_count"`
	OnlineCount     int       `json:"online_count"`
}

// ToRoomDTO 将StudyRoom转换为StudyRoomDTO
func (r *StudyRoom) ToRoomDTO() *StudyRoomDTO {
	if r == nil {
		return nil
	}
	
	return &StudyRoomDTO{
		ID:              r.ID,
		Name:            r.Name,
		Description:     r.Description,
		ShareLink:       r.ShareLink,
		MaxMembers:      r.MaxMembers,
		IsPrivate:       r.IsPrivate,
		Theme:           r.Theme,
		BackgroundImage: r.BackgroundImage,
		CreatedAt:       r.CreatedAt,
		ExpiresAt:       r.ExpiresAt,
		Owner:           r.Owner.ToProfileDTO(),
		MemberCount:     r.MemberCount,
		OnlineCount:     r.OnlineCount,
	}
}

// RoomMemberDTO 自习室成员DTO
type RoomMemberDTO struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Username    string    `json:"username"`
	Avatar      string    `json:"avatar,omitempty"`
	IsAnonymous bool      `json:"is_anonymous"`
	Role        string    `json:"role"`
	Status      string    `json:"status"`
	JoinedAt    time.Time `json:"joined_at"`
}

// ToMemberDTO 将RoomMember转换为RoomMemberDTO
func (m *RoomMember) ToMemberDTO() *RoomMemberDTO {
	if m == nil || m.User == nil {
		return nil
	}
	
	return &RoomMemberDTO{
		ID:          m.ID,
		UserID:      m.UserID,
		Username:    m.User.Username,
		Avatar:      m.User.Avatar,
		IsAnonymous: m.IsAnonymous,
		Role:        m.Role,
		Status:      m.Status,
		JoinedAt:    m.JoinedAt,
	}
} 