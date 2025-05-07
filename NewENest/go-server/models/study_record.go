package models

import (
	"time"
)

// StudyRecord 表示学习记录模型
type StudyRecord struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Duration    int       `json:"duration"` // 单位：秒
	Category    string    `json:"category"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// StudyRecordCreate 表示创建学习记录的请求数据
type StudyRecordCreate struct {
	StartTime   time.Time `json:"start_time" validate:"required"`
	EndTime     time.Time `json:"end_time" validate:"required,gtfield=StartTime"`
	Category    string    `json:"category" validate:"required"`
	Description string    `json:"description" validate:"max=500"`
}

// StudyRecordUpdate 表示更新学习记录的请求数据
type StudyRecordUpdate struct {
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time" validate:"omitempty,gtfield=StartTime"`
	Category    string    `json:"category"`
	Description string    `json:"description" validate:"omitempty,max=500"`
}

// StudyStats 表示学习统计数据
type StudyStats struct {
	TotalTime        int            `json:"total_time"` // 总学习时间（秒）
	DailyTime        int            `json:"daily_time"` // 今日学习时间（秒）
	WeeklyTime       int            `json:"weekly_time"` // 本周学习时间（秒）
	MonthlyTime      int            `json:"monthly_time"` // 本月学习时间（秒）
	CategoryDistribution map[string]int `json:"category_distribution"` // 分类占比
}

// StudyRecordRepository 学习记录数据访问接口
type StudyRecordRepository interface {
	// 基础CRUD操作
	FindByID(id int) (*StudyRecord, error)
	Create(record *StudyRecord) error
	Update(record *StudyRecord) error
	Delete(id int) error
	
	// 查询操作
	FindByUserID(userID int, page, pageSize int) ([]StudyRecord, int, error)
	FindByRoomID(roomID int, page, pageSize int) ([]StudyRecord, int, error)
	FindByDateRange(userID int, startDate, endDate time.Time) ([]StudyRecord, error)
	
	// 统计操作
	GetTotalStudyTime(userID int) (int, error)
	GetDailyStudyTime(userID int, date time.Time) (int, error)
	GetWeeklyStudyTime(userID int, startOfWeek time.Time) (map[string]int, error)
	GetMonthlyStudyTime(userID int, year int, month time.Month) (map[string]int, error)
}

// StudyRecordCreateDTO 学习记录创建数据传输对象
type StudyRecordCreateDTO struct {
	RoomID       *int      `json:"room_id,omitempty"`
	StartTime    time.Time `json:"start_time" validate:"required"`
	EndTime      *time.Time `json:"end_time,omitempty"`
	Duration     *int      `json:"duration,omitempty"`
	Interruptions int      `json:"interruptions" validate:"min=0"`
	Notes        string    `json:"notes,omitempty"`
}

// StudyRecordUpdateDTO 学习记录更新数据传输对象
type StudyRecordUpdateDTO struct {
	EndTime      *time.Time `json:"end_time,omitempty"`
	Duration     *int      `json:"duration,omitempty"`
	Interruptions *int     `json:"interruptions,omitempty" validate:"omitempty,min=0"`
	Notes        *string   `json:"notes,omitempty"`
}

// StudyStatsDTO 学习统计数据传输对象
type StudyStatsDTO struct {
	TotalStudyTime   int             `json:"total_study_time"`   // 总学习时间（分钟）
	DailyAverage     int             `json:"daily_average"`      // 日均学习时间（分钟）
	WeeklyAverage    int             `json:"weekly_average"`     // 周均学习时间（分钟）
	CurrentStreak    int             `json:"current_streak"`     // 当前连续学习天数
	LongestStreak    int             `json:"longest_streak"`     // 最长连续学习天数
	TotalSessions    int             `json:"total_sessions"`     // 总学习次数
	TotalDays        int             `json:"total_days"`         // 总学习天数
	DailyDistribution map[string]int `json:"daily_distribution"` // 每日分布
	WeeklyDistribution map[string]int `json:"weekly_distribution"` // 每周分布
}

// DailyStudyDTO 每日学习数据传输对象
type DailyStudyDTO struct {
	Date        string `json:"date"`         // 日期，格式：YYYY-MM-DD
	Duration    int    `json:"duration"`     // 持续时间（分钟）
	Sessions    int    `json:"sessions"`     // 学习次数
	Interruptions int  `json:"interruptions"` // 中断次数
}

// CalculateDuration 计算学习持续时间
func (r *StudyRecord) CalculateDuration() int {
	if r.EndTime.IsZero() {
		return 0
	}
	
	duration := r.EndTime.Sub(r.StartTime)
	return int(duration.Minutes())
} 