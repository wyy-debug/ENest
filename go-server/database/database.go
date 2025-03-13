package database

import (
	"database/sql"
	"fmt"
	"go-server/config"
	_ "github.com/lib/pq"
)

var db *sql.DB

// InitDB 初始化数据库连接
func InitDB() error {
	config := config.GetDatabaseConfig()
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
		config.SSLMode,
	)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	return nil
}

// GetDB 返回数据库连接实例
func GetDB() *sql.DB {
	return db
}

// CloseDB 关闭数据库连接
func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}