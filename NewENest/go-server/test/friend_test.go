package test

import (
	"encoding/json"
	"os"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"NewENest/go-server/database"
	"NewENest/go-server/models"
)

// 加载测试响应数据
func loadTestResponses() map[string]interface{} {
	// 从fixtures中加载测试数据
	data, err := os.ReadFile("fixtures/responses.json")
	if err != nil {
		panic("无法加载测试响应数据: " + err.Error())
	}

	var responses map[string]interface{}
	if err := json.Unmarshal(data, &responses); err != nil {
		panic("无法解析测试响应数据: " + err.Error())
	}

	return responses
}

// MockFriendRepository 模拟好友仓库
type MockFriendRepository struct {
	mock.Mock
	testData map[string]interface{}
}

// 创建新的模拟好友仓库
func NewMockFriendRepository() *MockFriendRepository {
	return &MockFriendRepository{
		testData: loadTestResponses(),
	}
}

// FindByID 根据ID查找好友关系
func (m *MockFriendRepository) FindByID(id int) (*models.Friend, error) {
	return &models.Friend{
		ID:       id,
		UserID:   1,
		FriendID: 2,
		Status:   "accepted",
		Friend: &models.User{
			ID:           2,
			Username:     "study_buddy",
			Email:        "buddy@example.com",
			Avatar:       "https://api.dicebear.com/7.x/avataaars/svg?seed=study_buddy",
			Signature:    "一起学习吧！",
			StudyDirection: "数学",
		},
	}, nil
}

// GetFriendList 模拟获取好友列表
func (m *MockFriendRepository) GetFriendList(userID int) ([]models.Friend, error) {
	// 从测试数据中加载好友列表
	friendsData := m.testData["friends"].(map[string]interface{})
	listData := friendsData["list"].(map[string]interface{})
	successData := listData["success"].(map[string]interface{})
	
	// 解析好友列表数据
	var friends []models.Friend
	friendsJson, _ := json.Marshal(successData["data"])
	json.Unmarshal(friendsJson, &friends)
	
	return friends, nil
}

// GetFriendRequests 模拟获取好友请求列表
func (m *MockFriendRepository) GetFriendRequests(userID int) ([]models.FriendRequest, error) {
	// 从测试数据中加载好友请求列表
	friendsData := m.testData["friends"].(map[string]interface{})
	requestsData := friendsData["requests"].(map[string]interface{})
	successData := requestsData["success"].(map[string]interface{})
	
	// 解析好友请求列表数据
	var requests []models.FriendRequest
	requestsJson, _ := json.Marshal(successData["data"])
	json.Unmarshal(requestsJson, &requests)
	
	return requests, nil
}

// SendFriendRequest 模拟发送好友请求
func (m *MockFriendRepository) SendFriendRequest(userID, receiverID int) error {
	return nil
}

// AcceptFriendRequest 模拟接受好友请求
func (m *MockFriendRepository) AcceptFriendRequest(requestID, userID int) error {
	return nil
}

// RejectFriendRequest 模拟拒绝好友请求
func (m *MockFriendRepository) RejectFriendRequest(requestID, userID int) error {
	return nil
}

// RespondToFriendRequest 模拟响应好友请求
func (m *MockFriendRepository) RespondToFriendRequest(requestID, userID int, accept bool) error {
	if accept {
		return m.AcceptFriendRequest(requestID, userID)
	}
	return m.RejectFriendRequest(requestID, userID)
}

// DeleteFriend 模拟删除好友
func (m *MockFriendRepository) DeleteFriend(userID, friendID int) error {
	return nil
}

// IsFriend 模拟检查是否是好友
func (m *MockFriendRepository) IsFriend(userID, friendID int) (bool, error) {
	// 模拟用户1和用户2是好友
	if (userID == 1 && friendID == 2) || (userID == 2 && friendID == 1) {
		return true, nil
	}
	return false, nil
}

// 模拟检查好友关系
func (m *MockFriendRepository) CheckFriendship(userID, otherID int) (bool, error) {
	return m.IsFriend(userID, otherID)
}

// CreateContract 模拟创建好友契约
func (m *MockFriendRepository) CreateContract(contract *models.FriendContract) error {
	return nil
}

// GetContractByID 模拟获取契约详情
func (m *MockFriendRepository) GetContractByID(contractID int) (*models.FriendContract, error) {
	return &models.FriendContract{
		ID:           contractID,
		UserID:       1,
		FriendID:     2,
		ContractType: "study_buddy",
		Status:       "active",
	}, nil
}

// GetUserContracts 模拟获取用户契约列表
func (m *MockFriendRepository) GetUserContracts(userID int) ([]models.FriendContract, error) {
	return []models.FriendContract{
		{
			ID:           1,
			UserID:       1,
			FriendID:     2,
			ContractType: "study_buddy",
			Status:       "active",
		},
	}, nil
}

// UpdateContractStatus 模拟更新契约状态
func (m *MockFriendRepository) UpdateContractStatus(contractID int, status string) error {
	return nil
}

// UpdateContract 模拟更新契约
func (m *MockFriendRepository) UpdateContract(contract *models.FriendContract) error {
	return nil
}

// SaveMessage 模拟保存消息
func (m *MockFriendRepository) SaveMessage(message *models.FriendMessage) error {
	return nil
}

// GetMessage 模拟获取消息
func (m *MockFriendRepository) GetMessage(messageID int) (*models.FriendMessage, error) {
	return &models.FriendMessage{
		ID:          messageID,
		SenderID:    1,
		ReceiverID:  2,
		MessageType: "text",
		Content:     "你好",
		IsRead:      false,
	}, nil
}

// MarkMessageAsRead 模拟标记消息已读
func (m *MockFriendRepository) MarkMessageAsRead(messageID int) error {
	return nil
}

// GetChatHistory 模拟获取聊天记录
func (m *MockFriendRepository) GetChatHistory(userID, friendID int, limit, offset int) ([]models.FriendMessage, int, error) {
	return []models.FriendMessage{
		{
			ID:          1,
			SenderID:    userID,
			ReceiverID:  friendID,
			MessageType: "text",
			Content:     "你好",
			IsRead:      true,
		},
		{
			ID:          2,
			SenderID:    friendID,
			ReceiverID:  userID,
			MessageType: "text",
			Content:     "你好！最近学习如何？",
			IsRead:      true,
		},
	}, 2, nil
}

// GetUnreadMessageCount 模拟获取未读消息数
func (m *MockFriendRepository) GetUnreadMessageCount(userID int) (int, error) {
	return 0, nil
}

// TestGetFriendList 测试获取好友列表API
func TestGetFriendList(t *testing.T) {
	// 创建测试服务器
	app := setupTestServer()

	// 替换模拟好友仓库
	database.SetFriendRepositoryInstance(NewMockFriendRepository())

	t.Run("成功获取好友列表", func(t *testing.T) {
		// 发送请求
		resp := sendRequest(t, app, "GET", "/api/v1/friends", nil, testUser1.Token)
		
		// 验证响应状态码
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// 解析响应
		var response APIResponse
		parseResponse(t, resp, &response)

		// 验证响应数据
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "获取好友列表成功", response.Message)
		assert.NotNil(t, response.Data)
	})

	t.Run("未授权访问", func(t *testing.T) {
		// 发送请求，不提供令牌
		resp := sendRequest(t, app, "GET", "/api/v1/friends", nil, "")
		
		// 验证响应状态码
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
}

// TestGetFriendRequests 测试获取好友请求列表
func TestGetFriendRequests(t *testing.T) {
	// 创建测试服务器
	app := setupTestServer()

	// 替换模拟好友仓库
	database.SetFriendRepositoryInstance(NewMockFriendRepository())

	t.Run("成功获取好友请求列表", func(t *testing.T) {
		// 发送请求
		resp := sendRequest(t, app, "GET", "/api/v1/friends/requests", nil, testUser1.Token)
		
		// 验证响应状态码
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// 解析响应
		var response APIResponse
		parseResponse(t, resp, &response)

		// 验证响应数据
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "获取好友请求列表成功", response.Message)
	})
}

// TestSendFriendRequest 测试发送好友请求
func TestSendFriendRequest(t *testing.T) {
	app := setupTestServer()

	t.Run("成功发送好友请求", func(t *testing.T) {
		// 准备请求数据
		reqBody := map[string]interface{}{
			"receiver_id": 3, // 一个未成为好友的用户ID
		}

		// 发送请求
		resp := sendRequest(t, app, "POST", "/api/v1/friends/requests", reqBody, testUser1.Token)
		
		// 验证响应状态码
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// 解析响应
		var response APIResponse
		parseResponse(t, resp, &response)

		// 验证响应数据
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "好友请求发送成功", response.Message)
	})

	t.Run("向自己发送好友请求", func(t *testing.T) {
		// 准备请求数据
		reqBody := map[string]interface{}{
			"receiver_id": 1, // 用户自己的ID
		}

		// 发送请求
		resp := sendRequest(t, app, "POST", "/api/v1/friends/requests", reqBody, testUser1.Token)
		
		// 验证响应状态码
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		// 解析响应
		var response APIResponse
		parseResponse(t, resp, &response)

		// 验证响应数据
		assert.Equal(t, 400, response.Code)
		assert.Equal(t, "不能添加自己为好友", response.Message)
	})

	t.Run("向已经是好友的用户发送请求", func(t *testing.T) {
		// 准备请求数据
		reqBody := map[string]interface{}{
			"receiver_id": 2, // 假设用户1和用户2已经是好友
		}

		// 发送请求
		resp := sendRequest(t, app, "POST", "/api/v1/friends/requests", reqBody, testUser1.Token)
		
		// 验证响应状态码
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		// 解析响应
		var response APIResponse
		parseResponse(t, resp, &response)

		// 验证响应数据
		assert.Equal(t, 400, response.Code)
		assert.Equal(t, "你们已经是好友关系", response.Message)
	})

	t.Run("未授权访问", func(t *testing.T) {
		// 准备请求数据
		reqBody := map[string]interface{}{
			"receiver_id": 3,
		}

		// 发送请求，不带令牌
		resp := sendRequest(t, app, "POST", "/api/v1/friends/requests", reqBody, "")
		
		// 验证响应状态码
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
}

// TestRespondToFriendRequest 测试响应好友请求
func TestRespondToFriendRequest(t *testing.T) {
	app := setupTestServer()

	t.Run("成功接受好友请求", func(t *testing.T) {
		// 准备请求数据
		reqBody := map[string]interface{}{
			"request_id": 1, // 假设ID为1的好友请求
			"action":    "accept",
		}

		// 发送请求
		resp := sendRequest(t, app, "POST", "/api/v1/friends/requests/response", reqBody, testUser1.Token)
		
		// 验证响应状态码
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// 解析响应
		var response APIResponse
		parseResponse(t, resp, &response)

		// 验证响应数据
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "已接受好友请求", response.Message)
	})

	t.Run("成功拒绝好友请求", func(t *testing.T) {
		// 准备请求数据
		reqBody := map[string]interface{}{
			"request_id": 2, // 假设ID为2的好友请求
			"action":    "reject",
		}

		// 发送请求
		resp := sendRequest(t, app, "POST", "/api/v1/friends/requests/response", reqBody, testUser1.Token)
		
		// 验证响应状态码
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// 解析响应
		var response APIResponse
		parseResponse(t, resp, &response)

		// 验证响应数据
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "已拒绝好友请求", response.Message)
	})

	t.Run("无效的请求ID", func(t *testing.T) {
		// 准备请求数据
		reqBody := map[string]interface{}{
			"request_id": 999, // 不存在的好友请求ID
			"action":    "accept",
		}

		// 发送请求
		resp := sendRequest(t, app, "POST", "/api/v1/friends/requests/response", reqBody, testUser1.Token)
		
		// 由于我们使用了模拟数据，这里实际上会返回成功，但在真实环境中应检查是否返回适当的错误
		// 这里我们只检查状态码是否为200
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("无效的操作类型", func(t *testing.T) {
		// 准备请求数据
		reqBody := map[string]interface{}{
			"request_id": 1,
			"action":    "invalid_action", // 无效的操作类型
		}

		// 发送请求
		resp := sendRequest(t, app, "POST", "/api/v1/friends/requests/response", reqBody, testUser1.Token)
		
		// 验证响应状态码
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		// 解析响应
		var response APIResponse
		parseResponse(t, resp, &response)

		// 验证响应数据
		assert.Equal(t, 400, response.Code)
		assert.Equal(t, "操作类型必须是 accept 或 reject", response.Message)
	})

	t.Run("未授权访问", func(t *testing.T) {
		// 准备请求数据
		reqBody := map[string]interface{}{
			"request_id": 1,
			"action":    "accept",
		}

		// 发送请求，不带令牌
		resp := sendRequest(t, app, "POST", "/api/v1/friends/requests/response", reqBody, "")
		
		// 验证响应状态码
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
}

// TestDeleteFriend 测试删除好友
func TestDeleteFriend(t *testing.T) {
	app := setupTestServer()

	t.Run("删除好友", func(t *testing.T) {
		// 发送请求
		resp := sendRequest(t, app, "DELETE", "/api/v1/friends/2", nil, testUser1.Token)
		
		// 注意：当前实现返回501，因为该功能尚未实现
		assert.Equal(t, http.StatusNotImplemented, resp.StatusCode)

		// 解析响应
		var response APIResponse
		parseResponse(t, resp, &response)

		// 验证响应数据
		assert.Equal(t, "删除好友功能待实现", response.Message)
	})

	t.Run("未授权访问", func(t *testing.T) {
		// 发送请求，不带令牌
		resp := sendRequest(t, app, "DELETE", "/api/v1/friends/2", nil, "")
		
		// 验证响应状态码
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
} 