package handlers

import (
	"net/http"
	
	"github.com/gofiber/fiber/v2"
)

// Response 表示API响应的通用结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success 返回成功响应
func Success(c *fiber.Ctx, data interface{}, message string) error {
	if message == "" {
		message = "操作成功"
	}
	
	return c.Status(http.StatusOK).JSON(Response{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}

// Created 返回创建成功响应
func Created(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(fiber.StatusCreated).JSON(Response{
		Code:    fiber.StatusCreated,
		Message: message,
		Data:    data,
	})
}

// BadRequest 返回错误请求响应
func BadRequest(c *fiber.Ctx, message string, data interface{}) error {
	if message == "" {
		message = "请求参数错误"
	}
	
	return c.Status(http.StatusBadRequest).JSON(Response{
		Code:    http.StatusBadRequest,
		Message: message,
		Data:    data,
	})
}

// Unauthorized 返回未授权响应
func Unauthorized(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "未授权的访问"
	}
	
	return c.Status(http.StatusUnauthorized).JSON(Response{
		Code:    http.StatusUnauthorized,
		Message: message,
	})
}

// Forbidden 返回禁止访问响应
func Forbidden(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "禁止访问"
	}
	
	return c.Status(http.StatusForbidden).JSON(Response{
		Code:    http.StatusForbidden,
		Message: message,
	})
}

// NotFound 返回未找到响应
func NotFound(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "资源不存在"
	}
	
	return c.Status(http.StatusNotFound).JSON(Response{
		Code:    http.StatusNotFound,
		Message: message,
	})
}

// ServerError 返回服务器错误响应
func ServerError(c *fiber.Ctx, message string, err error) error {
	if message == "" {
		message = "服务器内部错误"
	}
	
	return c.Status(http.StatusInternalServerError).JSON(Response{
		Code:    http.StatusInternalServerError,
		Message: message,
		Data:    err.Error(),
	})
} 