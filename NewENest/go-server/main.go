package main

import (
	"NewENest/go-server/config"
	"NewENest/go-server/routes"
	"NewENest/go-server/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 创建 Fiber 应用
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"code":    code,
				"message": err.Error(),
			})
		},
	})

	// 配置中间件
	corsConfig := cors.Config{
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}

	// 如果允许所有源，则禁用AllowCredentials
	if len(cfg.CORS.AllowOrigins) == 1 && cfg.CORS.AllowOrigins[0] == "*" {
		corsConfig.AllowCredentials = false
		corsConfig.AllowOrigins = "*"
	} else {
		corsConfig.AllowOrigins = cfg.CORS.AllowOrigins[0]
	}

	app.Use(cors.New(corsConfig))
	app.Use(logger.New())

	// 注册认证中间件
	middleware.RegisterAuthMiddleware(app)

	// 连接数据库
	db, err := connectDB(&cfg.Database)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	defer db.Close()

	// 设置路由
	routes.SetupRoutes(app, db)

	// 启动服务器
	go func() {
		if err := app.Listen(cfg.Server.Address); err != nil {
			log.Fatalf("服务器启动失败: %v", err)
		}
	}()

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("正在关闭服务器...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("服务器关闭失败: %v", err)
	}
	log.Println("服务器已关闭")
}

// 连接到PostgreSQL数据库
func connectDB(config *config.DatabaseConfig) (*sqlx.DB, error) {
	// 使用DSN方法获取连接字符串
	dsn := config.DSN()

	// 连接数据库
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// 测试连接
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("数据库连接成功")
	return db, nil
} 