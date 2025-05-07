package models

import (
	"time"
)

// StudyRoom 自习室模型
type StudyRoom struct {
	ID              int       `json:"id" db:"id"`
	OwnerID         int       `json:"owner_id" db:"owner_id"`
	Name            string    `json:"name" db:"name"`
	Description     string    `json:"description" db:"description"`
	ShareLink       string    `json:"share_link,omitempty" db:"share_link"`
	MaxMembers      int       `json:"max_members" db:"max_members"`
	IsPrivate       bool      `json:"is_private" db:"is_private"`
	Theme           string    `json:"theme" db:"theme"`
	BackgroundImage string    `json:"background_image,omitempty" db:"background_image"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	ExpiresAt       time.Time `json:"expires_at" db:"expires_at"`
	Owner           *User     `json:"owner,omitempty" db:"-"` // 关联的用户对象
}

// RoomMember 自习室成员模型
type RoomMember struct {
	ID          int       `json:"id" db:"id"`
	RoomID      int       `json:"room_id" db:"room_id"`
	UserID      int       `json:"user_id" db:"user_id"`
	IsAnonymous bool      `json:"is_anonymous" db:"is_anonymous"`
	Role        string    `json:"role" db:"role"`
	Status      string    `json:"status" db:"status"`
	JoinedAt    time.Time `json:"joined_at" db:"joined_at"`
	User        *User     `json:"user,omitempty" db:"-"` // 关联的用户对象
}

// StudyRoomRepository 自习室数据访问接口
type StudyRoomRepository interface {
	// 基础CRUD操作
	FindByID(id int) (*StudyRoom, error)
	FindByShareLink(shareLink string) (*StudyRoom, error)
	Create(room *StudyRoom) error
	Update(room *StudyRoom) error
	Delete(id int) error
	
	// 查询操作
	FindAll(page, pageSize int) ([]StudyRoom, int, error)
	FindByOwnerID(ownerID int) ([]StudyRoom, error)
	FindActivePublicRooms(page, pageSize int) ([]StudyRoom, int, error)
	Search(keyword string, page, pageSize int) ([]StudyRoom, int, error)
	
	// 成员管理
	AddMember(roomID, userID int, isAnonymous bool, role string) (*RoomMember, error)
	RemoveMember(roomID, userID int) error
	UpdateMemberStatus(roomID, userID int, status string) error
	UpdateMemberRole(roomID, userID int, role string) error
	
	// 成员查询
	GetRoomMembers(roomID int) ([]RoomMember, error)
	FindRoomsByMemberID(userID int) ([]StudyRoom, error)
	GetMemberCount(roomID int) (int, error)
	IsMember(roomID, userID int) (bool, error)
}

// StudyRoomCreateDTO 自习室创建数据传输对象
type StudyRoomCreateDTO struct {
	Name            string    `json:"name" validate:"required,min=2,max=100"`
	Description     string    `json:"description"`
	MaxMembers      int       `json:"max_members" validate:"required,min=1,max=100"`
	IsPrivate       bool      `json:"is_private"`
	Theme           string    `json:"theme" validate:"omitempty"`
	BackgroundImage string    `json:"background_image" validate:"omitempty"`
	ExpiresIn       int       `json:"expires_in" validate:"required,min=1"` // 过期时间（小时）
}

// StudyRoomUpdateDTO 自习室更新数据传输对象
type StudyRoomUpdateDTO struct {
	Name            *string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Description     *string `json:"description,omitempty"`
	MaxMembers      *int    `json:"max_members,omitempty" validate:"omitempty,min=1,max=100"`
	IsPrivate       *bool   `json:"is_private,omitempty"`
	Theme           *string `json:"theme,omitempty"`
	BackgroundImage *string `json:"background_image,omitempty"`
}

// StudyRoomDetailDTO 自习室详情数据传输对象
type StudyRoomDetailDTO struct {
	ID              int            `json:"id"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	ShareLink       string         `json:"share_link,omitempty"`
	MaxMembers      int            `json:"max_members"`
	IsPrivate       bool           `json:"is_private"`
	Theme           string         `json:"theme"`
	BackgroundImage string         `json:"background_image,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
	ExpiresAt       time.Time      `json:"expires_at"`
	Owner           *UserProfileDTO `json:"owner,omitempty"`
	MemberCount     int            `json:"member_count"`
	Members         []RoomMemberDTO `json:"members,omitempty"`
}

// RoomMemberDTO 自习室成员数据传输对象
type RoomMemberDTO struct {
	UserID      int       `json:"user_id"`
	Username    string    `json:"username,omitempty"`
	Avatar      string    `json:"avatar,omitempty"`
	IsAnonymous bool      `json:"is_anonymous"`
	Role        string    `json:"role"`
	Status      string    `json:"status"`
	JoinedAt    time.Time `json:"joined_at"`
}

// ToDetailDTO 将StudyRoom转换为StudyRoomDetailDTO
func (r *StudyRoom) ToDetailDTO(members []RoomMember, memberCount int) *StudyRoomDetailDTO {
	detail := &StudyRoomDetailDTO{
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
		MemberCount:     memberCount,
	}
	
	if r.Owner != nil {
		detail.Owner = r.Owner.ToProfileDTO()
	}
	
	if members != nil {
		detail.Members = make([]RoomMemberDTO, 0, len(members))
		for _, m := range members {
			memberDTO := RoomMemberDTO{
				UserID:      m.UserID,
				IsAnonymous: m.IsAnonymous,
				Role:        m.Role,
				Status:      m.Status,
				JoinedAt:    m.JoinedAt,
			}
			
			if m.User != nil && !m.IsAnonymous {
				memberDTO.Username = m.User.Username
				memberDTO.Avatar = m.User.Avatar
			}
			
			detail.Members = append(detail.Members, memberDTO)
		}
	}
	
	return detail
} 