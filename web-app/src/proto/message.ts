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


export interface AuthMessage {
  token: string;
  device_id?: string;
}

export interface StudyRoomMessage {
  operation: StudyRoomOperation;
  roomId?: number;
  name?: string;
  maxMembers?: number;
  isPrivate?: boolean;
  duration?: string;
  shareLink?: string;
}

export interface FriendMessage {
  operation: FriendOperation;
  friendId?: number;
  action?: string;
  contractType?: string;
  contractTerms?: string;
  contractId?: number;
}

export interface ProfileMessage {
  operation: ProfileOperation;
  username?: string;
  signature?: string;
  studyDirection?: string;
}

export interface ErrorMessage {
  code: number;
  message: string;
  detail?: string;
  requestId?: string;
}