declare module '../proto/index' {
  export default {
    proto: {
      Message: {
        encode(message: any): { finish(): Uint8Array };
        decode(bytes: Uint8Array): any;
      },
      AuthMessage: {
        encode(message: any): { finish(): Uint8Array };
        decode(bytes: Uint8Array): any;
      },
      HeartbeatMessage: {
        encode(message: any): { finish(): Uint8Array };
        decode(bytes: Uint8Array): any;
      },
      RegisterMessage: {
        encode(message: any): { finish(): Uint8Array };
        decode(bytes: Uint8Array): any;
      },
      ChatMessage: {
        encode(message: any): { finish(): Uint8Array };
        decode(bytes: Uint8Array): any;
      },
      SystemMessage: {
        encode(message: any): { finish(): Uint8Array };
        decode(bytes: Uint8Array): any;
      },
      StudyRoomMessage: {
        encode(message: any): { finish(): Uint8Array };
        decode(bytes: Uint8Array): any;
      },
      FriendMessage: {
        encode(message: any): { finish(): Uint8Array };
        decode(bytes: Uint8Array): any;
      },
      ProfileMessage: {
        encode(message: any): { finish(): Uint8Array };
        decode(bytes: Uint8Array): any;
      },
      ErrorMessage: {
        encode(message: any): { finish(): Uint8Array };
        decode(bytes: Uint8Array): any;
      }
    }
  };
} 