import * as $protobuf from "protobufjs";
import Long = require("long");
/** Namespace proto. */
export namespace proto {

    /** MessageType enum. */
    enum MessageType {
        UNKNOWN = 0,
        HEARTBEAT = 1,
        AUTH = 2,
        CHAT = 3,
        SYSTEM = 4,
        STUDY_ROOM = 5,
        FRIEND = 6,
        PROFILE = 7,
        ERROR = 8,
        REGISTER = 9
    }

    /** Properties of a Message. */
    interface IMessage {

        /** Message type */
        type?: (proto.MessageType|null);

        /** Message timestamp */
        timestamp?: (number|Long|null);

        /** Message payload */
        payload?: (Uint8Array|null);

        /** Message sessionId */
        sessionId?: (string|null);
    }

    /** Represents a Message. */
    class Message implements IMessage {

        /**
         * Constructs a new Message.
         * @param [properties] Properties to set
         */
        constructor(properties?: proto.IMessage);

        /** Message type. */
        public type: proto.MessageType;

        /** Message timestamp. */
        public timestamp: (number|Long);

        /** Message payload. */
        public payload: Uint8Array;

        /** Message sessionId. */
        public sessionId: string;

        /**
         * Creates a new Message instance using the specified properties.
         * @param [properties] Properties to set
         * @returns Message instance
         */
        public static create(properties?: proto.IMessage): proto.Message;

        /**
         * Encodes the specified Message message. Does not implicitly {@link proto.Message.verify|verify} messages.
         * @param message Message message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: proto.IMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified Message message, length delimited. Does not implicitly {@link proto.Message.verify|verify} messages.
         * @param message Message message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: proto.IMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Message message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Message
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): proto.Message;

        /**
         * Decodes a Message message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns Message
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): proto.Message;

        /**
         * Verifies a Message message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a Message message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Message
         */
        public static fromObject(object: { [k: string]: any }): proto.Message;

        /**
         * Creates a plain object from a Message message. Also converts values to other types if specified.
         * @param message Message
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: proto.Message, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Message to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for Message
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a HeartbeatMessage. */
    interface IHeartbeatMessage {

        /** HeartbeatMessage timestamp */
        timestamp?: (number|Long|null);
    }

    /** Represents a HeartbeatMessage. */
    class HeartbeatMessage implements IHeartbeatMessage {

        /**
         * Constructs a new HeartbeatMessage.
         * @param [properties] Properties to set
         */
        constructor(properties?: proto.IHeartbeatMessage);

        /** HeartbeatMessage timestamp. */
        public timestamp: (number|Long);

        /**
         * Creates a new HeartbeatMessage instance using the specified properties.
         * @param [properties] Properties to set
         * @returns HeartbeatMessage instance
         */
        public static create(properties?: proto.IHeartbeatMessage): proto.HeartbeatMessage;

        /**
         * Encodes the specified HeartbeatMessage message. Does not implicitly {@link proto.HeartbeatMessage.verify|verify} messages.
         * @param message HeartbeatMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: proto.IHeartbeatMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified HeartbeatMessage message, length delimited. Does not implicitly {@link proto.HeartbeatMessage.verify|verify} messages.
         * @param message HeartbeatMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: proto.IHeartbeatMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a HeartbeatMessage message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns HeartbeatMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): proto.HeartbeatMessage;

        /**
         * Decodes a HeartbeatMessage message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns HeartbeatMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): proto.HeartbeatMessage;

        /**
         * Verifies a HeartbeatMessage message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a HeartbeatMessage message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns HeartbeatMessage
         */
        public static fromObject(object: { [k: string]: any }): proto.HeartbeatMessage;

        /**
         * Creates a plain object from a HeartbeatMessage message. Also converts values to other types if specified.
         * @param message HeartbeatMessage
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: proto.HeartbeatMessage, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this HeartbeatMessage to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for HeartbeatMessage
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of an AuthMessage. */
    interface IAuthMessage {

        /** AuthMessage token */
        token?: (string|null);

        /** AuthMessage deviceId */
        deviceId?: (string|null);

        /** AuthMessage username */
        username?: (string|null);

        /** AuthMessage email */
        email?: (string|null);

        /** AuthMessage password */
        password?: (string|null);
    }

    /** Represents an AuthMessage. */
    class AuthMessage implements IAuthMessage {

        /**
         * Constructs a new AuthMessage.
         * @param [properties] Properties to set
         */
        constructor(properties?: proto.IAuthMessage);

        /** AuthMessage token. */
        public token: string;

        /** AuthMessage deviceId. */
        public deviceId: string;

        /** AuthMessage username. */
        public username: string;

        /** AuthMessage email. */
        public email: string;

        /** AuthMessage password. */
        public password: string;

        /**
         * Creates a new AuthMessage instance using the specified properties.
         * @param [properties] Properties to set
         * @returns AuthMessage instance
         */
        public static create(properties?: proto.IAuthMessage): proto.AuthMessage;

        /**
         * Encodes the specified AuthMessage message. Does not implicitly {@link proto.AuthMessage.verify|verify} messages.
         * @param message AuthMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: proto.IAuthMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified AuthMessage message, length delimited. Does not implicitly {@link proto.AuthMessage.verify|verify} messages.
         * @param message AuthMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: proto.IAuthMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an AuthMessage message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns AuthMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): proto.AuthMessage;

        /**
         * Decodes an AuthMessage message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns AuthMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): proto.AuthMessage;

        /**
         * Verifies an AuthMessage message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates an AuthMessage message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns AuthMessage
         */
        public static fromObject(object: { [k: string]: any }): proto.AuthMessage;

        /**
         * Creates a plain object from an AuthMessage message. Also converts values to other types if specified.
         * @param message AuthMessage
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: proto.AuthMessage, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this AuthMessage to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for AuthMessage
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a RegisterMessage. */
    interface IRegisterMessage {

        /** RegisterMessage username */
        username?: (string|null);

        /** RegisterMessage password */
        password?: (string|null);

        /** RegisterMessage email */
        email?: (string|null);
    }

    /** Represents a RegisterMessage. */
    class RegisterMessage implements IRegisterMessage {

        /**
         * Constructs a new RegisterMessage.
         * @param [properties] Properties to set
         */
        constructor(properties?: proto.IRegisterMessage);

        /** RegisterMessage username. */
        public username: string;

        /** RegisterMessage password. */
        public password: string;

        /** RegisterMessage email. */
        public email: string;

        /**
         * Creates a new RegisterMessage instance using the specified properties.
         * @param [properties] Properties to set
         * @returns RegisterMessage instance
         */
        public static create(properties?: proto.IRegisterMessage): proto.RegisterMessage;

        /**
         * Encodes the specified RegisterMessage message. Does not implicitly {@link proto.RegisterMessage.verify|verify} messages.
         * @param message RegisterMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: proto.IRegisterMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified RegisterMessage message, length delimited. Does not implicitly {@link proto.RegisterMessage.verify|verify} messages.
         * @param message RegisterMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: proto.IRegisterMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a RegisterMessage message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns RegisterMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): proto.RegisterMessage;

        /**
         * Decodes a RegisterMessage message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns RegisterMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): proto.RegisterMessage;

        /**
         * Verifies a RegisterMessage message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a RegisterMessage message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns RegisterMessage
         */
        public static fromObject(object: { [k: string]: any }): proto.RegisterMessage;

        /**
         * Creates a plain object from a RegisterMessage message. Also converts values to other types if specified.
         * @param message RegisterMessage
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: proto.RegisterMessage, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this RegisterMessage to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for RegisterMessage
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a ChatMessage. */
    interface IChatMessage {

        /** ChatMessage senderId */
        senderId?: (number|null);

        /** ChatMessage receiverId */
        receiverId?: (number|null);

        /** ChatMessage content */
        content?: (string|null);

        /** ChatMessage messageType */
        messageType?: (string|null);
    }

    /** Represents a ChatMessage. */
    class ChatMessage implements IChatMessage {

        /**
         * Constructs a new ChatMessage.
         * @param [properties] Properties to set
         */
        constructor(properties?: proto.IChatMessage);

        /** ChatMessage senderId. */
        public senderId: number;

        /** ChatMessage receiverId. */
        public receiverId: number;

        /** ChatMessage content. */
        public content: string;

        /** ChatMessage messageType. */
        public messageType: string;

        /**
         * Creates a new ChatMessage instance using the specified properties.
         * @param [properties] Properties to set
         * @returns ChatMessage instance
         */
        public static create(properties?: proto.IChatMessage): proto.ChatMessage;

        /**
         * Encodes the specified ChatMessage message. Does not implicitly {@link proto.ChatMessage.verify|verify} messages.
         * @param message ChatMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: proto.IChatMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified ChatMessage message, length delimited. Does not implicitly {@link proto.ChatMessage.verify|verify} messages.
         * @param message ChatMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: proto.IChatMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ChatMessage message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ChatMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): proto.ChatMessage;

        /**
         * Decodes a ChatMessage message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns ChatMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): proto.ChatMessage;

        /**
         * Verifies a ChatMessage message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a ChatMessage message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns ChatMessage
         */
        public static fromObject(object: { [k: string]: any }): proto.ChatMessage;

        /**
         * Creates a plain object from a ChatMessage message. Also converts values to other types if specified.
         * @param message ChatMessage
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: proto.ChatMessage, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this ChatMessage to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for ChatMessage
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a SystemMessage. */
    interface ISystemMessage {

        /** SystemMessage type */
        type?: (string|null);

        /** SystemMessage content */
        content?: (string|null);
    }

    /** Represents a SystemMessage. */
    class SystemMessage implements ISystemMessage {

        /**
         * Constructs a new SystemMessage.
         * @param [properties] Properties to set
         */
        constructor(properties?: proto.ISystemMessage);

        /** SystemMessage type. */
        public type: string;

        /** SystemMessage content. */
        public content: string;

        /**
         * Creates a new SystemMessage instance using the specified properties.
         * @param [properties] Properties to set
         * @returns SystemMessage instance
         */
        public static create(properties?: proto.ISystemMessage): proto.SystemMessage;

        /**
         * Encodes the specified SystemMessage message. Does not implicitly {@link proto.SystemMessage.verify|verify} messages.
         * @param message SystemMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: proto.ISystemMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified SystemMessage message, length delimited. Does not implicitly {@link proto.SystemMessage.verify|verify} messages.
         * @param message SystemMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: proto.ISystemMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a SystemMessage message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns SystemMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): proto.SystemMessage;

        /**
         * Decodes a SystemMessage message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns SystemMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): proto.SystemMessage;

        /**
         * Verifies a SystemMessage message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a SystemMessage message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns SystemMessage
         */
        public static fromObject(object: { [k: string]: any }): proto.SystemMessage;

        /**
         * Creates a plain object from a SystemMessage message. Also converts values to other types if specified.
         * @param message SystemMessage
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: proto.SystemMessage, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this SystemMessage to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for SystemMessage
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a StudyRoomMessage. */
    interface IStudyRoomMessage {

        /** StudyRoomMessage operation */
        operation?: (proto.StudyRoomMessage.Operation|null);

        /** StudyRoomMessage roomId */
        roomId?: (number|null);

        /** StudyRoomMessage name */
        name?: (string|null);

        /** StudyRoomMessage maxMembers */
        maxMembers?: (number|null);

        /** StudyRoomMessage isPrivate */
        isPrivate?: (boolean|null);

        /** StudyRoomMessage duration */
        duration?: (string|null);

        /** StudyRoomMessage shareLink */
        shareLink?: (string|null);
    }

    /** Represents a StudyRoomMessage. */
    class StudyRoomMessage implements IStudyRoomMessage {

        /**
         * Constructs a new StudyRoomMessage.
         * @param [properties] Properties to set
         */
        constructor(properties?: proto.IStudyRoomMessage);

        /** StudyRoomMessage operation. */
        public operation: proto.StudyRoomMessage.Operation;

        /** StudyRoomMessage roomId. */
        public roomId: number;

        /** StudyRoomMessage name. */
        public name: string;

        /** StudyRoomMessage maxMembers. */
        public maxMembers: number;

        /** StudyRoomMessage isPrivate. */
        public isPrivate: boolean;

        /** StudyRoomMessage duration. */
        public duration: string;

        /** StudyRoomMessage shareLink. */
        public shareLink: string;

        /**
         * Creates a new StudyRoomMessage instance using the specified properties.
         * @param [properties] Properties to set
         * @returns StudyRoomMessage instance
         */
        public static create(properties?: proto.IStudyRoomMessage): proto.StudyRoomMessage;

        /**
         * Encodes the specified StudyRoomMessage message. Does not implicitly {@link proto.StudyRoomMessage.verify|verify} messages.
         * @param message StudyRoomMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: proto.IStudyRoomMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified StudyRoomMessage message, length delimited. Does not implicitly {@link proto.StudyRoomMessage.verify|verify} messages.
         * @param message StudyRoomMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: proto.IStudyRoomMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a StudyRoomMessage message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns StudyRoomMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): proto.StudyRoomMessage;

        /**
         * Decodes a StudyRoomMessage message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns StudyRoomMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): proto.StudyRoomMessage;

        /**
         * Verifies a StudyRoomMessage message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a StudyRoomMessage message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns StudyRoomMessage
         */
        public static fromObject(object: { [k: string]: any }): proto.StudyRoomMessage;

        /**
         * Creates a plain object from a StudyRoomMessage message. Also converts values to other types if specified.
         * @param message StudyRoomMessage
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: proto.StudyRoomMessage, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this StudyRoomMessage to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for StudyRoomMessage
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    namespace StudyRoomMessage {

        /** Operation enum. */
        enum Operation {
            CREATE = 0,
            JOIN = 1,
            LEAVE = 2,
            DESTROY = 3,
            GET_DETAIL = 4
        }
    }

    /** Properties of a FriendMessage. */
    interface IFriendMessage {

        /** FriendMessage operation */
        operation?: (proto.FriendMessage.Operation|null);

        /** FriendMessage friendId */
        friendId?: (number|null);

        /** FriendMessage action */
        action?: (string|null);

        /** FriendMessage contractType */
        contractType?: (string|null);

        /** FriendMessage contractTerms */
        contractTerms?: (string|null);

        /** FriendMessage contractId */
        contractId?: (number|null);
    }

    /** Represents a FriendMessage. */
    class FriendMessage implements IFriendMessage {

        /**
         * Constructs a new FriendMessage.
         * @param [properties] Properties to set
         */
        constructor(properties?: proto.IFriendMessage);

        /** FriendMessage operation. */
        public operation: proto.FriendMessage.Operation;

        /** FriendMessage friendId. */
        public friendId: number;

        /** FriendMessage action. */
        public action: string;

        /** FriendMessage contractType. */
        public contractType: string;

        /** FriendMessage contractTerms. */
        public contractTerms: string;

        /** FriendMessage contractId. */
        public contractId: number;

        /**
         * Creates a new FriendMessage instance using the specified properties.
         * @param [properties] Properties to set
         * @returns FriendMessage instance
         */
        public static create(properties?: proto.IFriendMessage): proto.FriendMessage;

        /**
         * Encodes the specified FriendMessage message. Does not implicitly {@link proto.FriendMessage.verify|verify} messages.
         * @param message FriendMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: proto.IFriendMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified FriendMessage message, length delimited. Does not implicitly {@link proto.FriendMessage.verify|verify} messages.
         * @param message FriendMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: proto.IFriendMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a FriendMessage message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns FriendMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): proto.FriendMessage;

        /**
         * Decodes a FriendMessage message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns FriendMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): proto.FriendMessage;

        /**
         * Verifies a FriendMessage message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a FriendMessage message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns FriendMessage
         */
        public static fromObject(object: { [k: string]: any }): proto.FriendMessage;

        /**
         * Creates a plain object from a FriendMessage message. Also converts values to other types if specified.
         * @param message FriendMessage
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: proto.FriendMessage, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this FriendMessage to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for FriendMessage
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    namespace FriendMessage {

        /** Operation enum. */
        enum Operation {
            SEND_REQUEST = 0,
            HANDLE_REQUEST = 1,
            GET_LIST = 2,
            DELETE = 3,
            CREATE_CONTRACT = 4,
            TERMINATE_CONTRACT = 5,
            GET_CONTRACTS = 6
        }
    }

    /** Properties of a ProfileMessage. */
    interface IProfileMessage {

        /** ProfileMessage operation */
        operation?: (proto.ProfileMessage.Operation|null);

        /** ProfileMessage username */
        username?: (string|null);

        /** ProfileMessage signature */
        signature?: (string|null);

        /** ProfileMessage studyDirection */
        studyDirection?: (string|null);
    }

    /** Represents a ProfileMessage. */
    class ProfileMessage implements IProfileMessage {

        /**
         * Constructs a new ProfileMessage.
         * @param [properties] Properties to set
         */
        constructor(properties?: proto.IProfileMessage);

        /** ProfileMessage operation. */
        public operation: proto.ProfileMessage.Operation;

        /** ProfileMessage username. */
        public username: string;

        /** ProfileMessage signature. */
        public signature: string;

        /** ProfileMessage studyDirection. */
        public studyDirection: string;

        /**
         * Creates a new ProfileMessage instance using the specified properties.
         * @param [properties] Properties to set
         * @returns ProfileMessage instance
         */
        public static create(properties?: proto.IProfileMessage): proto.ProfileMessage;

        /**
         * Encodes the specified ProfileMessage message. Does not implicitly {@link proto.ProfileMessage.verify|verify} messages.
         * @param message ProfileMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: proto.IProfileMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified ProfileMessage message, length delimited. Does not implicitly {@link proto.ProfileMessage.verify|verify} messages.
         * @param message ProfileMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: proto.IProfileMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ProfileMessage message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ProfileMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): proto.ProfileMessage;

        /**
         * Decodes a ProfileMessage message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns ProfileMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): proto.ProfileMessage;

        /**
         * Verifies a ProfileMessage message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a ProfileMessage message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns ProfileMessage
         */
        public static fromObject(object: { [k: string]: any }): proto.ProfileMessage;

        /**
         * Creates a plain object from a ProfileMessage message. Also converts values to other types if specified.
         * @param message ProfileMessage
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: proto.ProfileMessage, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this ProfileMessage to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for ProfileMessage
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    namespace ProfileMessage {

        /** Operation enum. */
        enum Operation {
            UPDATE = 0,
            GET = 1
        }
    }

    /** Properties of an ErrorMessage. */
    interface IErrorMessage {

        /** ErrorMessage code */
        code?: (number|null);

        /** ErrorMessage message */
        message?: (string|null);

        /** ErrorMessage detail */
        detail?: (string|null);

        /** ErrorMessage requestId */
        requestId?: (string|null);
    }

    /** Represents an ErrorMessage. */
    class ErrorMessage implements IErrorMessage {

        /**
         * Constructs a new ErrorMessage.
         * @param [properties] Properties to set
         */
        constructor(properties?: proto.IErrorMessage);

        /** ErrorMessage code. */
        public code: number;

        /** ErrorMessage message. */
        public message: string;

        /** ErrorMessage detail. */
        public detail: string;

        /** ErrorMessage requestId. */
        public requestId: string;

        /**
         * Creates a new ErrorMessage instance using the specified properties.
         * @param [properties] Properties to set
         * @returns ErrorMessage instance
         */
        public static create(properties?: proto.IErrorMessage): proto.ErrorMessage;

        /**
         * Encodes the specified ErrorMessage message. Does not implicitly {@link proto.ErrorMessage.verify|verify} messages.
         * @param message ErrorMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: proto.IErrorMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified ErrorMessage message, length delimited. Does not implicitly {@link proto.ErrorMessage.verify|verify} messages.
         * @param message ErrorMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: proto.IErrorMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an ErrorMessage message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ErrorMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): proto.ErrorMessage;

        /**
         * Decodes an ErrorMessage message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns ErrorMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): proto.ErrorMessage;

        /**
         * Verifies an ErrorMessage message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates an ErrorMessage message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns ErrorMessage
         */
        public static fromObject(object: { [k: string]: any }): proto.ErrorMessage;

        /**
         * Creates a plain object from an ErrorMessage message. Also converts values to other types if specified.
         * @param message ErrorMessage
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: proto.ErrorMessage, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this ErrorMessage to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for ErrorMessage
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }
}
