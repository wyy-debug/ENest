package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"NewENest/go-server/middleware"
	"NewENest/go-server/models"
)

// CreateTask 创建新任务
func CreateTask(c *fiber.Ctx) error {
	// 从请求上下文获取用户ID
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return Unauthorized(c, "未授权")
	}

	// 解析请求体
	var input models.TaskCreate
	if err := c.BodyParser(&input); err != nil {
		return BadRequest(c, "请求格式无效", err)
	}

	// 创建任务（示例，实际应当保存到数据库）
	task := &models.Task{
		ID:          1, // 实际应由数据库生成
		UserID:      userID,
		Title:       input.Title,
		Description: input.Description,
		Deadline:    input.Deadline,
		Priority:    input.Priority,
		Status:      "pending",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return Created(c, task, "任务创建成功")
}

// GetTasks 获取用户所有任务
func GetTasks(c *fiber.Ctx) error {
	// 从请求上下文获取用户ID
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return Unauthorized(c, "未授权")
	}

	// 获取任务列表（示例，实际应当从数据库获取）
	tasks := []models.Task{
		{
			ID:          1,
			UserID:      userID,
			Title:       "完成项目报告",
			Description: "准备季度项目进展报告",
			Deadline:    time.Now().Add(72 * time.Hour),
			Priority:    "高",
			Status:      "进行中",
			CreatedAt:   time.Now().Add(-24 * time.Hour),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			UserID:      userID,
			Title:       "学习Go语言",
			Description: "完成Go语言基础课程",
			Deadline:    time.Now().Add(7 * 24 * time.Hour),
			Priority:    "中",
			Status:      "未开始",
			CreatedAt:   time.Now().Add(-48 * time.Hour),
			UpdatedAt:   time.Now().Add(-48 * time.Hour),
		},
	}

	return Success(c, tasks, "获取任务列表成功")
}

// GetTask 获取指定任务详情
func GetTask(c *fiber.Ctx) error {
	// 从请求上下文获取用户ID
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return Unauthorized(c, "未授权")
	}

	// 获取任务ID
	taskID, err := c.ParamsInt("id")
	if err != nil {
		return BadRequest(c, "无效的任务ID", err)
	}

	// 获取任务详情（示例，实际应当从数据库获取）
	task := models.Task{
		ID:          taskID,
		UserID:      userID,
		Title:       "完成项目报告",
		Description: "准备季度项目进展报告，包括完成的功能和遇到的挑战",
		Deadline:    time.Now().Add(72 * time.Hour),
		Priority:    "高",
		Status:      "进行中",
		CreatedAt:   time.Now().Add(-24 * time.Hour),
		UpdatedAt:   time.Now(),
	}

	// 检查任务是否属于该用户
	if task.UserID != userID {
		return Forbidden(c, "您无权访问此任务")
	}

	return Success(c, task, "获取任务详情成功")
}

// UpdateTask 更新任务
func UpdateTask(c *fiber.Ctx) error {
	// 从请求上下文获取用户ID
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return Unauthorized(c, "未授权")
	}

	// 获取任务ID
	taskID, err := c.ParamsInt("id")
	if err != nil {
		return BadRequest(c, "无效的任务ID", err)
	}

	// 解析请求体
	var input models.TaskUpdate
	if err := c.BodyParser(&input); err != nil {
		return BadRequest(c, "请求格式无效", err)
	}

	// 获取任务（示例，实际应当从数据库获取）
	task := models.Task{
		ID:          taskID,
		UserID:      userID,
		Title:       "完成项目报告",
		Description: "准备季度项目进展报告",
		Deadline:    time.Now().Add(72 * time.Hour),
		Priority:    "高",
		Status:      "进行中",
		CreatedAt:   time.Now().Add(-24 * time.Hour),
		UpdatedAt:   time.Now(),
	}

	// 检查任务是否属于该用户
	if task.UserID != userID {
		return Forbidden(c, "您无权修改此任务")
	}

	// 更新任务
	if input.Title != "" {
		task.Title = input.Title
	}
	if input.Description != "" {
		task.Description = input.Description
	}
	if !input.Deadline.IsZero() {
		task.Deadline = input.Deadline
	}
	if input.Priority != "" {
		task.Priority = input.Priority
	}
	if input.Status != "" {
		task.Status = input.Status
	}
	task.UpdatedAt = time.Now()

	// 此处应当将更新后的任务信息保存到数据库

	return Success(c, task, "任务更新成功")
}

// DeleteTask 删除任务
func DeleteTask(c *fiber.Ctx) error {
	// 从请求上下文获取用户ID
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return Unauthorized(c, "未授权")
	}

	// 获取任务ID
	taskID, err := c.ParamsInt("id")
	if err != nil {
		return BadRequest(c, "无效的任务ID", err)
	}

	// 获取任务（示例，实际应当从数据库获取）
	task := models.Task{
		ID:     taskID,
		UserID: userID,
	}

	// 检查任务是否属于该用户
	if task.UserID != userID {
		return Forbidden(c, "您无权删除此任务")
	}

	// 此处应当从数据库中删除任务

	return Success(c, nil, "任务删除成功")
} 