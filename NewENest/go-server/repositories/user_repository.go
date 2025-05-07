package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	
	"golang.org/x/crypto/bcrypt"
	
	"github.com/jmoiron/sqlx"
	
	"NewENest/go-server/models"
)

// UserRepositoryImpl 用户数据访问实现
type UserRepositoryImpl struct {
	db *sqlx.DB
}

// NewUserRepository 创建用户数据访问实现实例
func NewUserRepository(db *sqlx.DB) models.UserRepository {
	return &UserRepositoryImpl{db: db}
}

// FindByID 通过ID查找用户
func (r *UserRepositoryImpl) FindByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user, `
		SELECT id, username, email, password_hash, avatar, signature, study_direction, 
		       total_study_time, achievement_points, created_at, updated_at 
		FROM users 
		WHERE id = $1
	`, id)
	
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}
	
	return &user, nil
}

// FindByUsername 通过用户名查找用户
func (r *UserRepositoryImpl) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user, `
		SELECT id, username, email, password_hash, avatar, signature, study_direction, 
		       total_study_time, achievement_points, created_at, updated_at 
		FROM users 
		WHERE username = $1
	`, username)
	
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}
	
	return &user, nil
}

// FindByEmail 通过邮箱查找用户
func (r *UserRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user, `
		SELECT id, username, email, password_hash, avatar, signature, study_direction, 
		       total_study_time, achievement_points, created_at, updated_at 
		FROM users 
		WHERE email = $1
	`, email)
	
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}
	
	return &user, nil
}

// Create 创建新用户
func (r *UserRepositoryImpl) Create(user *models.User) error {
	// 检查用户名是否已存在
	exists, err := r.checkExists("username", user.Username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("username already exists")
	}
	
	// 检查邮箱是否已存在
	exists, err = r.checkExists("email", user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email already exists")
	}
	
	// 执行插入
	err = r.db.QueryRow(`
		INSERT INTO users (username, email, password_hash, avatar, signature, study_direction) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id, created_at, updated_at
	`, user.Username, user.Email, user.PasswordHash, user.Avatar, user.Signature, user.StudyDirection).Scan(
		&user.ID, &user.CreatedAt, &user.UpdatedAt)
	
	return err
}

// Update 更新用户信息
func (r *UserRepositoryImpl) Update(user *models.User) error {
	result, err := r.db.Exec(`
		UPDATE users 
		SET username = $1, email = $2, avatar = $3, signature = $4, study_direction = $5, 
		    total_study_time = $6, achievement_points = $7, updated_at = CURRENT_TIMESTAMP 
		WHERE id = $8
	`, user.Username, user.Email, user.Avatar, user.Signature, user.StudyDirection,
		user.TotalStudyTime, user.AchievementPoints, user.ID)
	
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	
	return nil
}

// Delete 删除用户
func (r *UserRepositoryImpl) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	
	return nil
}

// Authenticate 用户认证
func (r *UserRepositoryImpl) Authenticate(username, password string) (*models.User, error) {
	user, err := r.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	
	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}
	
	return user, nil
}

// ChangePassword 修改密码
func (r *UserRepositoryImpl) ChangePassword(userID int, newPassword string) error {
	// 生成密码哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	
	result, err := r.db.Exec(`
		UPDATE users 
		SET password_hash = $1, updated_at = CURRENT_TIMESTAMP 
		WHERE id = $2
	`, string(hashedPassword), userID)
	
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	
	return nil
}

// CreateSession 创建会话
func (r *UserRepositoryImpl) CreateSession(userID int, token string, expiresAt time.Time) (*models.Session, error) {
	var session models.Session
	err := r.db.QueryRow(`
		INSERT INTO sessions (user_id, token, expires_at) 
		VALUES ($1, $2, $3) 
		RETURNING id, user_id, token, expires_at, created_at
	`, userID, token, expiresAt).Scan(
		&session.ID, &session.UserID, &session.Token, &session.ExpiresAt, &session.CreatedAt)
	
	if err != nil {
		return nil, err
	}
	
	return &session, nil
}

// GetSessionByToken 通过Token获取会话
func (r *UserRepositoryImpl) GetSessionByToken(token string) (*models.Session, error) {
	var session models.Session
	err := r.db.Get(&session, `
		SELECT id, user_id, token, expires_at, created_at 
		FROM sessions 
		WHERE token = $1 AND expires_at > CURRENT_TIMESTAMP
	`, token)
	
	if err == sql.ErrNoRows {
		return nil, errors.New("session not found or expired")
	} else if err != nil {
		return nil, err
	}
	
	return &session, nil
}

// DeleteSession 删除会话
func (r *UserRepositoryImpl) DeleteSession(token string) error {
	_, err := r.db.Exec("DELETE FROM sessions WHERE token = $1", token)
	return err
}

// UpdateStudyTime 更新用户学习时间
func (r *UserRepositoryImpl) UpdateStudyTime(userID int, additionalMinutes int) error {
	result, err := r.db.Exec(`
		UPDATE users 
		SET total_study_time = total_study_time + $1, updated_at = CURRENT_TIMESTAMP 
		WHERE id = $2
	`, additionalMinutes, userID)
	
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	
	return nil
}

// UpdateAchievementPoints 更新用户成就点数
func (r *UserRepositoryImpl) UpdateAchievementPoints(userID int, points int) error {
	result, err := r.db.Exec(`
		UPDATE users 
		SET achievement_points = achievement_points + $1, updated_at = CURRENT_TIMESTAMP 
		WHERE id = $2
	`, points, userID)
	
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	
	return nil
}

// checkExists 检查指定字段的值是否已存在
func (r *UserRepositoryImpl) checkExists(field, value string) (bool, error) {
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM users WHERE %s = $1)", field)
	var exists bool
	err := r.db.QueryRow(query, value).Scan(&exists)
	return exists, err
} 