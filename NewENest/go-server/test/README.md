# ENest API 自动化测试

本目录包含 ENest 后端 API 的自动化测试用例。

## 测试文件结构

- `common_test.go` - 测试辅助函数和通用代码
- `auth_test.go` - 用户认证 API 测试用例
- `friend_test.go` - 好友相关 API 测试用例
- `study_room_test.go` - 自习室相关 API 测试用例

## 测试环境要求

- Go 1.20 或更高版本
- 确保项目的 `.env` 文件正确配置（或使用环境变量）

## 运行测试

### 运行所有测试

```bash
cd /path/to/NewENest/go-server
go test ./test/... -v
```

### 运行特定测试文件

```bash
cd /path/to/NewENest/go-server
go test ./test/auth_test.go ./test/common_test.go -v
```

### 运行特定测试函数

```bash
cd /path/to/NewENest/go-server
go test ./test/... -v -run TestLoginUser
```

## 测试实现说明

### 模拟数据

测试用例使用了模拟数据，便于在没有真实数据库的环境下进行测试。这些测试主要验证 API 的接口格式和基本逻辑，而不依赖于真实的数据库状态。

### 测试用户

测试中使用了两个预定义的测试用户：

1. 测试用户1:
   - ID: 1
   - 用户名: test_user
   - 邮箱: test@example.com
   - 密码: password123

2. 测试用户2:
   - ID: 2
   - 用户名: study_buddy
   - 邮箱: buddy@example.com
   - 密码: password123

这些用户在测试开始前自动生成令牌，用于需要认证的 API 端点。

## 测试覆盖的 API 端点

### 认证 API

- `POST /api/v1/auth/register` - 用户注册
- `POST /api/v1/auth/login` - 用户登录

### 好友 API

- `GET /api/v1/friends` - 获取好友列表
- `GET /api/v1/friends/requests` - 获取好友请求
- `POST /api/v1/friends/requests` - 发送好友请求
- `POST /api/v1/friends/requests/response` - 响应好友请求
- `DELETE /api/v1/friends/:id` - 删除好友

### 自习室 API

- `GET /api/v1/study-rooms` - 获取自习室列表
- `GET /api/v1/study-rooms/:id` - 获取自习室详情
- `POST /api/v1/study-rooms` - 创建自习室
- `POST /api/v1/study-rooms/join` - 加入自习室

## 扩展测试

若要添加新的测试用例，建议按照以下步骤：

1. 确定要测试的 API 端点
2. 在对应的测试文件中添加新的测试函数
3. 使用 `t.Run()` 为每个测试场景创建子测试
4. 使用 `sendRequest()` 和 `parseResponse()` 辅助函数发送请求并解析响应
5. 使用 `assert` 包进行断言验证

示例：

```go
func TestNewFeature(t *testing.T) {
    app := setupTestServer()

    t.Run("成功场景", func(t *testing.T) {
        // 准备请求数据
        reqBody := map[string]interface{}{
            "key": "value",
        }

        // 发送请求
        resp := sendRequest(t, app, "POST", "/api/v1/some-endpoint", reqBody, testUser1.Token)
        
        // 验证响应
        assert.Equal(t, http.StatusOK, resp.StatusCode)
        
        // 解析并验证响应数据
        var response APIResponse
        parseResponse(t, resp, &response)
        assert.Equal(t, 200, response.Code)
    })
} 