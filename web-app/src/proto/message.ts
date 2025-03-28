export enum MessageType {
  UNKNOWN = 0,
  HEARTBEAT = 1,
  AUTH = 2,
  CHAT = 3,
  SYSTEM = 4,
  STUDY_ROOM = 5,
  FRIEND = 6,
  PROFILE = 7,
  ERROR = 8
}

export enum StudyRoomOperation {
  CREATE = 0,
  JOIN = 1,
  LEAVE = 2,
  DESTROY = 3,
  GET_DETAIL = 4
}

export enum FriendOperation {
  SEND_REQUEST = 0,
  HANDLE_REQUEST = 1,
  GET_LIST = 2,
  DELETE = 3,
  CREATE_CONTRACT = 4,
  TERMINATE_CONTRACT = 5,
  GET_CONTRACTS = 6
}

export enum ProfileOperation {
  UPDATE = 0,
  GET = 1
}

export interface Message {
  type: MessageType;
  timestamp: number;
  payload: Uint8Array;
  session_id: string;
}

export interface HeartbeatMessage {
  timestamp: number;
}

export interface AuthMessage {
  token: string;
  device_id?: string;
  username?: string;
  email?: string;
  password_hash?: string;
}

export interface ChatMessage {
  sender_id: number;
  receiver_id: number;
  content: string;
  message_type: string;
}

export interface SystemMessage {
  type: string;
  content: string;
}

export interface StudyRoomMessage {
  operation: StudyRoomOperation;
  room_id?: number;
  name?: string;
  max_members?: number;
  is_private?: boolean;
  duration?: string;
  share_link?: string;
}

export interface FriendMessage {
  operation: FriendOperation;
  friend_id?: number;
  action?: string;
  contract_type?: string;
  contract_terms?: string;
  contract_id?: number;
}

export interface ProfileMessage {
  operation: ProfileOperation;
  username?: string;
  signature?: string;
  study_direction?: string;
}

export interface ErrorMessage {
  code: number;
  message: string;
  detail?: string;
  request_id?: string;
}