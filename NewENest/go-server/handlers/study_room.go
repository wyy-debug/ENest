package handlers

import (
	"NewENest/go-server/database"
	"NewENest/go-server/utils"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"fmt"
	"time"
	"github.com/jmoiron/sqlx"
)

// 全局数据库连接实例和仓库实例
var DB *sqlx.DB
var studyRoomRepo *database.StudyRoomRepository

// SetDB 设置数据库连接
func SetDB(db *sqlx.DB) {
	DB = db
	studyRoomRepo = database.NewStudyRoomRepository(db)
}

// GetStudyRooms 获取自习室列表处理器
func GetStudyRooms(c *fiber.Ctx) error {
	// 获取分页参数
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	
	pageSize, err := strconv.Atoi(c.Query("pageSize", "10"))
	if err != nil || pageSize < 1 || pageSize > 50 {
		pageSize = 10
	}
	
	// 从数据库获取自习室列表
	rooms, total, err := studyRoomRepo.GetAllStudyRooms(page, pageSize)
	if err != nil {
		log.Printf("获取自习室列表失败: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"code": 500,
			"message": "获取自习室列表失败，服务器错误",
		})
	}
	
	// 构建响应数据
	result := make([]fiber.Map, 0, len(rooms))
	for _, room := range rooms {
		// 获取房主信息
		owner, err := studyRoomRepo.GetUserByID(room.OwnerID)
		if err != nil {
			// 如果获取房主信息失败，使用默认值
			owner = &database.User{
				ID:       room.OwnerID,
				Username: "用户" + strconv.Itoa(room.OwnerID),
				Avatar:   "https://api.dicebear.com/7.x/avataaars/svg?seed=user" + strconv.Itoa(room.OwnerID),
			}
		}
		
		// 获取成员数量
		memberCount, err := studyRoomRepo.CountStudyRoomMembers(room.ID)
		if err != nil {
			memberCount = 0
		}
		
		// 格式化时间
		createdAt := utils.FormatTime(room.CreatedAt)
		expiresAt := utils.FormatTime(room.ExpiresAt)
		
		// 构建自习室数据
		roomData := fiber.Map{
			"id":               room.ID,
			"name":             room.Name,
			"description":      room.Description,
			"share_link":       room.ShareLink,
			"max_members":      room.MaxMembers,
			"is_private":       room.IsPrivate,
			"theme":            room.Theme,
			"background_image": room.BackgroundImage,
			"created_at":       createdAt,
			"expires_at":       expiresAt,
			"owner": fiber.Map{
				"id":       owner.ID,
				"username": owner.Username,
				"avatar":   owner.Avatar,
			},
			"member_count": memberCount,
		}
		
		result = append(result, roomData)
	}
	
	return c.Status(200).JSON(fiber.Map{
		"code": 200,
		"message": "获取自习室列表成功",
		"data": fiber.Map{
			"rooms": result,
			"total": total,
		},
	})
}

// GetStudyRoom 获取自习室详情处理器
func GetStudyRoom(c *fiber.Ctx) error {
	// 获取自习室ID
	roomIDStr := c.Params("id")
	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code": 400,
			"message": "无效的自习室ID",
		})
	}
	
	// 获取当前用户ID (仅用于日志或权限检查)
	userID, _ := utils.GetUserIDFromToken(c)
	log.Printf("GetStudyRoom - 用户 %d 正在查看自习室 %d", userID, roomID)
	
	// 从数据库获取自习室信息
	room, err := studyRoomRepo.GetStudyRoomByID(roomID)
	if err != nil {
		log.Printf("获取自习室失败: %v", err)
		return c.Status(404).JSON(fiber.Map{
			"code": 404,
			"message": "自习室不存在",
		})
	}
	
	// 获取房主信息
	owner, err := studyRoomRepo.GetUserByID(room.OwnerID)
	if err != nil {
		// 如果获取房主信息失败，使用默认值
		owner = &database.User{
			ID:       room.OwnerID,
			Username: "用户" + strconv.Itoa(room.OwnerID),
			Avatar:   "https://api.dicebear.com/7.x/avataaars/svg?seed=user" + strconv.Itoa(room.OwnerID),
		}
	}
	
	// 获取自习室成员列表
	membersData, err := studyRoomRepo.GetStudyRoomMembers(roomID)
	if err != nil {
		log.Printf("获取成员列表失败: %v", err)
		membersData = []map[string]interface{}{}
	}
	
	// 格式化成员数据
	members := make([]fiber.Map, 0, len(membersData))
	for _, m := range membersData {
		// 处理时间格式
		joinedAtTime, ok := m["joined_at"].(time.Time)
		joinedAt := ""
		if ok {
			joinedAt = utils.FormatTime(joinedAtTime)
		} else {
			joinedAt = utils.FormatTime(time.Now())
		}
		
		// 构建成员信息
		member := fiber.Map{
			"user_id":      m["user_id"],
			"username":     m["username"],
			"avatar":       m["avatar"],
			"is_anonymous": m["is_anonymous"],
			"role":         m["role"],
			"status":       m["status"],
			"joined_at":    joinedAt,
		}
		members = append(members, member)
	}
	
	// 格式化时间
	createdAt := utils.FormatTime(room.CreatedAt)
	expiresAt := utils.FormatTime(room.ExpiresAt)
	
	// 构建自习室详情响应
	roomDetail := fiber.Map{
		"id":               room.ID,
		"name":             room.Name,
		"description":      room.Description,
		"share_link":       room.ShareLink,
		"max_members":      room.MaxMembers,
		"is_private":       room.IsPrivate,
		"theme":            room.Theme,
		"background_image": room.BackgroundImage,
		"created_at":       createdAt,
		"expires_at":       expiresAt,
		"owner": fiber.Map{
			"id":       owner.ID,
			"username": owner.Username,
			"avatar":   owner.Avatar,
		},
		"member_count": len(members),
		"members":      members,
	}
	
	return c.Status(200).JSON(fiber.Map{
		"code": 200,
		"message": "获取自习室详情成功",
		"data": roomDetail,
	})
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
		log.Printf("CreateStudyRoom - 获取用户ID失败: %v", err)
		return c.Status(401).JSON(fiber.Map{
			"code": 401,
			"message": "未授权，请先登录",
		})
	}
	
	log.Printf("CreateStudyRoom - 成功获取用户ID: %d", userID)

	// 解析请求体
	req := new(CreateStudyRoomRequest)
	if err := c.BodyParser(req); err != nil {
		log.Printf("CreateStudyRoom - 解析请求体失败: %v", err)
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

	// 获取当前时间和过期时间
	currentTime := time.Now()
	expiresAt := currentTime.Add(time.Duration(req.ExpiresIn) * time.Hour)

	// 生成唯一的分享链接
	shareLink := fmt.Sprintf("room_%d_%s", userID, utils.GenerateRandomString(8))
	
	// 创建自习室
	room := &database.StudyRoom{
		OwnerID:         userID,
		Name:            req.Name,
		Description:     req.Description,
		ShareLink:       shareLink,
		MaxMembers:      req.MaxMembers,
		IsPrivate:       req.IsPrivate,
		Theme:           req.Theme,
		BackgroundImage: req.BackgroundImage,
		CreatedAt:       currentTime,
		ExpiresAt:       expiresAt,
	}
	
	// 将自习室保存到数据库
	roomID, err := studyRoomRepo.CreateStudyRoom(room)
	if err != nil {
		log.Printf("创建自习室失败: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"code": 500,
			"message": "创建自习室失败，服务器错误",
		})
	}
	
	// 获取创建者信息
	owner, err := studyRoomRepo.GetUserByID(userID)
	if err != nil {
		owner = &database.User{
			ID:       userID,
			Username: "用户" + strconv.Itoa(userID),
			Avatar:   "https://api.dicebear.com/7.x/avataaars/svg?seed=user" + strconv.Itoa(userID),
		}
	}
	
	// 添加创建者为成员
	member := &database.StudyRoomMember{
		RoomID:      roomID,
		UserID:      userID,
		IsAnonymous: false,
		Role:        "owner",
		Status:      "active",
		JoinedAt:    currentTime,
	}

	err = studyRoomRepo.AddMemberToStudyRoom(member)
	if err != nil {
		log.Printf("添加创建者为成员失败: %v", err)
		// 这里不返回错误，因为自习室已经创建成功
	}
	
	// 获取格式化的时间字符串
	createdAtStr := utils.FormatTime(currentTime)
	expiresAtStr := utils.FormatTime(expiresAt)
	
	// 构建要返回的自习室信息
	ownerInfo := fiber.Map{
		"id":       owner.ID,
		"username": owner.Username,
		"avatar":   owner.Avatar,
	}
	
	// 创建成员列表
	members := []fiber.Map{
		{
			"user_id":      userID,
			"username":     owner.Username,
			"avatar":       owner.Avatar,
			"is_anonymous": false,
			"role":         "owner",
			"status":       "active",
			"joined_at":    createdAtStr,
		},
	}
	
	// 构建响应
	newRoom := fiber.Map{
		"id":                roomID,
		"name":              req.Name,
		"description":       req.Description,
		"share_link":        shareLink,
		"max_members":       req.MaxMembers,
		"is_private":        req.IsPrivate,
		"theme":             req.Theme,
		"background_image":  req.BackgroundImage,
		"created_at":        createdAtStr,
		"expires_at":        expiresAtStr,
		"owner":             ownerInfo,
		"member_count":      1,
		"members":           members,
	}

	// 日志记录
	log.Printf("创建自习室成功: ID=%d, 名称=%s", roomID, req.Name)

	// 返回响应
	return c.Status(201).JSON(fiber.Map{
		"code": 201,
		"message": "自习室创建成功",
		"data": newRoom,
	})
}

// JoinStudyRoomRequest 加入自习室请求结构
type JoinStudyRoomRequest struct {
	ShareLink    string `json:"shareLink"`
	IsAnonymous  bool   `json:"isAnonymous"`
}

// JoinStudyRoom 加入自习室处理器
func JoinStudyRoom(c *fiber.Ctx) error {
	// 获取用户ID (如果用户已登录)
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		log.Printf("JoinStudyRoom - 获取用户ID失败: %v", err)
		return c.Status(401).JSON(fiber.Map{
			"code": 401,
			"message": "未授权，请先登录",
		})
	}
	
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

	// 根据分享链接获取自习室
	room, err := studyRoomRepo.GetStudyRoomByShareLink(req.ShareLink)
	if err != nil {
		log.Printf("获取自习室失败: %v", err)
		return c.Status(404).JSON(fiber.Map{
			"code": 404,
			"message": "自习室不存在或链接已失效",
		})
	}
	
	// 声明变量用于保存用户名和头像
	var username string
	var avatar string
	
	// 获取当前用户信息
	if req.IsAnonymous {
		// 如果是匿名加入，直接使用匿名值
		username = getAnonymousUsername(true)
		avatar = getAnonymousAvatar(true)
		log.Printf("用户 %d 以匿名方式加入自习室", userID)
	} else {
		// 否则获取用户真实信息
		user, err := studyRoomRepo.GetUserByID(userID)
		if err != nil {
			// 如果获取用户信息失败，使用默认值
			log.Printf("获取用户信息失败: %v，使用默认值", err)
			username = getAnonymousUsername(false)
			avatar = getAnonymousAvatar(false)
		} else {
			username = user.Username
			avatar = user.Avatar
		}
	}
	
	// 获取成员数量
	memberCount, err := studyRoomRepo.CountStudyRoomMembers(room.ID)
	if err == nil && memberCount >= room.MaxMembers {
		return c.Status(403).JSON(fiber.Map{
			"code": 403,
			"message": "自习室已满，无法加入",
		})
	}
	
	// 添加用户到自习室成员
	joinedAt := time.Now()
	member := &database.StudyRoomMember{
		RoomID:      room.ID,
		UserID:      userID,
		IsAnonymous: req.IsAnonymous,
		Role:        "member",
		Status:      "active",
		JoinedAt:    joinedAt,
	}
	
	err = studyRoomRepo.AddMemberToStudyRoom(member)
	if err != nil {
		log.Printf("添加成员失败: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"code": 500,
			"message": "加入自习室失败，服务器错误",
		})
	}
	
	// 将用户名和头像信息添加到日志中
	log.Printf("用户 %d 加入自习室 %d，用户名: %s, 头像: %s", userID, room.ID, username, avatar)
	
	// 更新成员数量
	memberCount++
	
	// 获取房主信息
	owner, err := studyRoomRepo.GetUserByID(room.OwnerID)
	if err != nil {
		// 如果获取房主信息失败，使用默认值
		owner = &database.User{
			ID:       room.OwnerID,
			Username: "用户" + strconv.Itoa(room.OwnerID),
			Avatar:   "https://api.dicebear.com/7.x/avataaars/svg?seed=user" + strconv.Itoa(room.OwnerID),
		}
	}
	
	// 格式化时间
	createdAt := utils.FormatTime(room.CreatedAt)
	expiresAt := utils.FormatTime(room.ExpiresAt)
	
	// 获取自习室所有成员
	membersData, err := studyRoomRepo.GetStudyRoomMembers(room.ID)
	if err != nil {
		log.Printf("获取成员列表失败: %v", err)
		membersData = []map[string]interface{}{}
	}
	
	// 格式化成员数据
	members := make([]fiber.Map, 0, len(membersData))
	for _, m := range membersData {
		// 处理时间格式
		memberJoinedAtTime, ok := m["joined_at"].(time.Time)
		memberJoinedAtStr := ""
		if ok {
			memberJoinedAtStr = utils.FormatTime(memberJoinedAtTime)
		} else {
			memberJoinedAtStr = utils.FormatTime(time.Now())
		}
		
		// 构建成员信息
		memberInfo := fiber.Map{
			"user_id":      m["user_id"],
			"username":     m["username"],
			"avatar":       m["avatar"],
			"is_anonymous": m["is_anonymous"],
			"role":         m["role"],
			"status":       m["status"],
			"joined_at":    memberJoinedAtStr,
		}
		members = append(members, memberInfo)
	}
	
	// 构建自习室响应
	roomData := fiber.Map{
		"id":               room.ID,
		"name":             room.Name,
		"description":      room.Description,
		"share_link":       room.ShareLink,
		"max_members":      room.MaxMembers,
		"is_private":       room.IsPrivate,
		"theme":            room.Theme,
		"background_image": room.BackgroundImage,
		"created_at":       createdAt,
		"expires_at":       expiresAt,
		"owner": fiber.Map{
			"id":       owner.ID,
			"username": owner.Username,
			"avatar":   owner.Avatar,
		},
		"member_count": memberCount,
		"members":      members,
	}

	// 返回加入结果
	return c.Status(200).JSON(fiber.Map{
		"code": 200,
		"message": "成功加入自习室",
		"data": roomData,
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

// UpdateStudyRoomRequest 更新自习室请求结构
type UpdateStudyRoomRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	MaxMembers      int    `json:"maxMembers"`
	IsPrivate       bool   `json:"isPrivate"`
	Theme           string `json:"theme"`
	BackgroundImage string `json:"backgroundImage"`
}

// UpdateStudyRoom 更新自习室处理器
func UpdateStudyRoom(c *fiber.Ctx) error {
	// 获取用户ID
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		log.Printf("UpdateStudyRoom - 获取用户ID失败: %v", err)
		return c.Status(401).JSON(fiber.Map{
			"code": 401,
			"message": "未授权，请先登录",
		})
	}
	
	// 解析自习室ID
	roomID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code": 400,
			"message": "无效的自习室ID",
		})
	}

	// 解析请求体
	req := new(UpdateStudyRoomRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code": 400,
			"message": "请求参数无效",
		})
	}

	// 查询自习室
	room, err := studyRoomRepo.GetStudyRoomByID(roomID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"code": 404,
			"message": "自习室不存在",
		})
	}
	
	// 检查权限：只有房主可以更新自习室信息
	if room.OwnerID != userID {
		return c.Status(403).JSON(fiber.Map{
			"code": 403,
			"message": "只有房主可以更新自习室信息",
		})
	}
	
	// 更新自习室信息
	room.Name = req.Name
	room.Description = req.Description
	room.MaxMembers = req.MaxMembers
	room.IsPrivate = req.IsPrivate
	room.Theme = req.Theme
	room.BackgroundImage = req.BackgroundImage
	
	// 保存到数据库
	err = studyRoomRepo.UpdateStudyRoom(room)
	if err != nil {
		log.Printf("更新自习室失败: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"code": 500,
			"message": "更新自习室失败，服务器错误",
		})
	}
	
	// 获取房主信息
	owner, err := studyRoomRepo.GetUserByID(room.OwnerID)
	if err != nil {
		// 如果获取房主信息失败，使用默认值
		owner = &database.User{
			ID:       room.OwnerID,
			Username: "用户" + strconv.Itoa(room.OwnerID),
			Avatar:   "https://api.dicebear.com/7.x/avataaars/svg?seed=user" + strconv.Itoa(room.OwnerID),
		}
	}
	
	// 获取自习室成员
	membersData, err := studyRoomRepo.GetStudyRoomMembers(roomID)
	if err != nil {
		log.Printf("获取成员列表失败: %v", err)
		membersData = []map[string]interface{}{}
	}
	
	// 格式化成员数据
	members := make([]fiber.Map, 0, len(membersData))
	for _, m := range membersData {
		joinedAtTime, ok := m["joined_at"].(time.Time)
		joinedAt := ""
		if ok {
			joinedAt = utils.FormatTime(joinedAtTime)
		} else {
			joinedAt = utils.FormatTime(time.Now())
		}
		
		member := fiber.Map{
			"user_id":      m["user_id"],
			"username":     m["username"],
			"avatar":       m["avatar"],
			"is_anonymous": m["is_anonymous"],
			"role":         m["role"],
			"status":       m["status"],
			"joined_at":    joinedAt,
		}
		members = append(members, member)
	}
	
	// 格式化时间
	createdAt := utils.FormatTime(room.CreatedAt)
	expiresAt := utils.FormatTime(room.ExpiresAt)
	
	// 构建更新后的自习室数据
	updatedRoom := fiber.Map{
		"id":               room.ID,
		"name":             room.Name,
		"description":      room.Description,
		"share_link":       room.ShareLink,
		"max_members":      room.MaxMembers,
		"is_private":       room.IsPrivate,
		"theme":            room.Theme,
		"background_image": room.BackgroundImage,
		"created_at":       createdAt,
		"expires_at":       expiresAt,
		"owner": fiber.Map{
			"id":       owner.ID,
			"username": owner.Username,
			"avatar":   owner.Avatar,
		},
		"member_count": len(members),
		"members":      members,
	}

	// 返回更新结果
	return c.Status(200).JSON(fiber.Map{
		"code": 200,
		"message": "自习室更新成功",
		"data": updatedRoom,
	})
}

// LeaveStudyRoomRequest 离开自习室请求结构
type LeaveStudyRoomRequest struct {
	RoomID int `json:"room_id"`
}

// LeaveStudyRoom 离开自习室处理器
func LeaveStudyRoom(c *fiber.Ctx) error {
	// 获取用户ID
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		log.Printf("LeaveStudyRoom - 获取用户ID失败: %v", err)
		return c.Status(401).JSON(fiber.Map{
			"code": 401,
			"message": "未授权，请先登录",
		})
	}
	
	// 解析请求体
	req := new(LeaveStudyRoomRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code": 400,
			"message": "请求参数无效",
		})
	}

	// 验证请求参数
	if req.RoomID <= 0 {
		return c.Status(400).JSON(fiber.Map{
			"code": 400,
			"message": "自习室ID无效",
		})
	}

	// 检查自习室是否存在
	room, err := studyRoomRepo.GetStudyRoomByID(req.RoomID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"code": 404,
			"message": "自习室不存在",
		})
	}
	
	// 检查用户是否为房主
	if room.OwnerID == userID {
		// TODO: 如果是房主离开，可以选择销毁自习室或转移所有权，此处简化处理
		log.Printf("房主 %d 离开自习室 %d", userID, req.RoomID)
	}

	// 将用户从自习室成员中移除
	err = studyRoomRepo.RemoveMemberFromStudyRoom(req.RoomID, userID)
	if err != nil {
		log.Printf("移除成员失败: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"code": 500,
			"message": "离开自习室失败，服务器错误",
		})
	}

	// 返回离开结果
	return c.Status(200).JSON(fiber.Map{
		"code": 200,
		"message": "成功离开自习室",
		"data": fiber.Map{
			"room_id": req.RoomID,
			"user_id": userID,
		},
	})
}

// GetStudyRoomByID 根据ID获取自习室详情
func GetStudyRoomByID(roomID int) fiber.Map {
	// 模拟数据库操作，根据ID返回不同的自习室
	var room fiber.Map
	
	if roomID == 1 {
		room = fiber.Map{
			"id": 1,
			"name": "编程自习室",
			"description": "专注于编程学习的自习室",
			"share_link": "code_room_123",
			"max_members": 20,
			"is_private": false,
			"theme": "coding",
			"background_image": "https://images.unsplash.com/photo-1605379399642-870262d3d051",
			"created_at": "2025-05-01T10:00:00Z",
			"expires_at": "2025-05-08T10:00:00Z",
			"owner": fiber.Map{
				"id": 1,
				"username": "test_user",
				"avatar": "https://api.dicebear.com/7.x/avataaars/svg?seed=test_user",
			},
			"member_count": 2,
			"members": []fiber.Map{
				{
					"user_id": 1,
					"username": "test_user",
					"avatar": "https://api.dicebear.com/7.x/avataaars/svg?seed=test_user",
					"is_anonymous": false,
					"role": "owner",
					"status": "online",
					"joined_at": "2025-05-01T10:00:00Z",
				},
				{
					"user_id": 2,
					"username": "study_buddy",
					"avatar": "https://api.dicebear.com/7.x/avataaars/svg?seed=study_buddy",
					"is_anonymous": false,
					"role": "member",
					"status": "online",
					"joined_at": "2025-05-01T10:30:00Z",
				},
			},
		}
	} else if roomID == 2 {
		room = fiber.Map{
			"id": 2,
			"name": "数学研讨室",
			"description": "一起解决数学难题",
			"share_link": "math_room_456",
			"max_members": 10,
			"is_private": false,
			"theme": "math",
			"background_image": "https://images.unsplash.com/photo-1635070041078-e363dbe005cb",
			"created_at": "2025-05-02T10:00:00Z",
			"expires_at": "2025-05-09T10:00:00Z",
			"owner": fiber.Map{
				"id": 2,
				"username": "study_buddy",
				"avatar": "https://api.dicebear.com/7.x/avataaars/svg?seed=study_buddy",
			},
			"member_count": 2,
			"members": []fiber.Map{
				{
					"user_id": 2,
					"username": "study_buddy",
					"avatar": "https://api.dicebear.com/7.x/avataaars/svg?seed=study_buddy",
					"is_anonymous": false,
					"role": "owner",
					"status": "online",
					"joined_at": "2025-05-02T10:00:00Z",
				},
				{
					"user_id": 3,
					"username": "focusmaster",
					"avatar": "https://api.dicebear.com/7.x/avataaars/svg?seed=focusmaster",
					"is_anonymous": false,
					"role": "member",
					"status": "away",
					"joined_at": "2025-05-02T10:15:00Z",
				},
			},
		}
	} else if roomID == 3 {
		room = fiber.Map{
			"id": 3,
			"name": "专注学习间",
			"description": "安静的学习环境",
			"share_link": "focus_room_789",
			"max_members": 5,
			"is_private": true,
			"theme": "minimal",
			"background_image": "https://images.unsplash.com/photo-1497032628192-86f99bcd76bc",
			"created_at": "2025-05-03T10:00:00Z",
			"expires_at": "2025-05-10T10:00:00Z",
			"owner": fiber.Map{
				"id": 3,
				"username": "focusmaster",
				"avatar": "https://api.dicebear.com/7.x/avataaars/svg?seed=focusmaster",
			},
			"member_count": 2,
			"members": []fiber.Map{
				{
					"user_id": 3,
					"username": "focusmaster",
					"avatar": "https://api.dicebear.com/7.x/avataaars/svg?seed=focusmaster",
					"is_anonymous": false,
					"role": "owner",
					"status": "online",
					"joined_at": "2025-05-03T10:00:00Z",
				},
				{
					"user_id": 1,
					"username": "test_user",
					"avatar": "https://api.dicebear.com/7.x/avataaars/svg?seed=test_user",
					"is_anonymous": false,
					"role": "member",
					"status": "offline",
					"joined_at": "2025-05-03T10:05:00Z",
				},
			},
		}
	} else {
		return nil
	}
	
	return room
} 