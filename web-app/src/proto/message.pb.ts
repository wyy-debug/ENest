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
            },
            username: {
              type: 'string',
              id: 3
            },
            email: {
              type: 'string',
              id: 4
            },
            password_hash: {
              type: 'string',
              id: 5
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
        },
        StudyRoomMessage: {
          nested: {
            Operation: {
              values: {
                CREATE: 0,
                JOIN: 1,
                LEAVE: 2,
                DESTROY: 3,
                GET_DETAIL: 4
              }
            }
          },
          fields: {
            operation: {
              type: 'Operation',
              id: 1
            },
            room_id: {
              type: 'int32',
              id: 2
            },
            name: {
              type: 'string',
              id: 3
            },
            max_members: {
              type: 'int32',
              id: 4
            },
            is_private: {
              type: 'bool',
              id: 5
            },
            duration: {
              type: 'string',
              id: 6
            },
            share_link: {
              type: 'string',
              id: 7
            }
          }
        },
        FriendMessage: {
          nested: {
            Operation: {
              values: {
                SEND_REQUEST: 0,
                HANDLE_REQUEST: 1,
                GET_LIST: 2,
                DELETE: 3,
                CREATE_CONTRACT: 4,
                TERMINATE_CONTRACT: 5,
                GET_CONTRACTS: 6
              }
            }
          },
          fields: {
            operation: {
              type: 'Operation',
              id: 1
            },
            friend_id: {
              type: 'int32',
              id: 2
            },
            action: {
              type: 'string',
              id: 3
            },
            contract_type: {
              type: 'string',
              id: 4
            },
            contract_terms: {
              type: 'string',
              id: 5
            },
            contract_id: {
              type: 'int32',
              id: 6
            }
          }
        },
        ProfileMessage: {
          nested: {
            Operation: {
              values: {
                UPDATE: 0,
                GET: 1
              }
            }
          },
          fields: {
            operation: {
              type: 'Operation',
              id: 1
            },
            username: {
              type: 'string',
              id: 2
            },
            signature: {
              type: 'string',
              id: 3
            },
            study_direction: {
              type: 'string',
              id: 4
            }
          }
        },
        ErrorMessage: {
          fields: {
            code: {
              type: 'int32',
              id: 1
            },
            message: {
              type: 'string',
              id: 2
            },
            detail: {
              type: 'string',
              id: 3
            },
            request_id: {
              type: 'string',
              id: 4
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
export const StudyRoomMessage = root.lookupType('proto.StudyRoomMessage');
export const FriendMessage = root.lookupType('proto.FriendMessage');
export const ProfileMessage = root.lookupType('proto.ProfileMessage');
export const ErrorMessage = root.lookupType('proto.ErrorMessage');
export { MessageType } from './message';