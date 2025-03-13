package main

import (
	"log"

	"go-server/database"
	"go-server/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	})) // CORS中间件

	// 设置路由
	api := app.Group("/api")

	// 健康检查路由
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// 认证相关路由
	auth := api.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)
	auth.Post("/logout", handlers.Logout)

	// 自习室相关路由
	room := api.Group("/room")
	room.Post("/create", handlers.CreateStudyRoom)
	room.Post("/join", handlers.JoinStudyRoom)
	room.Post("/join/:shareLink", handlers.JoinStudyRoomByShareLink)
	room.Post("/leave", handlers.LeaveStudyRoom)
	room.Post("/destroy", handlers.DestroyStudyRoom)

	// 好友系统相关路由
	friend := api.Group("/friend")
	friend.Post("/request", handlers.SendFriendRequest)
	friend.Post("/handle-request", handlers.HandleFriendRequest)
	friend.Get("/list", handlers.GetFriendList)
	friend.Delete("/delete", handlers.DeleteFriend)
	friend.Post("/contract/create", handlers.CreateFriendContract)
	friend.Post("/contract/terminate", handlers.TerminateFriendContract)
	friend.Get("/contracts", handlers.GetFriendContracts)

	// WebSocket聊天路由
	friend.Get("/ws", handlers.HandleChat)

	// 个人信息相关路由
	profile := api.Group("/profile")
	profile.Get("/info", handlers.GetUserProfile)
	profile.Put("/update", handlers.UpdateUserProfile)

	// 启动服务器
	log.Fatal(app.Listen(":3000"))
}
