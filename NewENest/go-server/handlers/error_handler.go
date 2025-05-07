package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// Response 统一响应结构
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
}

// AppError 应用错误类型
type AppError struct {
	Code    int
	Message string
	Err     error
}

// Error 实现 error 接口
func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

// NewAppError 创建新的应用错误
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// ErrorHandler 统一错误处理
func ErrorHandler(c *fiber.Ctx, err error) error {
	// 默认状态码和消息
	statusCode := fiber.StatusInternalServerError
	message := "服务器内部错误"
	var errMessages []string

	// 根据错误类型处理
	var appError *AppError
	if errors.As(err, &appError) {
		statusCode = appError.Code
		message = appError.Message
		
		if appError.Err != nil {
			errMessages = append(errMessages, appError.Err.Error())
		}
	} else if e, ok := err.(*fiber.Error); ok {
		statusCode = e.Code
		message = e.Message
	} else {
		// 其他错误作为内部错误处理并记录日志
		log.Error().Err(err).Msg("未处理的错误")
		errMessages = append(errMessages, err.Error())
	}

	// 特殊处理400错误
	if statusCode == fiber.StatusBadRequest {
		// 检查是否为验证错误
		errStr := err.Error()
		if strings.Contains(errStr, "validation") || strings.Contains(errStr, "invalid") {
			// 分割验证错误信息
			parts := strings.Split(errStr, ";")
			for _, part := range parts {
				if part = strings.TrimSpace(part); part != "" {
					errMessages = append(errMessages, part)
				}
			}
		}
	}

	// 构建响应
	response := Response{
		Success: false,
		Message: message,
		Errors:  errMessages,
	}

	return c.Status(statusCode).JSON(response)
}

// 返回成功响应
func Success(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(http.StatusOK).JSON(Response{
		Success: true,
		Data:    data,
		Message: message,
	})
}

// 返回错误响应
func Error(c *fiber.Ctx, statusCode int, message string, err error) error {
	return ErrorHandler(c, NewAppError(statusCode, message, err))
}

// 常见错误类型

// BadRequest 返回400错误
func BadRequest(c *fiber.Ctx, message string, err error) error {
	return Error(c, fiber.StatusBadRequest, message, err)
}

// Unauthorized 返回401错误
func Unauthorized(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "未授权的访问"
	}
	return Error(c, fiber.StatusUnauthorized, message, nil)
}

// Forbidden 返回403错误
func Forbidden(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "禁止访问"
	}
	return Error(c, fiber.StatusForbidden, message, nil)
}

// NotFound 返回404错误
func NotFound(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "资源不存在"
	}
	return Error(c, fiber.StatusNotFound, message, nil)
}

// ServerError 返回500错误
func ServerError(c *fiber.Ctx, message string, err error) error {
	if message == "" {
		message = "服务器内部错误"
	}
	return Error(c, fiber.StatusInternalServerError, message, err)
} 