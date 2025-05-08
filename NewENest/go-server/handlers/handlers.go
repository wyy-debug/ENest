package handlers

import (
	"github.com/gofiber/fiber/v2"
	"strings"
	"NewENest/go-server/middleware"
)

// Register 用户注册处理器
func Register(c *fiber.Ctx) error {
	// 解析请求体
	var registerData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&registerData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "请求格式无效",
		})
	}

	// 检查邮箱是否已被使用
	if registerData.Email == "test@example.com" || registerData.Email == "buddy@example.com" || registerData.Email == "focus@example.com" {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "邮箱已被占用",
		})
	}

	// 检查用户名是否已被使用
	if registerData.Username == "test_user" || registerData.Username == "study_buddy" || registerData.Username == "focusmaster" {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "用户名已被占用",
		})
	}

	// 返回成功响应
	return c.Status(201).JSON(fiber.Map{
		"code":    201,
		"message": "注册成功",
		"data": fiber.Map{
			"username": registerData.Username,
			"email":    registerData.Email,
		},
	})
}

// Login 用户登录处理器
func Login(c *fiber.Ctx) error {
	// 解析请求体
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "请求格式无效",
		})
	}

	// 检查是否是测试用户
	if loginData.Email == "test@example.com" && loginData.Password == "password123" {
		// 使用JWT中间件生成正确格式的token
		token, err := middleware.GenerateToken(1, "test@example.com")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"code":    500,
				"message": "生成token失败",
			})
		}
		
		// 返回成功响应
		return c.Status(200).JSON(fiber.Map{
			"token": token,
			"user": fiber.Map{
				"id":            1,
				"username":     "test_user",
				"email":        "test@example.com",
				"avatar":       "https://api.dicebear.com/7.x/avataaars/svg?seed=test_user",
				"signature":    "每天进步一点点",
				"studyDirection": "计算机科学",
				"totalStudyTime": 1200,
				"achievementPoints": 50,
			},
		})
	} else if loginData.Email == "buddy@example.com" && loginData.Password == "password123" {
		// 使用JWT中间件生成正确格式的token
		token, err := middleware.GenerateToken(2, "buddy@example.com")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"code":    500,
				"message": "生成token失败",
			})
		}
		
		// 另一个测试用户
		return c.Status(200).JSON(fiber.Map{
			"token": token,
			"user": fiber.Map{
				"id":            2,
				"username":     "study_buddy",
				"email":        "buddy@example.com",
				"avatar":       "https://api.dicebear.com/7.x/avataaars/svg?seed=study_buddy",
				"signature":    "一起学习吧！",
				"studyDirection": "数学",
				"totalStudyTime": 900,
				"achievementPoints": 30,
			},
		})
	}

	// 返回错误响应
	return c.Status(401).JSON(fiber.Map{
		"code":    401,
		"message": "邮箱或密码错误",
	})
}

// GetCurrentUser 获取当前用户信息处理器
func GetCurrentUser(c *fiber.Ctx) error {
	// 从请求头获取令牌
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"code":    401,
			"message": "未提供认证令牌",
		})
	}

	// 检查令牌格式
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(401).JSON(fiber.Map{
			"code":    401,
			"message": "认证令牌格式无效",
		})
	}

	// 获取令牌
	token := parts[1]

	// 根据令牌返回对应用户
	if token == "test_token_123456" {
		return c.Status(200).JSON(fiber.Map{
			"id":            1,
			"username":     "test_user",
			"email":        "test@example.com",
			"avatar":       "https://api.dicebear.com/7.x/avataaars/svg?seed=test_user",
			"signature":    "每天进步一点点",
			"studyDirection": "计算机科学",
			"totalStudyTime": 1200,
			"achievementPoints": 50,
		})
	} else if token == "buddy_token_123456" {
		return c.Status(200).JSON(fiber.Map{
			"id":            2,
			"username":     "study_buddy",
			"email":        "buddy@example.com",
			"avatar":       "https://api.dicebear.com/7.x/avataaars/svg?seed=study_buddy",
			"signature":    "一起学习吧！",
			"studyDirection": "数学",
			"totalStudyTime": 900,
			"achievementPoints": 30,
		})
	}

	// 返回错误响应
	return c.Status(401).JSON(fiber.Map{
		"code":    401,
		"message": "无效的认证令牌",
	})
}

// UpdateCurrentUser 更新当前用户信息处理器
func UpdateCurrentUser(c *fiber.Ctx) error {
	return c.Status(501).JSON(fiber.Map{
		"message": "更新用户信息功能待实现",
	})
}

// CreateTask 创建任务处理器
func CreateTask(c *fiber.Ctx) error {
	return c.Status(501).JSON(fiber.Map{
		"message": "创建任务功能待实现",
	})
}

// GetTasks 获取任务列表处理器
func GetTasks(c *fiber.Ctx) error {
	return c.Status(501).JSON(fiber.Map{
		"message": "获取任务列表功能待实现",
	})
}

// GetTask 获取任务详情处理器
func GetTask(c *fiber.Ctx) error {
	return c.Status(501).JSON(fiber.Map{
		"message": "获取任务详情功能待实现",
	})
}

// UpdateTask 更新任务处理器
func UpdateTask(c *fiber.Ctx) error {
	return c.Status(501).JSON(fiber.Map{
		"message": "更新任务功能待实现",
	})
}

// DeleteTask 删除任务处理器
func DeleteTask(c *fiber.Ctx) error {
	return c.Status(501).JSON(fiber.Map{
		"message": "删除任务功能待实现",
	})
}

// CreateStudyRecord 创建学习记录处理器
func CreateStudyRecord(c *fiber.Ctx) error {
	return c.Status(501).JSON(fiber.Map{
		"message": "创建学习记录功能待实现",
	})
}

// GetStudyRecords 获取学习记录列表处理器
func GetStudyRecords(c *fiber.Ctx) error {
	return c.Status(501).JSON(fiber.Map{
		"message": "获取学习记录列表功能待实现",
	})
}

// GetStudyRecord 获取学习记录详情处理器
func GetStudyRecord(c *fiber.Ctx) error {
	return c.Status(501).JSON(fiber.Map{
		"message": "获取学习记录详情功能待实现",
	})
}

// UpdateStudyRecord 更新学习记录处理器
func UpdateStudyRecord(c *fiber.Ctx) error {
	return c.Status(501).JSON(fiber.Map{
		"message": "更新学习记录功能待实现",
	})
}

// DeleteStudyRecord 删除学习记录处理器
func DeleteStudyRecord(c *fiber.Ctx) error {
	return c.Status(501).JSON(fiber.Map{
		"message": "删除学习记录功能待实现",
	})
}

// GetStudyStats 获取学习统计数据处理器
func GetStudyStats(c *fiber.Ctx) error {
	return c.Status(501).JSON(fiber.Map{
		"message": "获取学习统计数据功能待实现",
	})
}

// SearchUsers 搜索用户处理器
func SearchUsers(c *fiber.Ctx) error {
	// 获取搜索关键词
	keyword := c.Query("keyword")
	if keyword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "搜索关键词不能为空",
		})
	}
	
	// 模拟数据 - 实际应用中应该从数据库查询
	var results []fiber.Map
	
	// 如果搜索buddy，返回buddy用户
	if keyword == "buddy" || keyword == "study_buddy" || keyword == "study" {
		results = append(results, fiber.Map{
			"id":              2,
			"username":        "study_buddy",
			"email":           "buddy@example.com",
			"avatar":          "https://api.dicebear.com/7.x/avataaars/svg?seed=study_buddy",
			"signature":       "一起学习吧！",
			"study_direction": "数学",
			"is_friend":       false,
		})
	}
	
	// 如果搜索test，返回test用户
	if keyword == "test" || keyword == "test_user" {
		results = append(results, fiber.Map{
			"id":              1,
			"username":        "test_user",
			"email":           "test@example.com", 
			"avatar":          "https://api.dicebear.com/7.x/avataaars/svg?seed=test_user",
			"signature":       "每天进步一点点",
			"study_direction": "计算机科学",
			"is_friend":       false,
		})
	}
	
	// 如果搜索admin，返回admin用户
	if keyword == "admin" || keyword == "admin_user" {
		results = append(results, fiber.Map{
			"id":              3,
			"username":        "admin_user",
			"email":           "admin@example.com",
			"avatar":          "https://api.dicebear.com/7.x/avataaars/svg?seed=admin_user",
			"signature":       "系统管理员",
			"study_direction": "系统管理",
			"is_friend":       false,
		})
	}
	
	// 返回结果
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "搜索成功",
		"data":    results,
		"total":   len(results),
	})
}

// HandleNotFound 处理404请求的处理器
func HandleNotFound(c *fiber.Ctx, message string) error {
	return c.Status(404).JSON(fiber.Map{
		"message": message,
	})
} 