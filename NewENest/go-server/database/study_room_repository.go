package database

import (
	"fmt"
	"log"
	"time"
	"github.com/jmoiron/sqlx"
)

// StudyRoomRepository 自习室数据访问层
type StudyRoomRepository struct {
	db *sqlx.DB
}

// NewStudyRoomRepository 创建新的自习室数据访问层实例
func NewStudyRoomRepository(db *sqlx.DB) *StudyRoomRepository {
	return &StudyRoomRepository{db: db}
}

// StudyRoom 自习室数据结构
type StudyRoom struct {
	ID              int       `db:"id" json:"id"`
	OwnerID         int       `db:"owner_id" json:"owner_id"`
	Name            string    `db:"name" json:"name"`
	Description     string    `db:"description" json:"description"`
	ShareLink       string    `db:"share_link" json:"share_link"`
	MaxMembers      int       `db:"max_members" json:"max_members"`
	IsPrivate       bool      `db:"is_private" json:"is_private"`
	Theme           string    `db:"theme" json:"theme"`
	BackgroundImage string    `db:"background_image" json:"background_image"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	ExpiresAt       time.Time `db:"expires_at" json:"expires_at"`
}

// StudyRoomMember 自习室成员数据结构
type StudyRoomMember struct {
	ID          int       `db:"id" json:"id"`
	RoomID      int       `db:"room_id" json:"room_id"`
	UserID      int       `db:"user_id" json:"user_id"`
	IsAnonymous bool      `db:"is_anonymous" json:"is_anonymous"`
	Role        string    `db:"role" json:"role"`
	Status      string    `db:"status" json:"status"`
	JoinedAt    time.Time `db:"joined_at" json:"joined_at"`
}

// User 用户基本信息结构
type User struct {
	ID        int    `db:"id" json:"id"`
	Username  string `db:"username" json:"username"`
	Avatar    string `db:"avatar" json:"avatar"`
}

// CreateStudyRoom 创建新自习室
func (r *StudyRoomRepository) CreateStudyRoom(room *StudyRoom) (int, error) {
	query := `
		INSERT INTO study_rooms 
		(owner_id, name, description, share_link, max_members, is_private, theme, background_image, created_at, expires_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id
	`
	
	var id int
	err := r.db.QueryRowx(
		query,
		room.OwnerID,
		room.Name,
		room.Description,
		room.ShareLink,
		room.MaxMembers,
		room.IsPrivate,
		room.Theme,
		room.BackgroundImage,
		room.CreatedAt,
		room.ExpiresAt,
	).Scan(&id)
	
	if err != nil {
		log.Printf("创建自习室失败: %v", err)
		return 0, fmt.Errorf("创建自习室失败: %w", err)
	}
	
	return id, nil
}

// GetStudyRoomByID 根据ID获取自习室
func (r *StudyRoomRepository) GetStudyRoomByID(id int) (*StudyRoom, error) {
	query := `
		SELECT id, owner_id, name, description, share_link, max_members, is_private, theme, background_image, created_at, expires_at
		FROM study_rooms
		WHERE id = $1
	`
	
	var room StudyRoom
	err := r.db.Get(&room, query, id)
	if err != nil {
		log.Printf("获取自习室失败: %v", err)
		return nil, fmt.Errorf("获取自习室失败: %w", err)
	}
	
	return &room, nil
}

// GetStudyRoomByShareLink 根据分享链接获取自习室
func (r *StudyRoomRepository) GetStudyRoomByShareLink(shareLink string) (*StudyRoom, error) {
	query := `
		SELECT id, owner_id, name, description, share_link, max_members, is_private, theme, background_image, created_at, expires_at
		FROM study_rooms
		WHERE share_link = $1
	`
	
	var room StudyRoom
	err := r.db.Get(&room, query, shareLink)
	if err != nil {
		log.Printf("获取自习室失败: %v", err)
		return nil, fmt.Errorf("获取自习室失败: %w", err)
	}
	
	return &room, nil
}

// GetAllStudyRooms 获取所有自习室列表，带分页
func (r *StudyRoomRepository) GetAllStudyRooms(page, pageSize int) ([]StudyRoom, int, error) {
	// 获取总数
	var total int
	countQuery := `SELECT COUNT(*) FROM study_rooms`
	err := r.db.Get(&total, countQuery)
	if err != nil {
		return nil, 0, fmt.Errorf("获取自习室总数失败: %w", err)
	}
	
	// 获取列表
	query := `
		SELECT id, owner_id, name, description, share_link, max_members, is_private, theme, background_image, created_at, expires_at
		FROM study_rooms
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	
	offset := (page - 1) * pageSize
	var rooms []StudyRoom
	err = r.db.Select(&rooms, query, pageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("获取自习室列表失败: %w", err)
	}
	
	return rooms, total, nil
}

// GetUserByID 根据用户ID获取用户基本信息
func (r *StudyRoomRepository) GetUserByID(id int) (*User, error) {
	query := `SELECT id, username, avatar FROM users WHERE id = $1`
	
	var user User
	err := r.db.Get(&user, query, id)
	if err != nil {
		return nil, fmt.Errorf("获取用户信息失败: %w", err)
	}
	
	return &user, nil
}

// AddMemberToStudyRoom 添加成员到自习室
func (r *StudyRoomRepository) AddMemberToStudyRoom(member *StudyRoomMember) error {
	query := `
		INSERT INTO room_members
		(room_id, user_id, is_anonymous, role, status, joined_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (room_id, user_id) DO UPDATE
		SET is_anonymous = $3, role = $4, status = $5
	`
	
	_, err := r.db.Exec(
		query,
		member.RoomID,
		member.UserID,
		member.IsAnonymous,
		member.Role,
		member.Status,
		member.JoinedAt,
	)
	
	if err != nil {
		return fmt.Errorf("添加成员失败: %w", err)
	}
	
	return nil
}

// GetStudyRoomMembers 获取自习室所有成员
func (r *StudyRoomRepository) GetStudyRoomMembers(roomID int) ([]map[string]interface{}, error) {
	query := `
		SELECT 
			rm.user_id, u.username, u.avatar, 
			rm.is_anonymous, rm.role, rm.status, rm.joined_at
		FROM room_members rm
		JOIN users u ON rm.user_id = u.id
		WHERE rm.room_id = $1
		ORDER BY rm.joined_at ASC
	`
	
	rows, err := r.db.Queryx(query, roomID)
	if err != nil {
		return nil, fmt.Errorf("获取自习室成员失败: %w", err)
	}
	defer rows.Close()
	
	var members []map[string]interface{}
	for rows.Next() {
		member := make(map[string]interface{})
		err := rows.MapScan(member)
		if err != nil {
			return nil, fmt.Errorf("解析成员数据失败: %w", err)
		}
		members = append(members, member)
	}
	
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("获取成员列表错误: %w", err)
	}
	
	return members, nil
}

// CountStudyRoomMembers 获取自习室成员数量
func (r *StudyRoomRepository) CountStudyRoomMembers(roomID int) (int, error) {
	query := `SELECT COUNT(*) FROM room_members WHERE room_id = $1`
	
	var count int
	err := r.db.Get(&count, query, roomID)
	if err != nil {
		return 0, fmt.Errorf("获取成员数量失败: %w", err)
	}
	
	return count, nil
}

// UpdateStudyRoom 更新自习室信息
func (r *StudyRoomRepository) UpdateStudyRoom(room *StudyRoom) error {
	query := `
		UPDATE study_rooms
		SET 
			name = $1,
			description = $2,
			max_members = $3,
			is_private = $4,
			theme = $5,
			background_image = $6
		WHERE id = $7
	`
	
	_, err := r.db.Exec(
		query,
		room.Name,
		room.Description,
		room.MaxMembers,
		room.IsPrivate,
		room.Theme,
		room.BackgroundImage,
		room.ID,
	)
	
	if err != nil {
		return fmt.Errorf("更新自习室失败: %w", err)
	}
	
	return nil
}

// RemoveMemberFromStudyRoom 从自习室移除成员
func (r *StudyRoomRepository) RemoveMemberFromStudyRoom(roomID, userID int) error {
	query := `DELETE FROM room_members WHERE room_id = $1 AND user_id = $2`
	
	_, err := r.db.Exec(query, roomID, userID)
	if err != nil {
		return fmt.Errorf("移除成员失败: %w", err)
	}
	
	return nil
}

// UpdateMemberStatus 更新成员状态
func (r *StudyRoomRepository) UpdateMemberStatus(roomID, userID int, status string) error {
	query := `UPDATE room_members SET status = $1 WHERE room_id = $2 AND user_id = $3`
	
	_, err := r.db.Exec(query, status, roomID, userID)
	if err != nil {
		return fmt.Errorf("更新成员状态失败: %w", err)
	}
	
	return nil
} 