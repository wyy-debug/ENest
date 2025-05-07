import * as $protobuf from 'protobufjs/minimal';

// Common aliases
const $Reader = $protobuf.Reader;
const $Writer = $protobuf.Writer;
const $util = $protobuf.util;

// 定义根命名空间
const $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

// 导入MessageType枚举
export const MessageType = {
  UNKNOWN: 0,
  HEARTBEAT: 1,
  AUTH: 2,
  CHAT: 3,
  SYSTEM: 4,
  STUDY_ROOM: 5,
  FRIEND: 6,
  PROFILE: 7,
  ERROR: 8,
  REGISTER: 9
};

// StudyRoomOperation枚举
export const StudyRoomOperation = {
  CREATE: 0,
  JOIN: 1,
  LEAVE: 2,
  DESTROY: 3,
  GET_DETAIL: 4
};

// 创建空的proto对象
const proto = {
  Message: {},
  HeartbeatMessage: {},
  AuthMessage: {},
  RegisterMessage: {},
  ChatMessage: {},
  SystemMessage: {},
  StudyRoomMessage: {},
  FriendMessage: {},
  ProfileMessage: {},
  ErrorMessage: {}
};

// 导出proto对象
export { proto };
export default { proto, MessageType, StudyRoomOperation }; 