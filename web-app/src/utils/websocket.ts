import { ElMessage } from 'element-plus';
import { CRYPTO_CONFIG } from '../config/crypto';
import protoRoot from '../proto/index';

// 使用新生成的MessageType枚举
const { MessageType } = protoRoot.proto;

interface WebSocketMessage {
  type: typeof MessageType;
  timestamp: number;
  payload: Uint8Array;
  session_id: string;
}

class WebSocketClient {
  private static instance: WebSocketClient;
  private ws: WebSocket | null = null;
  private reconnectAttempts = 0;
  private maxReconnectAttempts = 5;
  private reconnectInterval = 3000;
  private heartbeatInterval = 30000;
  private heartbeatTimer: ReturnType<typeof setInterval> | null = null;
  private messageHandlers: Map<typeof MessageType, (payload: Uint8Array) => void> = new Map();
  private cryptoKey: CryptoKey | null = null;

  private constructor() {}

  public static getInstance(): WebSocketClient {
    if (!WebSocketClient.instance) {
      WebSocketClient.instance = new WebSocketClient();
    }
    return WebSocketClient.instance;
  }

  private async initCrypto(): Promise<void> {
    // 从配置文件中读取密钥
    const keyData = CRYPTO_CONFIG.AES_KEY;
    
    this.cryptoKey = await window.crypto.subtle.importKey(
      'raw',
      keyData,
      {
        name: 'AES-GCM',
        length: 256
      },
      true,
      ['encrypt', 'decrypt']
    );
  }

  private async encryptMessage(data: Uint8Array): Promise<Uint8Array> {
    if (!this.cryptoKey) {
      await this.initCrypto();
    }

    // 生成12字节的随机nonce，与服务器端保持一致
    const nonce = window.crypto.getRandomValues(new Uint8Array(12));
    
    // 使用AES-GCM模式加密数据，与服务器端保持一致
    const encryptedData = await window.crypto.subtle.encrypt(
      {
        name: 'AES-GCM',
        iv: nonce
      },
      this.cryptoKey!,
      data
    );

    // 按照服务器端的格式组合nonce和密文
    const result = new Uint8Array(nonce.length + encryptedData.byteLength);
    result.set(nonce);
    result.set(new Uint8Array(encryptedData), nonce.length);
    return result;
  }

  private async decryptMessage(data: Uint8Array): Promise<Uint8Array> {
    if (!this.cryptoKey) {
      await this.initCrypto();
    }

    // 按照服务器端的格式提取nonce和密文
    const nonce = data.slice(0, 12);
    const ciphertext = data.slice(12);

    // 使用AES-GCM模式解密数据
    const decryptedData = await window.crypto.subtle.decrypt(
      {
        name: 'AES-GCM',
        iv: nonce
      },
      this.cryptoKey!,
      ciphertext
    );

    return new Uint8Array(decryptedData);
  }

  private onOpenHandlers: Array<() => Promise<void>> = [];
  private onErrorHandlers: Array<(error: Error) => void> = [];

  public onOpen(handler: () => Promise<void>): void {
    this.onOpenHandlers.push(handler);
  }

  public onError(handler: (error: Error) => void): void {
    this.onErrorHandlers.push(handler);
  }

  public connect(url: string): void {
    if (this.ws?.readyState === WebSocket.OPEN) return;

    this.ws = new WebSocket(url);

    this.ws.onopen = async () => {
      console.log('WebSocket connected');
      this.reconnectAttempts = 0;
      await this.initCrypto();
      //this.startHeartbeat();
      //this.authenticate();
      this.onOpenHandlers.forEach(handler => handler());
    };

    this.ws.onmessage = async (event) => {
      try {
        if (event.data instanceof Blob) {
          const buffer = await event.data.arrayBuffer();
          const encryptedData = new Uint8Array(buffer);
          const decryptedData = await this.decryptMessage(encryptedData);
          
          // 使用新的protobuf解析消息
          const message = protoRoot.proto.Message.decode(decryptedData);
          
          console.log('Received message:', message);
          if (message.type === MessageType.HEARTBEAT) {
            console.log('Received heartbeat');
            this.sendMessage(MessageType.HEARTBEAT, new Uint8Array([]));
            return;
          }

          const handler = this.messageHandlers.get(message.type);
          if (handler) {
            handler(message.payload);
          } else {
            console.warn(`No handler registered for message type: ${message.type}`);
          }
        }
      } catch (error) {
        console.error('Failed to process message:', error);
      }
    };

    this.ws.onclose = () => {
      console.log('WebSocket disconnected');
      this.stopHeartbeat();
      this.reconnect();
    };

    this.ws.onerror = (error: Event) => {
      const wsError = error instanceof Error ? error : new Error('WebSocket连接错误');
      console.error('WebSocket error:', wsError);
      ElMessage.error('WebSocket连接错误');
      this.onErrorHandlers.forEach(handler => handler(wsError));
    };
  }

  private reconnect(): void {
    if (this.reconnectAttempts >= this.maxReconnectAttempts) {
      ElMessage.error('WebSocket重连失败，请刷新页面重试');
      return;
    }

    setTimeout(() => {
      this.reconnectAttempts++;
      this.connect(this.ws!.url);
    }, this.reconnectInterval);
  }

  private startHeartbeat(): void {
    this.heartbeatTimer = setInterval(() => {
      this.sendMessage(MessageType.HEARTBEAT, new Uint8Array([]));
    }, this.heartbeatInterval);
  }

  private stopHeartbeat(): void {
    if (this.heartbeatTimer) {
      clearInterval(this.heartbeatTimer);
      this.heartbeatTimer = null;
    }
  }

  private authenticate(): void {
    const token = localStorage.getItem('session_token');
    if (token == 'undefined' || token == null) {
      console.log('No session token found, skipping authentication');
      ElMessage.warning('未找到登录凭证，请先登录');
      return;
    }

    try {
      const authMessage = protoRoot.proto.AuthMessage.create({
        token: token,
        deviceId: '123123213'
      });
      const authPayload = protoRoot.proto.AuthMessage.encode(authMessage).finish();
      this.sendMessage(MessageType.AUTH, authPayload);
    } catch (error) {
      console.error('Failed to create auth message:', error);
      ElMessage.error('认证消息创建失败');
    }
  }

  public async sendMessage(type: typeof MessageType, payload: Uint8Array): Promise<void> {
    console.log('MessageType' + type);
    if (this.ws?.readyState !== WebSocket.OPEN) {
      ElMessage.error('WebSocket未连接');
      return;
    }

    try {
      // 构建protobuf消息
      const message = protoRoot.proto.Message.create({
        type: type,
        timestamp: Math.floor(Date.now() / 1000),
        payload: payload,
        sessionId: localStorage.getItem('session_id') || ''
      });
      console.log('send message Type' + type);
      // 使用protobuf序列化消息
      const messageBuffer = protoRoot.proto.Message.encode(message).finish();
      const encryptedData = await this.encryptMessage(messageBuffer);
      this.ws.send(encryptedData);
    } catch (error) {
      console.error('Failed to send message:', error);
      ElMessage.error('发送消息失败');
    }
  }

  public registerHandler(type: typeof MessageType, handler: (payload: Uint8Array) => void): void {
    this.messageHandlers.set(type, handler);
  }

  public disconnect(): void {
    if (this.ws) {
      this.ws.close();
      this.ws = null;
    }
    this.stopHeartbeat();
    this.cryptoKey = null;
  }
}

export const wsClient = WebSocketClient.getInstance();