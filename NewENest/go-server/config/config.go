package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config 应用程序配置
type Config struct {
	// 服务器配置
	Server ServerConfig `split_words:"true"`
	
	// 数据库配置
	Database DatabaseConfig `split_words:"true"`
	
	// CORS配置
	CORS CORSConfig `split_words:"true"`
	
	// JWT配置
	JWT JWTConfig `split_words:"true"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	// 服务器监听地址
	Address string `default:":8080"`
	
	// 环境模式
	Environment string `default:"development"`
}

// DatabaseConfig 数据库配置结构
type DatabaseConfig struct {
	// 主机地址
	Host string `default:"localhost"`
	
	// 端口
	Port string `default:"5432"`
	
	// 用户名
	Username string `required:"true"`
	
	// 密码
	Password string `required:"true"`
	
	// 数据库名称
	Database string `required:"true"`
	
	// SSL模式
	SSLMode string `default:"disable" split_words:"true"`
	
	// 连接字符串
	ConnectionString string `ignored:"true"`
}

// DSN 返回数据库连接字符串
func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.Username, c.Password, c.Database, c.SSLMode,
	)
}

// GetDatabaseConfig 从环境变量获取数据库配置
func GetDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		Username: getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		Database: getEnv("DB_NAME", "enest"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// CORSConfig CORS配置
type CORSConfig struct {
	// 允许的源
	AllowOrigins []string `default:"*" split_words:"true"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	// 密钥
	Secret string `required:"true"`
	
	// 过期时间（秒）
	ExpirationSeconds int `default:"86400" split_words:"true"`
}

// LoadConfig 从环境变量加载配置
func LoadConfig() (*Config, error) {
	// 尝试从.env文件加载环境变量，如果文件不存在则忽略错误
	_ = godotenv.Load()
	
	var config Config
	err := envconfig.Process("NEST", &config)
	if err != nil {
		return nil, err
	}
	
	// 处理特殊情况: 如果CORS允许的源是字符串，则按逗号分割
	if len(config.CORS.AllowOrigins) == 1 && config.CORS.AllowOrigins[0] == "*" {
		// 默认允许所有源，保持不变
	} else if len(config.CORS.AllowOrigins) == 1 && strings.Contains(config.CORS.AllowOrigins[0], ",") {
		// 按逗号分割
		config.CORS.AllowOrigins = strings.Split(config.CORS.AllowOrigins[0], ",")
	}
	
	// 开发环境下使用默认值
	if config.Server.Environment == "development" {
		setDefaultsForDevelopment(&config)
	}
	
	// 设置数据库连接字符串
	config.Database.ConnectionString = config.Database.DSN()
	
	return &config, nil
}

// setDefaultsForDevelopment 为开发环境设置默认值
func setDefaultsForDevelopment(config *Config) {
	// 设置数据库默认值，如果它们没有通过环境变量设置
	if os.Getenv("NEST_DATABASE_USERNAME") == "" {
		config.Database.Username = "postgres"
	}
	
	if os.Getenv("NEST_DATABASE_PASSWORD") == "" {
		config.Database.Password = "postgres"
	}
	
	if os.Getenv("NEST_DATABASE_DATABASE") == "" {
		config.Database.Database = "nest"
	}
	
	// 设置JWT默认密钥，如果没有设置
	if os.Getenv("NEST_JWT_SECRET") == "" {
		config.JWT.Secret = "development_secret_key_do_not_use_in_production"
	}
} 