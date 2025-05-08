package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/joho/godotenv"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"NewENest/go-server/config"
	"NewENest/go-server/routes"
)

func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Println("未找到.env文件，使用环境变量")
	}

	// 加载配置
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化日志
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// 连接数据库
	db, err := connectDB(appConfig.Database)
	if err != nil {
		logger.Fatal().Err(err).Msg("数据库连接失败")
	}
	defer db.Close()

	// 创建Fiber应用
	app := fiber.New(fiber.Config{
		AppName:               "ENest API",
		ServerHeader:          "Fiber",
		ErrorHandler:          func(c *fiber.Ctx, err error) error {
			// 默认错误处理
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"code":    code,
				"message": err.Error(),
			})
		},
		DisableStartupMessage: false,
	})

	// 设置路由
	routes.SetupRoutes(app, db)

	// 优雅关闭
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		logger.Info().Msg("服务器正在关闭...")
		if err := app.Shutdown(); err != nil {
			logger.Fatal().Err(err).Msg("服务器关闭失败")
		}
	}()

	// 启动服务器
	port := os.Getenv("PORT")
	if port == "" {
		// 使用配置中的地址，如果没有设置环境变量PORT
		address := appConfig.Server.Address
		logger.Info().Msgf("服务器启动在地址: %s", address)
		if err := app.Listen(address); err != nil {
			logger.Fatal().Err(err).Msg("服务器启动失败")
		}
	} else {
		// 使用环境变量中的端口
		logger.Info().Msgf("服务器启动在端口: %s", port)
		if err := app.Listen(":" + port); err != nil {
			logger.Fatal().Err(err).Msg("服务器启动失败")
		}
	}
}

// 连接到PostgreSQL数据库
func connectDB(config config.DatabaseConfig) (*sqlx.DB, error) {
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

	log.Println("Connected to database successfully")
	return db, nil
} 