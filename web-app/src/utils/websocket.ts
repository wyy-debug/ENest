import { MessageType } from '../proto/message';
import { ElMessage } from 'element-plus';

interface WebSocketMessage {
  type: MessageType;
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
  private heartbeatTimer: NodeJS.Timeout | null = null;
  private messageHandlers: Map<MessageType, (payload: Uint8Array) => void> = new Map();

  private constructor() {}

  public static getInstance(): WebSocketClient {
    if (!WebSocketClient.instance) {
      WebSocketClient.instance = new WebSocketClient();
    }
    return WebSocketClient.instance;
  }

  public connect(url: string): void {
    if (this.ws?.readyState === WebSocket.OPEN) return;

    this.ws = new WebSocket(url);

    this.ws.onopen = () => {
      console.log('WebSocket connected');
      this.reconnectAttempts = 0;
      this.startHeartbeat();
      this.authenticate();
    };

    this.ws.onmessage = (event) => {
      try {
        const message: WebSocketMessage = JSON.parse(event.data);
        const handler = this.messageHandlers.get(message.type);
        if (handler) {
          handler(message.payload);
        }
      } catch (error) {
        console.error('Failed to parse message:', error);
      }
    };

    this.ws.onclose = () => {
      console.log('WebSocket disconnected');
      this.stopHeartbeat();
      this.reconnect();
    };

    this.ws.onerror = (error) => {
      console.error('WebSocket error:', error);
      ElMessage.error('WebSocket连接错误');
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
    if (token) {
      const authPayload = new TextEncoder().encode(JSON.stringify({ token }));
      this.sendMessage(MessageType.AUTH, authPayload);
    }
  }

  public sendMessage(type: MessageType, payload: Uint8Array): void {
    if (this.ws?.readyState !== WebSocket.OPEN) {
      ElMessage.error('WebSocket未连接');
      return;
    }

    const message: WebSocketMessage = {
      type,
      timestamp: Date.now(),
      payload,
      session_id: localStorage.getItem('session_token') || ''
    };

    this.ws.send(JSON.stringify(message));
  }

  public registerHandler(type: MessageType, handler: (payload: Uint8Array) => void): void {
    this.messageHandlers.set(type, handler);
  }

  public disconnect(): void {
    if (this.ws) {
      this.ws.close();
      this.ws = null;
    }
    this.stopHeartbeat();
  }
}

export const wsClient = WebSocketClient.getInstance();