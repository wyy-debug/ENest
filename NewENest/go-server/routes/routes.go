package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"NewENest/go-server/handlers"
	"NewENest/go-server/middleware"
)

// SetupRoutes 配置所有API路由
func SetupRoutes(app *fiber.App) {
	// 全局中间件
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))

	// API版本前缀
	api := app.Group("/api/v1")

	// 公开路由
	api.Post("/auth/register", handlers.Register)
	api.Post("/auth/login", handlers.Login)

	// 受保护的路由
	protected := api.Group("", middleware.JWTMiddleware())

	// 用户相关路由
	protected.Get("/user", handlers.GetCurrentUser)
	protected.Put("/user", handlers.UpdateCurrentUser)

	// 任务相关路由
	tasks := protected.Group("/tasks")
	tasks.Post("/", handlers.CreateTask)
	tasks.Get("/", handlers.GetTasks)
	tasks.Get("/:id", handlers.GetTask)
	tasks.Put("/:id", handlers.UpdateTask)
	tasks.Delete("/:id", handlers.DeleteTask)

	// 学习记录相关路由
	records := protected.Group("/study-records")
	records.Post("/", handlers.CreateStudyRecord)
	records.Get("/", handlers.GetStudyRecords)
	records.Get("/:id", handlers.GetStudyRecord)
	records.Put("/:id", handlers.UpdateStudyRecord)
	records.Delete("/:id", handlers.DeleteStudyRecord)
	records.Get("/stats", handlers.GetStudyStats)

	// 404处理
	app.Use(func(c *fiber.Ctx) error {
		return handlers.NotFound(c, "未找到请求的资源")
	})
} 