package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// Response 表示API响应的通用结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// Success 返回成功响应
func Success(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Code:    fiber.StatusOK,
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
func BadRequest(c *fiber.Ctx, message string, err error) error {
	var errors interface{}
	if err != nil {
		errors = err.Error()
	}
	
	return c.Status(fiber.StatusBadRequest).JSON(Response{
		Code:    fiber.StatusBadRequest,
		Message: message,
		Errors:  errors,
	})
}

// Unauthorized 返回未授权响应
func Unauthorized(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(Response{
		Code:    fiber.StatusUnauthorized,
		Message: message,
	})
}

// Forbidden 返回禁止访问响应
func Forbidden(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusForbidden).JSON(Response{
		Code:    fiber.StatusForbidden,
		Message: message,
	})
}

// NotFound 返回未找到响应
func NotFound(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNotFound).JSON(Response{
		Code:    fiber.StatusNotFound,
		Message: message,
	})
}

// ServerError 返回服务器错误响应
func ServerError(c *fiber.Ctx, message string, err error) error {
	var errors interface{}
	if err != nil {
		errors = err.Error()
	}
	
	return c.Status(fiber.StatusInternalServerError).JSON(Response{
		Code:    fiber.StatusInternalServerError,
		Message: message,
		Errors:  errors,
	})
} 