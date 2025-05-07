package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"

	"NewENest/go-server/routes"
)

func main() {
	// 初始化日志
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// 创建Fiber应用
	app := fiber.New(fiber.Config{
		AppName:      "NewENest API",
		ServerHeader: "Fiber",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
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
	})

	// 设置路由
	routes.SetupRoutes(app)

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
		port = "8080"
	}

	logger.Info().Msgf("服务器启动在端口: %s", port)
	if err := app.Listen(":" + port); err != nil {
		logger.Fatal().Err(err).Msg("服务器启动失败")
	}
} 