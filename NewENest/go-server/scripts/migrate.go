package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Println("未找到.env文件，使用环境变量")
	}

	// 获取数据库连接参数
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPass := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "postgres")
	dbSSLMode := getEnv("DB_SSLMODE", "disable")

	// 构建DSN
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPass, dbName, dbSSLMode)

	// 连接数据库
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	defer db.Close()

	log.Println("数据库连接成功")

	// 迁移脚本路径
	migrationsDir := filepath.Join("database", "migrations")

	// 读取迁移脚本文件
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		log.Fatalf("读取迁移目录失败: %v", err)
	}

	// 筛选SQL文件并排序
	var migrations []string
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			migrations = append(migrations, file.Name())
		}
	}
	sort.Strings(migrations)

	// 创建migrations表（如果不存在）
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id SERIAL PRIMARY KEY,
			version VARCHAR(50) NOT NULL UNIQUE,
			applied_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatalf("创建migrations表失败: %v", err)
	}

	// 获取已应用的迁移
	var appliedMigrations []string
	err = db.Select(&appliedMigrations, "SELECT version FROM migrations ORDER BY id")
	if err != nil {
		log.Fatalf("获取已应用迁移失败: %v", err)
	}

	// 记录已应用迁移版本
	appliedMap := make(map[string]bool)
	for _, m := range appliedMigrations {
		appliedMap[m] = true
	}

	// 应用新的迁移
	for _, migration := range migrations {
		version := strings.TrimSuffix(migration, filepath.Ext(migration))
		if appliedMap[version] {
			log.Printf("迁移 %s 已应用，跳过", version)
			continue
		}

		// 读取迁移脚本
		scriptPath := filepath.Join(migrationsDir, migration)
		content, err := ioutil.ReadFile(scriptPath)
		if err != nil {
			log.Fatalf("读取迁移脚本 %s 失败: %v", migration, err)
		}

		// 开始事务
		tx, err := db.Beginx()
		if err != nil {
			log.Fatalf("开始事务失败: %v", err)
		}

		// 执行迁移脚本
		log.Printf("应用迁移: %s", version)
		_, err = tx.Exec(string(content))
		if err != nil {
			tx.Rollback()
			log.Fatalf("执行迁移 %s 失败: %v", version, err)
		}

		// 确保迁移记录
		if !strings.Contains(string(content), "INSERT INTO migrations") {
			_, err = tx.Exec("INSERT INTO migrations (version) VALUES ($1)", version)
			if err != nil {
				tx.Rollback()
				log.Fatalf("记录迁移版本 %s 失败: %v", version, err)
			}
		}

		// 提交事务
		err = tx.Commit()
		if err != nil {
			log.Fatalf("提交事务失败: %v", err)
		}

		log.Printf("迁移 %s 应用成功", version)
	}

	log.Println("所有迁移已成功应用")
}

// 辅助函数，用于获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
} 