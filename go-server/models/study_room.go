package models

import (
	"database/sql"
	"errors"
	"go-server/database"
	"time"
	"crypto/rand"
	"encoding/base64"
)

type StudyRoom struct {
	ID          int       `json:"id"`
	OwnerID     int       `json:"owner_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ShareLink   string    `json:"share_link,omitempty"`
	MaxMembers  int       `json:"max_members"`
	IsPrivate   bool      `json:"is_private"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiresAt   time.Time `json:"expires_at"`
}

type RoomMember struct {
	ID       int       `json:"id"`
	RoomID   int       `json:"room_id"`
	UserID   int       `json:"user_id"`
	JoinedAt time.Time `json:"joined_at"`
}

// generateShareLink 生成唯一的分享链接
func generateShareLink() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// CreateStudyRoom 创建新的自习室
func CreateStudyRoom(ownerID int, name string, maxMembers int, isPrivate bool, duration time.Duration) (*StudyRoom, error) {
	db := database.GetDB()
	var room StudyRoom

	// 生成分享链接
	shareLink, err := generateShareLink()
	if err != nil {
		return nil, err
	}

	// 设置过期时间
	expiresAt := time.Now().Add(duration)

	// 插入自习室记录
	err = db.QueryRow(
		`INSERT INTO study_rooms (owner_id, name, share_link, max_members, is_private, expires_at) 
		 VALUES ($1, $2, $3, $4, $5, $6) 
		 RETURNING id, owner_id, name, share_link, max_members, is_private, created_at, expires_at`,
		ownerID, name, shareLink, maxMembers, isPrivate, expiresAt,
	).Scan(&room.ID, &room.OwnerID, &room.Name, &room.ShareLink, &room.MaxMembers, &room.IsPrivate, &room.CreatedAt, &room.ExpiresAt)

	if err != nil {
		return nil, err
	}

	// 自动将创建者加入自习室
	_, err = JoinStudyRoom(room.ID, ownerID)
	if err != nil {
		return nil, err
	}

	return &room, nil
}

// GetStudyRoom 获取自习室信息
func GetStudyRoom(roomID int) (*StudyRoom, error) {
	db := database.GetDB()
	var room StudyRoom

	err := db.QueryRow(
		"SELECT id, owner_id, name, share_link, max_members, is_private, created_at, expires_at FROM study_rooms WHERE id = $1",
		roomID,
	).Scan(&room.ID, &room.OwnerID, &room.Name, &room.ShareLink, &room.MaxMembers, &room.IsPrivate, &room.CreatedAt, &room.ExpiresAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("study room not found")
	} else if err != nil {
		return nil, err
	}

	return &room, nil
}

// StudyRoomDetail 自习室详细信息，包括成员信息
type StudyRoomDetail struct {
	StudyRoom
	MemberCount int             `json:"member_count"`
	Members     []RoomMemberInfo `json:"members"`
}

// RoomMemberInfo 自习室成员信息
type RoomMemberInfo struct {
	UserID      int       `json:"user_id"`
	Username    string    `json:"username,omitempty"`
	IsAnonymous bool      `json:"is_anonymous"`
	JoinedAt    time.Time `json:"joined_at"`
}

// GetStudyRoomDetail 获取自习室详细信息，包括成员列表
func GetStudyRoomDetail(roomID int) (*StudyRoomDetail, error) {
	db := database.GetDB()

	// 获取自习室基本信息
	room, err := GetStudyRoom(roomID)
	if err != nil {
		return nil, err
	}

	detail := &StudyRoomDetail{
		StudyRoom: *room,
	}

	// 获取成员信息
	rows, err := db.Query(
		`SELECT rm.user_id, u.username, rm.is_anonymous, rm.joined_at 
		 FROM room_members rm 
		 LEFT JOIN users u ON rm.user_id = u.id 
		 WHERE rm.room_id = $1`,
		roomID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var member RoomMemberInfo
		var username sql.NullString
		err := rows.Scan(&member.UserID, &username, &member.IsAnonymous, &member.JoinedAt)
		if err != nil {
			return nil, err
		}

		// 如果用户选择匿名，则不返回用户名
		if !member.IsAnonymous && username.Valid {
			member.Username = username.String
		}

		detail.Members = append(detail.Members, member)
	}

	detail.MemberCount = len(detail.Members)
	return detail, nil
}

// GetStudyRoomByShareLink 通过分享链接获取自习室
func GetStudyRoomByShareLink(shareLink string) (*StudyRoom, error) {
	db := database.GetDB()
	var room StudyRoom

	err := db.QueryRow(
		"SELECT id, owner_id, name, share_link, max_members, is_private, created_at, expires_at FROM study_rooms WHERE share_link = $1",
		shareLink,
	).Scan(&room.ID, &room.OwnerID, &room.Name, &room.ShareLink, &room.MaxMembers, &room.IsPrivate, &room.CreatedAt, &room.ExpiresAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("study room not found")
	} else if err != nil {
		return nil, err
	}

	return &room, nil
}

// JoinStudyRoom 加入自习室
func JoinStudyRoom(roomID, userID int) (*RoomMember, error) {
	db := database.GetDB()

	// 检查自习室是否存在且未过期
	room, err := GetStudyRoom(roomID)
	if err != nil {
		return nil, err
	}

	if time.Now().After(room.ExpiresAt) {
		return nil, errors.New("study room has expired")
	}

	// 检查成员数量是否达到上限
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM room_members WHERE room_id = $1", roomID).Scan(&count)
	if err != nil {
		return nil, err
	}

	if count >= room.MaxMembers {
		return nil, errors.New("study room is full")
	}

	// 添加成员
	var member RoomMember
	err = db.QueryRow(
		`INSERT INTO room_members (room_id, user_id) 
		 VALUES ($1, $2) 
		 RETURNING id, room_id, user_id, joined_at`,
		roomID, userID,
	).Scan(&member.ID, &member.RoomID, &member.UserID, &member.JoinedAt)

	if err != nil {
		return nil, err
	}

	return &member, nil
}

// LeaveStudyRoom 离开自习室
func LeaveStudyRoom(roomID, userID int) error {
	db := database.GetDB()
	_, err := db.Exec("DELETE FROM room_members WHERE room_id = $1 AND user_id = $2", roomID, userID)
	return err
}

// DestroyStudyRoom 销毁自习室（仅房主可操作）
func DestroyStudyRoom(roomID, ownerID int) error {
	db := database.GetDB()

	// 验证操作者是否为房主
	room, err := GetStudyRoom(roomID)
	if err != nil {
		return err
	}

	if room.OwnerID != ownerID {
		return errors.New("only room owner can destroy the study room")
	}

	// 删除自习室（会级联删除所有成员记录）
	_, err = db.Exec("DELETE FROM study_rooms WHERE id = $1", roomID)
	return err
}