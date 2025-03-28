# ENest

## Protobuf 代码生成

本项目使用 Protocol Buffers 进行前后端通信。以下是生成 Go 和 TypeScript 代码的步骤：

### 前提条件

1. 安装 protoc 编译器
2. 安装 Go protobuf 插件：
   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@latest
   ```
3. 安装 Node.js 依赖：
   ```bash
   npm install
   ```

### 生成代码

运行以下命令生成 Go 和 TypeScript 代码：

```bash
npm run generate:proto
```

这将：
- 从 `proto/*.proto` 文件生成 Go 代码到 `go-server/proto/` 目录
- 从 `proto/*.proto` 文件生成 TypeScript 代码到 `web-app/src/proto/` 目录

### 在前端项目中使用

1. 导入protobuf消息类型和枚举

```typescript
import protoRoot from '../proto/index';

// 导入消息类型枚举
const { MessageType } = protoRoot.proto;

// 使用其他枚举
const { Operation: StudyRoomOperation } = protoRoot.proto.StudyRoomMessage;
const { Operation: FriendOperation } = protoRoot.proto.FriendMessage;
```

2. 创建和编码消息

```typescript
// 创建消息
const chatMessage = protoRoot.proto.ChatMessage.create({
  senderId: 1001,
  receiverId: 2002,
  content: '你好，这是一条测试消息！',
  messageType: 'text'
});

// 编码消息为二进制
const payload = protoRoot.proto.ChatMessage.encode(chatMessage).finish();

// 作为基础消息的payload发送
const message = protoRoot.proto.Message.create({
  type: MessageType.CHAT,
  timestamp: Date.now(),
  payload: payload,
  sessionId: 'session-id'
});

// 编码基础消息
const messageBuffer = protoRoot.proto.Message.encode(message).finish();
```

3. 解码消息

```typescript
// 解码消息
const message = protoRoot.proto.Message.decode(messageBuffer);

// 根据类型解码payload
if (message.type === MessageType.CHAT) {
  const chatMessage = protoRoot.proto.ChatMessage.decode(message.payload);
  console.log('发送者:', chatMessage.senderId);
  console.log('内容:', chatMessage.content);
}
```

### 在Go服务器中使用

```go
import (
  "log"
  pb "go-server/proto"
)

func handleMessage(data []byte) {
  // 解析基础消息
  message := &pb.Message{}
  if err := message.UnmarshalVT(data); err != nil {
    log.Printf("解析消息失败: %v", err)
    return
  }
  
  // 根据类型处理payload
  switch message.Type {
  case pb.MessageType_CHAT:
    chatMessage := &pb.ChatMessage{}
    if err := chatMessage.UnmarshalVT(message.Payload); err != nil {
      log.Printf("解析聊天消息失败: %v", err)
      return
    }
    log.Printf("接收到聊天消息: %s", chatMessage.Content)
    
  case pb.MessageType_STUDY_ROOM:
    roomMessage := &pb.StudyRoomMessage{}
    if err := roomMessage.UnmarshalVT(message.Payload); err != nil {
      log.Printf("解析自习室消息失败: %v", err)
      return
    }
    log.Printf("接收到自习室消息, 操作: %v", roomMessage.Operation)
  }
}
```

### 自定义

如需自定义生成过程，请编辑 `scripts/generateProto.js` 文件。