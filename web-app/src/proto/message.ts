// message.ts
// 从生成的proto文件导入

import protoBundle from './index';
import type { IMessage, IAuthMessage } from './proto';

// 导出消息类型枚举
export enum MessageType {
  UNKNOWN = 0,
  HEARTBEAT = 1,    // 心跳包
  AUTH = 2,         // 认证消息
  CHAT = 3,         // 聊天消息
  SYSTEM = 4,       // 系统消息
  STUDY_ROOM = 5,   // 自习室消息
  FRIEND = 6,       // 好友消息
  PROFILE = 7,      // 个人信息消息
  ERROR = 8,        // 错误消息
  REGISTER = 9,     // 注册消息
}

// 定义操作类型枚举
export enum StudyRoomOp {
  CREATE = 0,      // 创建自习室
  JOIN = 1,        // 加入自习室
  LEAVE = 2,       // 离开自习室
  DESTROY = 3,     // 销毁自习室
  GET_DETAIL = 4,  // 获取详情
}

// 导出好友操作类型枚举
export enum FriendOp {
  SEND_REQUEST = 0,     // 发送好友请求
  HANDLE_REQUEST = 1,   // 处理好友请求
  GET_LIST = 2,         // 获取好友列表
  DELETE = 3,           // 删除好友
  CREATE_CONTRACT = 4,  // 创建好友契约
  TERMINATE_CONTRACT = 5, // 终止好友契约
  GET_CONTRACTS = 6,    // 获取契约列表
}

// 导出个人信息操作类型枚举
export enum ProfileOp {
  UPDATE = 0,    // 更新个人信息
  GET = 1,       // 获取个人信息
}

// 导出消息类型
export type StudyRoomMessage = {
  operation: StudyRoomOp;
  roomId?: number;
  name?: string;
  maxMembers?: number;
  isPrivate?: boolean;
  duration?: string;
}

// 导出好友消息类型
export type FriendMessage = {
  operation: FriendOp;
  friendId?: number;
  action?: string;
  contractType?: string;
  contractTerms?: string;
  contractId?: number;
}

// 导出个人信息消息类型
export type ProfileMessage = {
  operation: ProfileOp;
  username?: string;
  signature?: string;
  studyDirection?: string;
}

// 导出认证消息类型
export type AuthMessage = {
  token: string;
  deviceId: string;
  username?: string;
  email?: string;
  password?: string;
}

// 导出proto对象
export const proto = protoBundle.proto;

// 为方便使用，创建一些工具函数
export const Message = {
  create: (data: Partial<IMessage>) => protoBundle.proto.Message.create(data),
  encode: (message: any) => protoBundle.proto.Message.encode(message).finish(),
  decode: (buffer: Uint8Array) => protoBundle.proto.Message.decode(buffer)
};

export default protoBundle; 