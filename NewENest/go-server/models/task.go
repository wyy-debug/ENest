package models

import (
	"time"
)

// Task 表示任务模型
type Task struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Priority    string    `json:"priority"` // 高、中、低
	Status      string    `json:"status"`   // 未开始、进行中、已完成
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TaskCreate 表示创建任务的请求数据
type TaskCreate struct {
	Title       string    `json:"title" validate:"required,min=1,max=100"`
	Description string    `json:"description" validate:"max=1000"`
	Deadline    time.Time `json:"deadline" validate:"required"`
	Priority    string    `json:"priority" validate:"required,oneof=高 中 低"`
}

// TaskUpdate 表示更新任务的请求数据
type TaskUpdate struct {
	Title       string    `json:"title" validate:"omitempty,min=1,max=100"`
	Description string    `json:"description" validate:"omitempty,max=1000"`
	Deadline    time.Time `json:"deadline"`
	Priority    string    `json:"priority" validate:"omitempty,oneof=高 中 低"`
	Status      string    `json:"status" validate:"omitempty,oneof=未开始 进行中 已完成"`
} 