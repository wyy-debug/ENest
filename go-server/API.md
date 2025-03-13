# ENest Go Server API 文档

## 目录
- [用户认证](#用户认证)
  - [用户注册](#用户注册)
  - [用户登录](#用户登录)
  - [用户登出](#用户登出)
- [自习室管理](#自习室管理)
  - [创建自习室](#创建自习室)
  - [获取自习室详情](#获取自习室详情)
  - [加入自习室](#加入自习室)
  - [通过分享链接加入自习室](#通过分享链接加入自习室)
  - [离开自习室](#离开自习室)
  - [销毁自习室](#销毁自习室)

## 用户认证

### 用户注册

**请求方法**：POST

**请求URL**：`/api/auth/register`

**请求参数**：
```json
{
    "username": "用户名",
    "password": "密码",
    "email": "邮箱地址"
}
```

**响应格式**：
```json
{
    "id": 1,
    "username": "用户名",
    "email": "邮箱地址",
    "created_at": "2024-01-01T12:00:00Z"
}
```

### 用户登录

**请求方法**：POST

**请求URL**：`/api/auth/login`

**请求参数**：
```json
{
    "username": "用户名",
    "password": "密码"
}
```

**响应格式**：
```json
{
    "token": "JWT令牌",
    "user": {
        "id": 1,
        "username": "用户名",
        "email": "邮箱地址"
    }
}
```

### 用户登出

**请求方法**：POST

**请求URL**：`/api/auth/logout`

**请求参数**：无

**响应格式**：
```json
{
    "message": "Successfully logged out"
}
```

## 自习室管理

### 创建自习室

**请求方法**：POST

**请求URL**：`/api/study-room/create`

**请求参数**：
```json
{
    "name": "自习室名称",
    "max_members": 10,
    "is_private": false,
    "duration": "2h"  // 持续时间，支持时间单位：ns, us, ms, s, m, h
}
```

**响应格式**：
```json
{
    "id": 1,
    "owner_id": 1,
    "name": "自习室名称",
    "share_link": "unique_share_link",
    "max_members": 10,
    "is_private": false,
    "created_at": "2024-01-01T12:00:00Z",
    "expires_at": "2024-01-01T14:00:00Z"
}
```

### 获取自习室详情

**请求方法**：POST

**请求URL**：`/api/study-room/detail`

**请求参数**：
```json
{
    "room_id": 1
}
```

**响应格式**：
```json
{
    "id": 1,
    "owner_id": 1,
    "name": "自习室名称",
    "share_link": "unique_share_link",
    "max_members": 10,
    "is_private": false,
    "created_at": "2024-01-01T12:00:00Z",
    "expires_at": "2024-01-01T14:00:00Z",
    "member_count": 2,
    "members": [
        {
            "user_id": 1,
            "username": "用户名",
            "is_anonymous": false,
            "joined_at": "2024-01-01T12:00:00Z"
        }
    ]
}
```

### 加入自习室

**请求方法**：POST

**请求URL**：`/api/study-room/join`

**请求参数**：
```json
{
    "room_id": 1
}
```

**响应格式**：
```json
{
    "id": 1,
    "room_id": 1,
    "user_id": 1,
    "joined_at": "2024-01-01T12:00:00Z"
}
```

### 通过分享链接加入自习室

**请求方法**：GET

**请求URL**：`/api/study-room/join/:shareLink`

**请求参数**：无（shareLink在URL路径中）

**响应格式**：
```json
{
    "id": 1,
    "room_id": 1,
    "user_id": 1,
    "joined_at": "2024-01-01T12:00:00Z"
}
```

### 离开自习室

**请求方法**：POST

**请求URL**：`/api/study-room/leave`

**请求参数**：
```json
{
    "room_id": 1
}
```

**响应格式**：
```json
{
    "message": "Successfully left the study room"
}
```

### 销毁自习室

**请求方法**：POST

**请求URL**：`/api/study-room/destroy`

**请求参数**：
```json
{
    "room_id": 1
}
```

**响应格式**：
```json
{
    "message": "Successfully destroyed the study room"
}
```

## 好友系统

### 发送好友请求

**请求方法**：POST

**请求URL**：`/api/friend/request`

**请求参数**：
```json
{
    "friend_id": 1
}
```

**响应格式**：
```json
{
    "message": "Friend request sent successfully"
}
```

### 处理好友请求

**请求方法**：POST

**请求URL**：`/api/friend/handle-request`

**请求参数**：
```json
{
    "request_id": 1,
    "action": "accept"  // accept 或 reject
}
```

**响应格式**：
```json
{
    "message": "Friend request accepted/rejected successfully"
}
```

### 获取好友列表

**请求方法**：GET

**请求URL**：`/api/friend/list`

**请求参数**：无

**响应格式**：
```json
[
    {
        "id": 1,
        "user_id": 1,
        "friend_id": 2,
        "status": "accepted",
        "created_at": "2024-01-01T12:00:00Z",
        "updated_at": "2024-01-01T12:00:00Z",
        "friend": {
            "username": "好友用户名",
            "email": "好友邮箱"
        }
    }
]
```

### 获取好友请求列表

**请求方法**：GET

**请求URL**：`/api/friend/requests`

**请求参数**：无

**响应格式**：
```json
[
    {
        "id": 1,
        "sender_id": 2,
        "receiver_id": 1,
        "status": "pending",
        "created_at": "2024-01-01T12:00:00Z",
        "updated_at": "2024-01-01T12:00:00Z",
        "sender": {
            "username": "请求者用户名",
            "email": "请求者邮箱"
        }
    }
]
```

### 删除好友

**请求方法**：POST

**请求URL**：`/api/friend/delete`

**请求参数**：
```json
{
    "friend_id": 1
}
```

**响应格式**：
```json
{
    "message": "Friend deleted successfully"
}
```

### 创建好友契约

**请求方法**：POST

**请求URL**：`/api/friend/contract/create`

**请求参数**：
```json
{
    "friend_id": 1,
    "contract_type": "study_buddy",  // study_buddy, accountability_partner
    "contract_terms": "契约条款内容"
}
```

**响应格式**：
```json
{
    "id": 1,
    "user_id": 1,
    "friend_id": 2,
    "contract_type": "study_buddy",
    "contract_terms": "契约条款内容",
    "status": "active",
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
}
```

### 终止好友契约

**请求方法**：POST

**请求URL**：`/api/friend/contract/terminate`

**请求参数**：
```json
{
    "contract_id": 1
}
```

**响应格式**：
```json
{
    "message": "Contract terminated successfully"
}
```

### 获取好友契约列表

**请求方法**：GET

**请求URL**：`/api/friend/contract/list`

**请求参数**：无

**响应格式**：
```json
[
    {
        "id": 1,
        "user_id": 1,
        "friend_id": 2,
        "contract_type": "study_buddy",
        "contract_terms": "契约条款内容",
        "status": "active",
        "created_at": "2024-01-01T12:00:00Z",
        "updated_at": "2024-01-01T12:00:00Z",
        "friend": {
            "username": "好友用户名",
            "email": "好友邮箱"
        }
    }
]
```

### 好友聊天WebSocket连接

**WebSocket URL**：`/ws/friend/chat`

**发送消息格式**：
```json
{
    "sender_id": 1,
    "receiver_id": 2,
    "message_type": "text",  // text, image
    "content": "消息内容"
}
```

**接收消息格式**：
```json
{
    "id": 1,
    "sender_id": 1,
    "receiver_id": 2,
    "message_type": "text",
    "content": "消息内容",
    "is_read": false,
    "created_at": "2024-01-01T12:00:00Z"
}
```

## 错误响应

当API调用出现错误时，将返回以下格式的响应：

```json
{
    "error": "错误信息描述"
}
```

常见错误状态码：
- 400：请求参数错误
- 401：未授权或会话已过期
- 403：权限不足
- 404：资源不存在
- 409：用户名或邮箱已存在
- 422：无效的用户凭据
- 500：服务器内部错误