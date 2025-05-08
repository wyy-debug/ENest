package services

import (
	"errors"
	"time"
	
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"

	"NewENest/go-server/models"
)

// UserService 用户服务接口
type UserService interface {
	Register(dto models.UserRegisterDTO) (*models.AuthResponseDTO, error)
	Login(dto models.UserLoginDTO) (*models.AuthResponseDTO, error)
	GetCurrentUser(userID int) (*models.UserProfileDTO, error)
	UpdateUser(userID int, dto models.UserUpdateDTO) (*models.UserProfileDTO, error)
}

// UserServiceImpl 用户服务实现
type UserServiceImpl struct {
	userRepo models.UserRepository
	jwtSecret string
	jwtExpiration time.Duration
}

// NewUserService 创建新的用户服务
func NewUserService(userRepo models.UserRepository, jwtSecret string, jwtExpiration time.Duration) UserService {
	return &UserServiceImpl{
		userRepo:      userRepo,
		jwtSecret:     jwtSecret,
		jwtExpiration: jwtExpiration,
	}
}

// Register 用户注册
func (s *UserServiceImpl) Register(dto models.UserRegisterDTO) (*models.AuthResponseDTO, error) {
	// 检查用户名是否已存在
	existingUser, err := s.userRepo.FindByUsername(dto.Username)
	if err == nil && existingUser != nil {
		return nil, errors.New("用户名已被占用")
	}
	
	// 检查邮箱是否已存在
	existingUser, err = s.userRepo.FindByEmail(dto.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("邮箱已被占用")
	}
	
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}
	
	// 创建新用户
	newUser := &models.User{
		Username:     dto.Username,
		Email:        dto.Email,
		PasswordHash: string(hashedPassword),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	
	// 保存用户
	if err := s.userRepo.Create(newUser); err != nil {
		return nil, errors.New("创建用户失败")
	}
	
	// 生成令牌
	token, err := s.generateToken(newUser.ID)
	if err != nil {
		return nil, errors.New("生成令牌失败")
	}
	
	// 返回响应
	return &models.AuthResponseDTO{
		User:  newUser.ToProfileDTO(),
		Token: token,
	}, nil
}

// Login 用户登录
func (s *UserServiceImpl) Login(dto models.UserLoginDTO) (*models.AuthResponseDTO, error) {
	// 根据邮箱查找用户
	user, err := s.userRepo.FindByEmail(dto.Email)
	if err != nil || user == nil {
		return nil, errors.New("用户不存在或密码错误")
	}
	
	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(dto.Password)); err != nil {
		return nil, errors.New("用户不存在或密码错误")
	}
	
	// 生成令牌
	token, err := s.generateToken(user.ID)
	if err != nil {
		return nil, errors.New("生成令牌失败")
	}
	
	// 返回响应
	return &models.AuthResponseDTO{
		User:  user.ToProfileDTO(),
		Token: token,
	}, nil
}

// GetCurrentUser 获取当前用户信息
func (s *UserServiceImpl) GetCurrentUser(userID int) (*models.UserProfileDTO, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil || user == nil {
		return nil, errors.New("用户不存在")
	}
	
	return user.ToProfileDTO(), nil
}

// UpdateUser 更新用户信息
func (s *UserServiceImpl) UpdateUser(userID int, dto models.UserUpdateDTO) (*models.UserProfileDTO, error) {
	// 获取当前用户
	user, err := s.userRepo.FindByID(userID)
	if err != nil || user == nil {
		return nil, errors.New("用户不存在")
	}
	
	// 更新用户信息
	if dto.Username != "" {
		// 检查用户名是否已被占用
		existingUser, err := s.userRepo.FindByUsername(dto.Username)
		if err == nil && existingUser != nil && existingUser.ID != userID {
			return nil, errors.New("用户名已被占用")
		}
		user.Username = dto.Username
	}
	
	if dto.Avatar != "" {
		user.Avatar = dto.Avatar
	}
	
	if dto.Signature != "" {
		user.Signature = dto.Signature
	}
	
	if dto.StudyDirection != "" {
		user.StudyDirection = dto.StudyDirection
	}
	
	user.UpdatedAt = time.Now()
	
	// 保存更新
	if err := s.userRepo.Update(user); err != nil {
		return nil, errors.New("更新用户失败")
	}
	
	return user.ToProfileDTO(), nil
}

// generateToken 生成JWT令牌
func (s *UserServiceImpl) generateToken(userID int) (*models.UserTokenDTO, error) {
	// 设置过期时间
	expiresAt := time.Now().Add(s.jwtExpiration).Unix()
	
	// 创建JWT声明
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expiresAt,
	}
	
	// 生成令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, err
	}
	
	return &models.UserTokenDTO{
		Token:     signedToken,
		ExpiresAt: expiresAt,
	}, nil
} 