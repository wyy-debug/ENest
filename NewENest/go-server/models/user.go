package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type User struct {
	ID               int       `json:"id" db:"id"`
	Username         string    `json:"username" db:"username"`
	Email            string    `json:"email" db:"email"`
	PasswordHash     string    `json:"-" db:"password_hash"`
	Avatar           string    `json:"avatar,omitempty" db:"avatar"`
	Signature        string    `json:"signature,omitempty" db:"signature"`
	StudyDirection   string    `json:"study_direction,omitempty" db:"study_direction"`
	TotalStudyTime   int       `json:"total_study_time" db:"total_study_time"`
	AchievementPoints int      `json:"achievement_points" db:"achievement_points"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

// UserProfileDTO 用户概要信息DTO
type UserProfileDTO struct {
	ID               int       `json:"id"`
	Username         string    `json:"username"`
	Avatar           string    `json:"avatar,omitempty"`
	Signature        string    `json:"signature,omitempty"`
	StudyDirection   string    `json:"study_direction,omitempty"`
	TotalStudyTime   int       `json:"total_study_time"`
	AchievementPoints int      `json:"achievement_points"`
}

// ToProfileDTO 将User转换为UserProfileDTO
func (u *User) ToProfileDTO() *UserProfileDTO {
	if u == nil {
		return nil
	}
	
	return &UserProfileDTO{
		ID:               u.ID,
		Username:         u.Username,
		Avatar:           u.Avatar,
		Signature:        u.Signature,
		StudyDirection:   u.StudyDirection,
		TotalStudyTime:   u.TotalStudyTime,
		AchievementPoints: u.AchievementPoints,
	}
}

// UserRepository 用户仓库接口
type UserRepository interface {
	FindByID(id int) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
	UpdateStudyTime(userID int, minutes int) error
	UpdateAchievementPoints(userID int, points int) error
}

// UserLoginDTO 用户登录DTO
type UserLoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// UserRegisterDTO 用户注册DTO
type UserRegisterDTO struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// UserUpdateDTO 用户更新DTO
type UserUpdateDTO struct {
	Username       string `json:"username,omitempty" validate:"omitempty,min=3,max=50"`
	Avatar         string `json:"avatar,omitempty"`
	Signature      string `json:"signature,omitempty"`
	StudyDirection string `json:"study_direction,omitempty"`
}

// UserTokenDTO 用户令牌DTO
type UserTokenDTO struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

// AuthResponseDTO 认证响应DTO
type AuthResponseDTO struct {
	User  *UserProfileDTO `json:"user"`
	Token *UserTokenDTO   `json:"token"`
}

// UserResponse 是返回给客户端的用户信息
type UserResponse struct {
	ID                int       `json:"id"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	Avatar            string    `json:"avatar"`
	Signature         string    `json:"signature"`
	StudyDirection    string    `json:"study_direction"`
	TotalStudyTime    int       `json:"total_study_time"`
	AchievementPoints int       `json:"achievement_points"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// UserCreate 表示创建用户的请求数据
type UserCreate struct {
	Username  string `json:"username" validate:"required,min=2,max=50"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
}

// UserLogin 表示用户登录的请求数据
type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// UserUpdate 表示更新用户信息的请求数据
type UserUpdate struct {
	Username       string `json:"username" validate:"omitempty,min=2,max=50"`
	Avatar         string `json:"avatar"`
	Signature      string `json:"signature" validate:"omitempty,max=200"`
	StudyDirection string `json:"study_direction"`
}

// HashPassword 创建密码的哈希
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash 检查密码是否与哈希匹配
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ToResponse 将用户模型转换为响应形式
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:                u.ID,
		Username:          u.Username,
		Email:             u.Email,
		Avatar:            u.Avatar,
		Signature:         u.Signature,
		StudyDirection:    u.StudyDirection,
		TotalStudyTime:    u.TotalStudyTime,
		AchievementPoints: u.AchievementPoints,
		CreatedAt:         u.CreatedAt,
		UpdatedAt:         u.UpdatedAt,
	}
}

// Session 会话模型
type Session struct {
	ID        int       `json:"id" db:"id"`
}