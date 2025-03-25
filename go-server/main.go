package main

import (
	"log"

	"go-server/database"
	"go-server/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
)

func main() {
	// 初始化数据库连接
	if err := database.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer database.CloseDB()

	// 创建一个新的Fiber实例
	app := fiber.New(fiber.Config{
		AppName: "E-StudyRoom API Server",
	})

	// 使用中间件
	app.Use(logger.New()) // 日志中间件
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, Sec-WebSocket-Key, Sec-WebSocket-Protocol, Sec-WebSocket-Version, Sec-WebSocket-Extensions",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		ExposeHeaders:    "Sec-WebSocket-Accept",
	})) // CORS中间件

	// WebSocket路由
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/ws", websocket.New(handlers.HandleWebSocket, websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}))

	// 设置API路由
	// api := app.Group("/api")

	// // 健康检查路由
	// api.Get("/health", func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{
	// 		"status": "ok",
	// 	})
	// })

	// 启动服务器
	log.Fatal(app.Listen("0.0.0.0:3000"))
}
