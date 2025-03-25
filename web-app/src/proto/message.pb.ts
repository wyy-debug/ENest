import * as protobuf from 'protobufjs';

const root = protobuf.Root.fromJSON({
  nested: {
    proto: {
      nested: {
        MessageType: {
          values: {
            UNKNOWN: 0,
            HEARTBEAT: 1,
            AUTH: 2,
            CHAT: 3,
            SYSTEM: 4,
            STUDY_ROOM: 5,
            FRIEND: 6,
            PROFILE: 7,
            ERROR: 8
          }
        },
        Message: {
          fields: {
            type: {
              type: 'MessageType',
              id: 1
            },
            timestamp: {
              type: 'int64',
              id: 2
            },
            payload: {
              type: 'bytes',
              id: 3
            },
            session_id: {
              type: 'string',
              id: 4
            }
          }
        },
        HeartbeatMessage: {
          fields: {
            timestamp: {
              type: 'int64',
              id: 1
            }
          }
        },
        AuthMessage: {
          fields: {
            token: {
              type: 'string',
              id: 1
            },
            device_id: {
              type: 'string',
              id: 2
            }
          }
        },
        ChatMessage: {
          fields: {
            sender_id: {
              type: 'int32',
              id: 1
            },
            receiver_id: {
              type: 'int32',
              id: 2
            },
            content: {
              type: 'string',
              id: 3
            },
            message_type: {
              type: 'string',
              id: 4
            }
          }
        },
        SystemMessage: {
          fields: {
            type: {
              type: 'string',
              id: 1
            },
            content: {
              type: 'string',
              id: 2
            }
          }
        }
      }
    }
  }
});

export const Message = root.lookupType('proto.Message');
export const HeartbeatMessage = root.lookupType('proto.HeartbeatMessage');
export const AuthMessage = root.lookupType('proto.AuthMessage');
export const ChatMessage = root.lookupType('proto.ChatMessage');
export const SystemMessage = root.lookupType('proto.SystemMessage');
export { MessageType } from './message';