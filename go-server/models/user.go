package models

import (
	"database/sql"
	"errors"
	"go-server/database"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int       `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	PasswordHash   string    `json:"password_hash,omitempty"`
	Avatar         string    `json:"avatar,omitempty"`
	Signature      string    `json:"signature,omitempty"`
	StudyDirection string    `json:"study_direction,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Session struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateUser 创建新用户
func CreateUser(username, email, password string) (*User, error) {
	// 检查用户名是否已存在
	if exists, _ := checkUserExists("username", username); exists {
		return nil, errors.New("username already exists")
	}

	// 检查邮箱是否已存在
	if exists, _ := checkUserExists("email", email); exists {
		return nil, errors.New("email already exists")
	}

	// 生成密码哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 插入新用户
	db := database.GetDB()
	var user User
	err = db.QueryRow(
		`INSERT INTO users (username, email, password_hash, avatar, signature, study_direction) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id, username, email, password_hash, avatar, signature, study_direction, created_at, updated_at`,
		username, email, string(hashedPassword), "", "", "",
	).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Avatar, &user.Signature, &user.StudyDirection, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// AuthenticateUser 验证用户登录
func AuthenticateUser(username, password string) (*User, error) {
	db := database.GetDB()
	var user User
	err := db.QueryRow(
		"SELECT id, username, email, password_hash, created_at, updated_at FROM users WHERE username = $1",
		username,
	).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}

// CreateSession 创建新的会话
func CreateSession(userID int, token string, expiresAt time.Time) (*Session, error) {
	db := database.GetDB()
	var session Session
	err := db.QueryRow(
		`INSERT INTO sessions (user_id, token, expires_at) 
		VALUES ($1, $2, $3) 
		RETURNING id, user_id, token, expires_at, created_at`,
		userID, token, expiresAt,
	).Scan(&session.ID, &session.UserID, &session.Token, &session.ExpiresAt, &session.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &session, nil
}

// GetSessionByToken 通过token获取会话
func GetSessionByToken(token string) (*Session, error) {
	db := database.GetDB()
	var session Session
	err := db.QueryRow(
		"SELECT id, user_id, token, expires_at, created_at FROM sessions WHERE token = $1",
		token,
	).Scan(&session.ID, &session.UserID, &session.Token, &session.ExpiresAt, &session.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("session not found")
	} else if err != nil {
		return nil, err
	}

	return &session, nil
}

// DeleteSession 删除会话
func DeleteSession(token string) error {
	db := database.GetDB()
	_, err := db.Exec("DELETE FROM sessions WHERE token = $1", token)
	return err
}

// UpdateUserProfile 更新用户个人信息
func UpdateUserProfile(userID int, updates map[string]string) error {
	db := database.GetDB()

	// 如果要更新用户名，先检查是否已存在
	if username, ok := updates["username"]; ok {
		if exists, _ := checkUserExists("username", username); exists {
			return errors.New("username already exists")
		}
	}

	// 构建更新语句
	setClauses := make([]string, 0)
	values := make([]interface{}, 0)
	valueIndex := 1

	for field, value := range updates {
		switch field {
		case "username", "signature", "study_direction":
			setClauses = append(setClauses, field+" = $"+strconv.Itoa(valueIndex))
			values = append(values, value)
			valueIndex++
		}
	}

	if len(setClauses) == 0 {
		return errors.New("no valid fields to update")
	}

	// 添加更新时间
	setClauses = append(setClauses, "updated_at = CURRENT_TIMESTAMP")

	// 构建并执行SQL语句
	query := "UPDATE users SET " + strings.Join(setClauses, ", ") + " WHERE id = $" + string(valueIndex)
	values = append(values, userID)

	result, err := db.Exec(query, values...)
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

// GetUserByID 通过用户ID获取用户信息
func GetUserByID(userID int) (*User, error) {
	db := database.GetDB()
	var user User
	err := db.QueryRow(
		"SELECT id, username, email, avatar, signature, study_direction, created_at, updated_at FROM users WHERE id = $1",
		userID,
	).Scan(&user.ID, &user.Username, &user.Email, &user.Avatar, &user.Signature, &user.StudyDirection, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

// checkUserExists 检查用户是否已存在
func checkUserExists(field, value string) (bool, error) {
	db := database.GetDB()
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE " + field + " = $1)"
	err := db.QueryRow(query, value).Scan(&exists)
	return exists, err
}
