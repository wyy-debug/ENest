package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"NewENest/go-server/middleware"
	"NewENest/go-server/models"
)

// CreateStudyRecord 创建新学习记录
func CreateStudyRecord(c *fiber.Ctx) error {
	// 从请求上下文获取用户ID
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return Unauthorized(c, "未授权")
	}

	// 解析请求体
	var input models.StudyRecordCreate
	if err := c.BodyParser(&input); err != nil {
		return BadRequest(c, "请求格式无效", err)
	}

	// 计算学习时长
	duration := int(input.EndTime.Sub(input.StartTime).Seconds())
	if duration <= 0 {
		return BadRequest(c, "结束时间必须晚于开始时间", nil)
	}

	// 创建学习记录（示例，实际应当保存到数据库）
	record := &models.StudyRecord{
		ID:          1, // 实际应由数据库生成
		UserID:      userID,
		StartTime:   input.StartTime,
		EndTime:     input.EndTime,
		Duration:    duration,
		Category:    input.Category,
		Description: input.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// 此处应当更新用户总学习时间
	// userService.UpdateTotalStudyTime(userID, duration)

	return Created(c, record, "学习记录创建成功")
}

// GetStudyRecords 获取用户所有学习记录
func GetStudyRecords(c *fiber.Ctx) error {
	// 从请求上下文获取用户ID
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return Unauthorized(c, "未授权")
	}

	// 获取分页参数
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("page_size", 10)
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// 获取时间范围参数
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	var startDate, endDate time.Time
	var hasDateFilter bool

	if startDateStr != "" {
		if parsedTime, err := time.Parse("2006-01-02", startDateStr); err == nil {
			startDate = parsedTime
			hasDateFilter = true
		}
	}

	if endDateStr != "" {
		if parsedTime, err := time.Parse("2006-01-02", endDateStr); err == nil {
			// 设置为当天的结束时间
			endDate = parsedTime.Add(24*time.Hour - time.Second)
			hasDateFilter = true
		}
	}

	// 获取学习记录（示例，实际应当从数据库获取）
	records := []models.StudyRecord{
		{
			ID:          1,
			UserID:      userID,
			StartTime:   time.Now().Add(-3 * time.Hour),
			EndTime:     time.Now().Add(-1 * time.Hour),
			Duration:    7200, // 2小时
			Category:    "编程",
			Description: "学习Go语言",
			CreatedAt:   time.Now().Add(-3 * time.Hour),
			UpdatedAt:   time.Now().Add(-1 * time.Hour),
		},
		{
			ID:          2,
			UserID:      userID,
			StartTime:   time.Now().Add(-1 * 24 * time.Hour),
			EndTime:     time.Now().Add(-1*24*time.Hour + 2*time.Hour),
			Duration:    7200, // 2小时
			Category:    "外语",
			Description: "学习英语",
			CreatedAt:   time.Now().Add(-1 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-1*24*time.Hour + 2*time.Hour),
		},
	}

	// 应用时间过滤（在实际数据库查询中应当在查询时过滤）
	if hasDateFilter {
		filteredRecords := []models.StudyRecord{}
		for _, record := range records {
			if (startDate.IsZero() || !record.StartTime.Before(startDate)) &&
				(endDate.IsZero() || !record.StartTime.After(endDate)) {
				filteredRecords = append(filteredRecords, record)
			}
		}
		records = filteredRecords
	}

	// 返回响应
	return Success(c, fiber.Map{
		"records":    records,
		"total":      len(records),
		"page":       page,
		"page_size":  pageSize,
		"total_page": (len(records) + pageSize - 1) / pageSize,
	}, "获取学习记录成功")
}

// GetStudyRecord 获取指定学习记录详情
func GetStudyRecord(c *fiber.Ctx) error {
	// 从请求上下文获取用户ID
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return Unauthorized(c, "未授权")
	}

	// 获取记录ID
	recordID, err := c.ParamsInt("id")
	if err != nil {
		return BadRequest(c, "无效的记录ID", err)
	}

	// 获取学习记录（示例，实际应当从数据库获取）
	record := models.StudyRecord{
		ID:          recordID,
		UserID:      userID,
		StartTime:   time.Now().Add(-3 * time.Hour),
		EndTime:     time.Now().Add(-1 * time.Hour),
		Duration:    7200, // 2小时
		Category:    "编程",
		Description: "学习Go语言基础知识，包括接口、协程和通道",
		CreatedAt:   time.Now().Add(-3 * time.Hour),
		UpdatedAt:   time.Now().Add(-1 * time.Hour),
	}

	// 检查记录是否属于该用户
	if record.UserID != userID {
		return Forbidden(c, "您无权访问此记录")
	}

	return Success(c, record, "获取学习记录详情成功")
}

// UpdateStudyRecord 更新学习记录
func UpdateStudyRecord(c *fiber.Ctx) error {
	// 从请求上下文获取用户ID
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return Unauthorized(c, "未授权")
	}

	// 获取记录ID
	recordID, err := c.ParamsInt("id")
	if err != nil {
		return BadRequest(c, "无效的记录ID", err)
	}

	// 解析请求体
	var input models.StudyRecordUpdate
	if err := c.BodyParser(&input); err != nil {
		return BadRequest(c, "请求格式无效", err)
	}

	// 获取学习记录（示例，实际应当从数据库获取）
	record := models.StudyRecord{
		ID:          recordID,
		UserID:      userID,
		StartTime:   time.Now().Add(-3 * time.Hour),
		EndTime:     time.Now().Add(-1 * time.Hour),
		Duration:    7200, // 2小时
		Category:    "编程",
		Description: "学习Go语言",
		CreatedAt:   time.Now().Add(-3 * time.Hour),
		UpdatedAt:   time.Now().Add(-1 * time.Hour),
	}

	// 检查记录是否属于该用户
	if record.UserID != userID {
		return Forbidden(c, "您无权修改此记录")
	}

	// 记录原始时长
	originalDuration := record.Duration

	// 更新记录
	if !input.StartTime.IsZero() {
		record.StartTime = input.StartTime
	}
	if !input.EndTime.IsZero() {
		record.EndTime = input.EndTime
	}
	if input.Category != "" {
		record.Category = input.Category
	}
	if input.Description != "" {
		record.Description = input.Description
	}

	// 重新计算时长
	record.Duration = int(record.EndTime.Sub(record.StartTime).Seconds())
	if record.Duration <= 0 {
		return BadRequest(c, "结束时间必须晚于开始时间", nil)
	}

	// 更新时间
	record.UpdatedAt = time.Now()

	// 此处应当更新用户总学习时间，增加或减少差值
	// durationDiff := record.Duration - originalDuration
	// userService.UpdateTotalStudyTime(userID, durationDiff)

	return Success(c, record, "学习记录更新成功")
}

// DeleteStudyRecord 删除学习记录
func DeleteStudyRecord(c *fiber.Ctx) error {
	// 从请求上下文获取用户ID
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return Unauthorized(c, "未授权")
	}

	// 获取记录ID
	recordID, err := c.ParamsInt("id")
	if err != nil {
		return BadRequest(c, "无效的记录ID", err)
	}

	// 获取学习记录（示例，实际应当从数据库获取）
	record := models.StudyRecord{
		ID:       recordID,
		UserID:   userID,
		Duration: 7200, // 2小时
	}

	// 检查记录是否属于该用户
	if record.UserID != userID {
		return Forbidden(c, "您无权删除此记录")
	}

	// 此处应当从数据库中删除记录
	// 并更新用户总学习时间，减去记录时长
	// userService.UpdateTotalStudyTime(userID, -record.Duration)

	return Success(c, nil, "学习记录删除成功")
}

// GetStudyStats 获取学习统计数据
func GetStudyStats(c *fiber.Ctx) error {
	// 从请求上下文获取用户ID
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return Unauthorized(c, "未授权")
	}

	// 获取时间范围参数
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	var startDate, endDate time.Time

	if startDateStr != "" {
		if parsedTime, err := time.Parse("2006-01-02", startDateStr); err == nil {
			startDate = parsedTime
		}
	}

	if endDateStr != "" {
		if parsedTime, err := time.Parse("2006-01-02", endDateStr); err == nil {
			// 设置为当天的结束时间
			endDate = parsedTime.Add(24*time.Hour - time.Second)
		}
	}

	// 如果没有指定开始日期，默认为过去30天
	if startDate.IsZero() {
		startDate = time.Now().AddDate(0, 0, -30)
	}

	// 如果没有指定结束日期，默认为当前时间
	if endDate.IsZero() {
		endDate = time.Now()
	}

	// 获取学习统计数据（示例，实际应当从数据库计算）
	stats := models.StudyStats{
		TotalTime:   36000,  // 10小时
		DailyTime:   7200,   // 2小时
		WeeklyTime:  18000,  // 5小时
		MonthlyTime: 36000,  // 10小时
		CategoryDistribution: map[string]int{
			"编程": 18000, // 5小时
			"外语": 10800, // 3小时
			"数学": 7200,  // 2小时
		},
	}

	return Success(c, stats, "获取学习统计数据成功")
} 