package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"NewENest/go-server/handlers"
)

// TestRegisterUser 测试用户注册
func TestRegisterUser(t *testing.T) {
	app := setupTestServer()

	t.Run("成功注册新用户", func(t *testing.T) {
		// 准备请求体
		registerData := fiber.Map{
			"username": "newuser",
			"email":    "newuser@example.com",
			"password": "password123",
		}

		// 序列化为JSON
		reqBody, _ := json.Marshal(registerData)

		// 创建请求
		req := httptest.NewRequest("POST", "/api/v1/auth/register", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		// 获取响应
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}

		// 验证状态码
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		// 验证响应体
		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatal(err)
		}

		// 验证响应字段
		assert.Equal(t, float64(201), response["code"])
		assert.Equal(t, "注册成功", response["message"])
		assert.NotNil(t, response["data"])
	})

	t.Run("使用已存在的用户名注册", func(t *testing.T) {
		// 准备请求体
		registerData := fiber.Map{
			"username": "test_user", // 已存在的用户名
			"email":    "new@example.com",
			"password": "password123",
		}

		// 序列化为JSON
		reqBody, _ := json.Marshal(registerData)

		// 发送请求
		req := httptest.NewRequest("POST", "/api/v1/auth/register", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		// 获取响应
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}

		// 验证状态码
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("使用已存在的邮箱注册", func(t *testing.T) {
		// 准备请求体
		registerData := fiber.Map{
			"username": "newuser2",
			"email":    "test@example.com", // 已存在的邮箱
			"password": "password123",
		}

		// 序列化为JSON
		reqBody, _ := json.Marshal(registerData)

		// 发送请求
		req := httptest.NewRequest("POST", "/api/v1/auth/register", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		// 获取响应
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}

		// 验证状态码
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("请求参数不完整", func(t *testing.T) {
		// 准备请求体
		registerData := fiber.Map{
			"username": "newuser3",
			// 缺少email
			"password": "password123",
		}

		// 序列化为JSON
		reqBody, _ := json.Marshal(registerData)

		// 发送请求
		req := httptest.NewRequest("POST", "/api/v1/auth/register", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		// 获取响应
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}

		// 验证状态码
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}

// TestLoginUser 测试用户登录
func TestLoginUser(t *testing.T) {
	app := fiber.New()
	app.Post("/login", handlers.LoginV2)

	t.Run("成功登录", func(t *testing.T) {
		// 准备请求体
		loginData := fiber.Map{
			"email":    "test@example.com",
			"password": "password123",
		}

		// 序列化为JSON
		reqBody, _ := json.Marshal(loginData)

		// 发送请求
		req := httptest.NewRequest("POST", "/login", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		// 获取响应
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}

		// 验证状态码
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// 验证响应体
		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatal(err)
		}

		// 验证响应字段
		assert.Equal(t, float64(200), response["code"])
		assert.Equal(t, "登录成功", response["message"])
		assert.NotNil(t, response["data"])

		// 验证数据中包含token和user
		data, ok := response["data"].(map[string]interface{})
		assert.True(t, ok)
		assert.NotNil(t, data["token"])

		// 验证user数据
		user, ok := data["user"].(map[string]interface{})
		assert.True(t, ok)
		assert.Equal(t, "test_user", user["username"])
	})

	t.Run("邮箱不存在", func(t *testing.T) {
		// 准备请求体
		loginData := fiber.Map{
			"email":    "nonexistent@example.com",
			"password": "password123",
		}

		// 序列化为JSON
		reqBody, _ := json.Marshal(loginData)

		// 发送请求
		req := httptest.NewRequest("POST", "/login", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		// 获取响应
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}

		// 验证状态码
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("密码错误", func(t *testing.T) {
		// 准备请求体
		loginData := fiber.Map{
			"email":    "test@example.com",
			"password": "wrongpassword",
		}

		// 序列化为JSON
		reqBody, _ := json.Marshal(loginData)

		// 发送请求
		req := httptest.NewRequest("POST", "/login", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		// 获取响应
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}

		// 验证状态码
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("请求参数不完整", func(t *testing.T) {
		// 准备请求体
		loginData := fiber.Map{
			"email": "test@example.com",
			// 缺少password
		}

		// 序列化为JSON
		reqBody, _ := json.Marshal(loginData)

		// 发送请求
		req := httptest.NewRequest("POST", "/login", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		// 获取响应
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}

		// 验证状态码
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
} 