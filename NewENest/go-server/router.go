package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	
	"NewENest/go-server/handlers"
)

// setupRouter 配置并返回一个配置好的路由器
func setupRouter() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "ENest API",
	})

	// 中间件
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))

	// API版本前缀
	api := app.Group("/api/v1")

	// 健康检查
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status": "ok",
			"version": "1.0.0",
		})
	})

	// 自习室路由
	studyRooms := api.Group("/study-rooms")
	studyRooms.Get("/", handlers.GetStudyRooms)        // 获取自习室列表
	studyRooms.Get("/:id", handlers.GetStudyRoom)      // 获取自习室详情
	studyRooms.Post("/", handlers.CreateStudyRoom)     // 创建自习室
	studyRooms.Post("/join", handlers.JoinStudyRoom)   // 加入自习室

	// 404处理
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code": 404,
			"message": "路径不存在",
		})
	})

	return app
} 