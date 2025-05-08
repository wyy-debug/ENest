package handlers

import (
	"NewENest/go-server/utils"
	"github.com/gofiber/fiber/v2"
)

// GetStudyRooms 获取自习室列表处理器
func GetStudyRooms(c *fiber.Ctx) error {
	// 返回示例自习室数据
	studyRooms := []fiber.Map{
		{
			"id": 1,
			"name": "编程自习室",
			"description": "专注于编程学习的自习室",
			"shareLink": "code_room_123",
			"maxMembers": 20,
			"isPrivate": false,
			"theme": "coding",
			"backgroundImage": "https://images.unsplash.com/photo-1605379399642-870262d3d051",
			"createdAt": "2025-05-01T10:00:00Z",
			"expiresAt": "2025-05-08T10:00:00Z",
			"owner": fiber.Map{
				"id": 1,
				"username": "test_user",
				"avatar": "https://api.dicebear.com/7.x/avataaars/svg?seed=test_user",
			},
			"memberCount": 2,
			"onlineCount": 1,
		},
		{
			"id": 2,
			"name": "数学研讨室",
			"description": "一起解决数学难题",
			"shareLink": "math_room_456",
			"maxMembers": 10,
			"isPrivate": false,
			"theme": "math",
			"backgroundImage": "https://images.unsplash.com/photo-1635070041078-e363dbe005cb",
			"createdAt": "2025-05-02T10:00:00Z",
			"expiresAt": "2025-05-09T10:00:00Z",
			"owner": fiber.Map{
				"id": 2,
				"username": "study_buddy",
				"avatar": "https://api.dicebear.com/7.x/avataaars/svg?seed=study_buddy",
			},
			"memberCount": 2,
			"onlineCount": 1,
		},
		{
			"id": 3,
			"name": "专注学习间",
			"description": "安静的学习环境",
			"shareLink": "focus_room_789",
			"maxMembers": 5,
			"isPrivate": true,
			"theme": "minimal",
			"backgroundImage": "https://images.unsplash.com/photo-1497032628192-86f99bcd76bc",
			"createdAt": "2025-05-03T10:00:00Z",
			"expiresAt": "2025-05-10T10:00:00Z",
			"owner": fiber.Map{
				"id": 3,
				"username": "focusmaster",
				"avatar": "https://api.dicebear.com/7.x/avataaars/svg?seed=focusmaster",
			},
			"memberCount": 2,
			"onlineCount": 1,
		},
	}

	return c.Status(200).JSON(fiber.Map{
		"data": studyRooms,
		"total": len(studyRooms),
	})
}

// GetStudyRoom 获取自习室详情处理器
func GetStudyRoom(c *fiber.Ctx) error {
	roomID := c.Params("id")
	
	// 获取当前用户ID (仅用于日志或权限检查)
	userID, _ := utils.GetUserIDFromToken(c)
	_ = userID // 避免未使用变量的警告
	
	// 根据ID返回不同的自习室信息
	var room fiber.Map
	
	if roomID == "1" {
		room = fiber.Map{
			"id": 1,
			"name": "编程自习室",
			"description": "专注于编程学习的自习室",
			"shareLink": "code_room_123",
			"maxMembers": 20,
			"isPrivate": false,
			"theme": "coding",
			"backgroundImage": "https://images.unsplash.com/photo-1605379399642-870262d3d051",
			"createdAt": "2025-05-01T10:00:00Z",
			"expiresAt": "2025-05-08T10:00:00Z",
			"owner": fiber.Map{
				"id": 1,
				"username": "test_user",
				"avatar": "https://api.dicebear.com/7.x/avataaars/svg?seed=test_user",
			},
			"memberCount": 2,
			"onlineCount": 1,
			"members": []fiber.Map{
				{
					"id": 1,
					"userID": 1,
					"username": "test_user",
					"avatar": "https://api.dicebear.com/7.x/avataaars/svg?seed=test_user",
					"isAnonymous": false,
					"role": "owner",
					"status": "online",
					"joinedAt": "2025-05-01T10:00:00Z",
				},
				{
					"id": 2,
					"userID": 2,
					"username": "study_buddy",
					"avatar": "https://api.dicebear.com/7.x/avataaars/svg?seed=study_buddy",
					"isAnonymous": false,
					"role": "member",
					"status": "online",
					"joinedAt": "2025-05-01T10:30:00Z",
				},
			},
		}
	} else if roomID == "2" {
		room = fiber.Map{
			"id": 2,
			"name": "数学研讨室",
			"description": "一起解决数学难题",
			"shareLink": "math_room_456",
			"maxMembers": 10,
			"isPrivate": false,
			"theme": "math",
			"backgroundImage": "https://images.unsplash.com/photo-1635070041078-e363dbe005cb",
			"createdAt": "2025-05-02T10:00:00Z",
			"expiresAt": "2025-05-09T10:00:00Z",
			"owner": fiber.Map{
				"id": 2,
				"username": "study_buddy",
				"avatar": "https://api.dicebear.com/7.x/avataaars/svg?seed=study_buddy",
			},
			"memberCount": 2,
			"onlineCount": 1,
			"members": []fiber.Map{
				{
					"id": 3,
					"userID": 2,
					"username": "study_buddy",
					"avatar": "https://api.dicebear.com/7.x/avataaars/svg?seed=study_buddy",
					"isAnonymous": false,
					"role": "owner",
					"status": "online",
					"joinedAt": "2025-05-02T10:00:00Z",
				},
				{
					"id": 4,
					"userID": 3,
					"username": "focusmaster",
					"avatar": "https://api.dicebear.com/7.x/avataaars/svg?seed=focusmaster",
					"isAnonymous": false,
					"role": "member",
					"status": "away",
					"joinedAt": "2025-05-02T10:15:00Z",
				},
			},
		}
	} else if roomID == "3" {
		room = fiber.Map{
			"id": 3,
			"name": "专注学习间",
			"description": "安静的学习环境",
			"shareLink": "focus_room_789",
			"maxMembers": 5,
			"isPrivate": true,
			"theme": "minimal",
			"backgroundImage": "https://images.unsplash.com/photo-1497032628192-86f99bcd76bc",
			"createdAt": "2025-05-03T10:00:00Z",
			"expiresAt": "2025-05-10T10:00:00Z",
			"owner": fiber.Map{
				"id": 3,
				"username": "focusmaster",
				"avatar": "https://api.dicebear.com/7.x/avataaars/svg?seed=focusmaster",
			},
			"memberCount": 2,
			"onlineCount": 1,
			"members": []fiber.Map{
				{
					"id": 5,
					"userID": 3,
					"username": "focusmaster",
					"avatar": "https://api.dicebear.com/7.x/avataaars/svg?seed=focusmaster",
					"isAnonymous": false,
					"role": "owner",
					"status": "online",
					"joinedAt": "2025-05-03T10:00:00Z",
				},
				{
					"id": 6,
					"userID": 1,
					"username": "Anonymous",
					"avatar": "",
					"isAnonymous": true,
					"role": "member",
					"status": "online",
					"joinedAt": "2025-05-03T11:00:00Z",
				},
			},
		}
	} else {
		// 如果ID不存在，返回404
		return c.Status(404).JSON(fiber.Map{
			"code": 404,
			"message": "自习室不存在",
		})
	}
	
	return c.Status(200).JSON(room)
}

// CreateStudyRoomRequest 创建自习室请求结构
type CreateStudyRoomRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	MaxMembers      int    `json:"maxMembers"`
	IsPrivate       bool   `json:"isPrivate"`
	Theme           string `json:"theme"`
	BackgroundImage string `json:"backgroundImage"`
	ExpiresIn       int    `json:"expiresIn"` // 过期时间，单位为小时
}

// CreateStudyRoom 创建自习室处理器
func CreateStudyRoom(c *fiber.Ctx) error {
	// 从Token获取用户ID
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"code": 401,
			"message": "未授权",
		})
	}

	// 解析请求体
	req := new(CreateStudyRoomRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code": 400,
			"message": "请求参数无效",
		})
	}

	// 验证请求参数
	if req.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"code": 400,
			"message": "自习室名称不能为空",
		})
	}

	if req.MaxMembers <= 0 {
		req.MaxMembers = 10 // 设置默认值
	}

	if req.ExpiresIn <= 0 {
		req.ExpiresIn = 24 // 默认24小时
	}

	// 创建示例自习室
	newRoom := fiber.Map{
		"id":              4, // 示例ID
		"name":            req.Name,
		"description":     req.Description,
		"shareLink":       "generated_link_" + req.Name,
		"maxMembers":      req.MaxMembers,
		"isPrivate":       req.IsPrivate,
		"theme":           req.Theme,
		"backgroundImage": req.BackgroundImage,
		"createdAt":       "2025-05-04T10:00:00Z", // 示例时间
		"expiresAt":       "2025-05-05T10:00:00Z", // 示例时间
		"owner": fiber.Map{
			"id":       userID,
			"username": "test_user",
			"avatar":   "https://api.dicebear.com/7.x/avataaars/svg?seed=test_user",
		},
		"memberCount": 1,
		"onlineCount": 1,
	}

	return c.Status(201).JSON(newRoom)
}

// JoinStudyRoomRequest 加入自习室请求结构
type JoinStudyRoomRequest struct {
	ShareLink    string `json:"shareLink"`
	IsAnonymous  bool   `json:"isAnonymous"`
}

// JoinStudyRoom 加入自习室处理器
func JoinStudyRoom(c *fiber.Ctx) error {
	// 获取用户ID (如果用户已登录)
	userID, _ := utils.GetUserIDFromToken(c)
	
	// 解析请求体
	req := new(JoinStudyRoomRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code": 400,
			"message": "请求参数无效",
		})
	}

	// 验证请求参数
	if req.ShareLink == "" {
		return c.Status(400).JSON(fiber.Map{
			"code": 400,
			"message": "分享链接不能为空",
		})
	}

	// 查找对应的自习室
	var roomID int
	switch req.ShareLink {
	case "code_room_123":
		roomID = 1
	case "math_room_456":
		roomID = 2
	case "focus_room_789":
		roomID = 3
	default:
		return c.Status(404).JSON(fiber.Map{
			"code": 404,
			"message": "自习室不存在",
		})
	}

	// 返回加入结果
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "成功加入自习室",
		"roomId":  roomID,
		"member": fiber.Map{
			"id":          99, // 示例ID
			"userID":      userID,
			"username":    getAnonymousUsername(req.IsAnonymous),
			"avatar":      getAnonymousAvatar(req.IsAnonymous),
			"isAnonymous": req.IsAnonymous,
			"role":        "member",
			"status":      "online",
			"joinedAt":    "2025-05-04T12:00:00Z", // 示例时间
		},
	})
}

// 辅助函数，根据是否匿名返回用户名
func getAnonymousUsername(isAnonymous bool) string {
	if isAnonymous {
		return "Anonymous"
	}
	return "current_user"
}

// 辅助函数，根据是否匿名返回头像
func getAnonymousAvatar(isAnonymous bool) string {
	if isAnonymous {
		return ""
	}
	return "https://api.dicebear.com/7.x/avataaars/svg?seed=current_user"
} 