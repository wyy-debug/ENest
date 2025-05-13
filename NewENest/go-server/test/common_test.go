package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"

	"NewENest/go-server/handlers"
	"NewENest/go-server/middleware"
	"NewENest/go-server/database"
)

// TestUser 测试用户结构
type TestUser struct {
	ID       int
	Username string
	Email    string
	Password string
	Token    string
}

// 测试用户数据
var (
	testUser1 = TestUser{
		ID:       1,
		Username: "test_user",
		Email:    "test@example.com",
		Password: "password123",
	}

	testUser2 = TestUser{
		ID:       2,
		Username: "study_buddy",
		Email:    "buddy@example.com",
		Password: "password123",
	}
)

// 创建测试服务器
func setupTestServer() *fiber.App {
	// 加载环境变量
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println("未找到.env文件，使用默认配置")
	}

	// 创建Fiber应用
	app := fiber.New(fiber.Config{
		AppName:      "ENest API Test",
		ErrorHandler: defaultErrorHandler,
	})

	// 设置路由
	setupRoutes(app)

	return app
}

// 默认错误处理
func defaultErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": err.Error(),
	})
}

// 设置测试路由
func setupRoutes(app *fiber.App) {
	// 在此简化了路由设置，实际测试中可能需要更完整的路由设置
	// 或者直接使用 routes.SetupRoutes(app, db)
	
	api := app.Group("/api/v1")

	// 公开路由
	api.Post("/auth/register", handlers.Register)
	api.Post("/auth/login", handlers.LoginV2)

	// 自习室公开路由
	api.Get("/study-rooms", handlers.GetStudyRooms)
	api.Get("/study-rooms/:id", handlers.GetStudyRoom)
	api.Post("/study-rooms/join", handlers.JoinStudyRoom)

	// 添加模拟的好友仓库
	mockRepo := NewMockFriendRepository() // 使用相同的模拟实现但类型转换为models.FriendRepository
	
	// 设置全局仓库实例，确保数据库层可以访问到
	database.SetFriendRepositoryInstance(mockRepo)
	
	// 创建处理器
	friendHandler := handlers.NewFriendHandler(mockRepo)

	// 受保护的路由
	protected := api.Group("", middleware.JWTMiddleware())
	
	// 注册好友路由
	friendHandler.RegisterRoutes(protected)

	// 用户相关路由
	protected.Get("/user", handlers.GetCurrentUser)
	protected.Put("/user", handlers.UpdateCurrentUser)
	protected.Get("/users/search", handlers.SearchUsers)

	// 自习室受保护路由
	protected.Post("/study-rooms", handlers.CreateStudyRoom)

	// 404处理
	app.Use(func(c *fiber.Ctx) error {
		return handlers.HandleNotFound(c, "未找到请求的资源")
	})
}

// 生成测试JWT令牌 - 修正后使用与服务器相同的JWTClaims结构
func generateToken(userID int, username string) (string, error) {
	// 使用与应用相同的JWT密钥
	jwtSecret := os.Getenv("NEST_JWT_SECRET")
	if jwtSecret == "" {
		// 确保此处的密钥与服务器配置的密钥相同
		jwtSecret = "newenest_secret_key"
	}

	// 创建与服务器端相同的JWT声明结构
	type JWTClaims struct {
		UserID int    `json:"user_id"`
		Email  string `json:"email"`
		jwt.StandardClaims
	}

	// 构造邮箱地址 - 与原测试保持一致
	email := "test" + fmt.Sprintf("%d", userID) + "@example.com"
	
	// 创建JWT声明
	claims := JWTClaims{
		UserID: userID,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "newenest",
		},
	}

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名令牌
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// 发送HTTP请求并返回响应
func sendRequest(t *testing.T, app *fiber.App, method, url string, body interface{}, token string) *http.Response {
	// 将请求体转换为JSON
	var reqBody []byte
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			t.Fatalf("无法将请求体转换为JSON: %v", err)
		}
	}

	// 创建HTTP请求
	req := httptest.NewRequest(method, url, bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	
	// 添加授权头（如果提供了令牌）
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	// 发送请求并获取响应
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("发送请求时出错: %v", err)
	}

	return resp
}

// 解析响应体
func parseResponse(t *testing.T, resp *http.Response, target interface{}) {
	err := json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		t.Fatalf("解析响应体时出错: %v", err)
	}
}

// 通用响应结构
type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 设置测试套件
func TestMain(m *testing.M) {
	// 在所有测试之前进行的设置
	
	// 生成测试用户的令牌
	var err error
	testUser1.Token, err = generateToken(testUser1.ID, testUser1.Username)
	if err != nil {
		fmt.Printf("无法为测试用户1生成令牌: %v\n", err)
		os.Exit(1)
	}

	testUser2.Token, err = generateToken(testUser2.ID, testUser2.Username)
	if err != nil {
		fmt.Printf("无法为测试用户2生成令牌: %v\n", err)
		os.Exit(1)
	}

	// 运行测试
	code := m.Run()

	// 在所有测试之后进行的清理
	
	// 退出，返回测试结果
	os.Exit(code)
} 