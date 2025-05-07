package services

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"
	
	"golang.org/x/crypto/bcrypt"
	
	"NewENest/go-server/models"
)

// UserService 用户服务接口
type UserService interface {
	// 用户管理
	Register(dto models.NewUserCreateDTO) (*models.User, error)
	Login(username, password string) (*models.User, string, error)
	Logout(token string) error
	GetUserProfile(userID int) (*models.UserProfileDTO, error)
	UpdateUserProfile(userID int, dto models.UserUpdateDTO) error
	ChangePassword(userID int, oldPassword, newPassword string) error
	
	// 会话管理
	ValidateToken(token string) (*models.User, error)
}

// UserServiceImpl 用户服务实现
type UserServiceImpl struct {
	userRepo models.UserRepository
}

// NewUserService 创建用户服务实例
func NewUserService(userRepo models.UserRepository) UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

// Register 用户注册
func (s *UserServiceImpl) Register(dto models.NewUserCreateDTO) (*models.User, error) {
	// 验证用户名和邮箱是否已存在
	_, err := s.userRepo.FindByUsername(dto.Username)
	if err == nil {
		return nil, errors.New("username already exists")
	}
	
	_, err = s.userRepo.FindByEmail(dto.Email)
	if err == nil {
		return nil, errors.New("email already exists")
	}
	
	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	
	// 创建用户对象
	user := &models.User{
		Username:     dto.Username,
		Email:        dto.Email,
		PasswordHash: string(hashedPassword),
	}
	
	// 保存用户信息
	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	
	return user, nil
}

// Login 用户登录
func (s *UserServiceImpl) Login(username, password string) (*models.User, string, error) {
	// 验证用户名和密码
	user, err := s.userRepo.Authenticate(username, password)
	if err != nil {
		return nil, "", err
	}
	
	// 生成会话令牌
	tokenBytes := make([]byte, 32)
	_, err = rand.Read(tokenBytes)
	if err != nil {
		return nil, "", err
	}
	token := base64.URLEncoding.EncodeToString(tokenBytes)
	
	// 设置会话过期时间（默认7天）
	expiresAt := time.Now().Add(7 * 24 * time.Hour)
	
	// 创建会话
	_, err = s.userRepo.CreateSession(user.ID, token, expiresAt)
	if err != nil {
		return nil, "", err
	}
	
	return user, token, nil
}

// Logout 用户登出
func (s *UserServiceImpl) Logout(token string) error {
	return s.userRepo.DeleteSession(token)
}

// GetUserProfile 获取用户资料
func (s *UserServiceImpl) GetUserProfile(userID int) (*models.UserProfileDTO, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	
	return user.ToProfileDTO(), nil
}

// UpdateUserProfile 更新用户资料
func (s *UserServiceImpl) UpdateUserProfile(userID int, dto models.UserUpdateDTO) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}
	
	// 更新用户信息
	if dto.Username != nil && *dto.Username != user.Username {
		// 检查新用户名是否已存在
		_, err := s.userRepo.FindByUsername(*dto.Username)
		if err == nil {
			return errors.New("username already exists")
		}
		user.Username = *dto.Username
	}
	
	if dto.Signature != nil {
		user.Signature = *dto.Signature
	}
	
	if dto.StudyDirection != nil {
		user.StudyDirection = *dto.StudyDirection
	}
	
	if dto.Avatar != nil {
		user.Avatar = *dto.Avatar
	}
	
	return s.userRepo.Update(user)
}

// ChangePassword 修改密码
func (s *UserServiceImpl) ChangePassword(userID int, oldPassword, newPassword string) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}
	
	// 验证旧密码
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(oldPassword))
	if err != nil {
		return errors.New("invalid old password")
	}
	
	// 修改密码
	return s.userRepo.ChangePassword(userID, newPassword)
}

// ValidateToken 验证令牌有效性
func (s *UserServiceImpl) ValidateToken(token string) (*models.User, error) {
	session, err := s.userRepo.GetSessionByToken(token)
	if err != nil {
		return nil, err
	}
	
	// 检查会话是否过期
	if time.Now().After(session.ExpiresAt) {
		s.userRepo.DeleteSession(token)
		return nil, errors.New("session expired")
	}
	
	// 获取用户信息
	user, err := s.userRepo.FindByID(session.UserID)
	if err != nil {
		return nil, err
	}
	
	return user, nil
} 