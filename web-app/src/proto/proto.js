/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.proto = (function() {

    /**
     * Namespace proto.
     * @exports proto
     * @namespace
     */
    var proto = {};

    /**
     * MessageType enum.
     * @name proto.MessageType
     * @enum {number}
     * @property {number} UNKNOWN=0 UNKNOWN value
     * @property {number} HEARTBEAT=1 HEARTBEAT value
     * @property {number} AUTH=2 AUTH value
     * @property {number} CHAT=3 CHAT value
     * @property {number} SYSTEM=4 SYSTEM value
     * @property {number} STUDY_ROOM=5 STUDY_ROOM value
     * @property {number} FRIEND=6 FRIEND value
     * @property {number} PROFILE=7 PROFILE value
     * @property {number} ERROR=8 ERROR value
     * @property {number} REGISTER=9 REGISTER value
     */
    proto.MessageType = (function() {
        var valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "UNKNOWN"] = 0;
        values[valuesById[1] = "HEARTBEAT"] = 1;
        values[valuesById[2] = "AUTH"] = 2;
        values[valuesById[3] = "CHAT"] = 3;
        values[valuesById[4] = "SYSTEM"] = 4;
        values[valuesById[5] = "STUDY_ROOM"] = 5;
        values[valuesById[6] = "FRIEND"] = 6;
        values[valuesById[7] = "PROFILE"] = 7;
        values[valuesById[8] = "ERROR"] = 8;
        values[valuesById[9] = "REGISTER"] = 9;
        return values;
    })();

    proto.Message = (function() {

        /**
         * Properties of a Message.
         * @memberof proto
         * @interface IMessage
         * @property {proto.MessageType|null} [type] Message type
         * @property {number|Long|null} [timestamp] Message timestamp
         * @property {Uint8Array|null} [payload] Message payload
         * @property {string|null} [sessionId] Message sessionId
         */

        /**
         * Constructs a new Message.
         * @memberof proto
         * @classdesc Represents a Message.
         * @implements IMessage
         * @constructor
         * @param {proto.IMessage=} [properties] Properties to set
         */
        function Message(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Message type.
         * @member {proto.MessageType} type
         * @memberof proto.Message
         * @instance
         */
        Message.prototype.type = 0;

        /**
         * Message timestamp.
         * @member {number|Long} timestamp
         * @memberof proto.Message
         * @instance
         */
        Message.prototype.timestamp = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Message payload.
         * @member {Uint8Array} payload
         * @memberof proto.Message
         * @instance
         */
        Message.prototype.payload = $util.newBuffer([]);

        /**
         * Message sessionId.
         * @member {string} sessionId
         * @memberof proto.Message
         * @instance
         */
        Message.prototype.sessionId = "";

        /**
         * Creates a new Message instance using the specified properties.
         * @function create
         * @memberof proto.Message
         * @static
         * @param {proto.IMessage=} [properties] Properties to set
         * @returns {proto.Message} Message instance
         */
        Message.create = function create(properties) {
            return new Message(properties);
        };

        /**
         * Encodes the specified Message message. Does not implicitly {@link proto.Message.verify|verify} messages.
         * @function encode
         * @memberof proto.Message
         * @static
         * @param {proto.IMessage} message Message message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Message.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.type != null && Object.hasOwnProperty.call(message, "type"))
                writer.uint32(/* id 1, wireType 0 =*/8).int32(message.type);
            if (message.timestamp != null && Object.hasOwnProperty.call(message, "timestamp"))
                writer.uint32(/* id 2, wireType 0 =*/16).int64(message.timestamp);
            if (message.payload != null && Object.hasOwnProperty.call(message, "payload"))
                writer.uint32(/* id 3, wireType 2 =*/26).bytes(message.payload);
            if (message.sessionId != null && Object.hasOwnProperty.call(message, "sessionId"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.sessionId);
            return writer;
        };

        /**
         * Encodes the specified Message message, length delimited. Does not implicitly {@link proto.Message.verify|verify} messages.
         * @function encodeDelimited
         * @memberof proto.Message
         * @static
         * @param {proto.IMessage} message Message message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Message.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Message message from the specified reader or buffer.
         * @function decode
         * @memberof proto.Message
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {proto.Message} Message
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Message.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.proto.Message();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.type = reader.int32();
                        break;
                    }
                case 2: {
                        message.timestamp = reader.int64();
                        break;
                    }
                case 3: {
                        message.payload = reader.bytes();
                        break;
                    }
                case 4: {
                        message.sessionId = reader.string();
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Message message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof proto.Message
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {proto.Message} Message
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Message.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Message message.
         * @function verify
         * @memberof proto.Message
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Message.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.type != null && message.hasOwnProperty("type"))
                switch (message.type) {
                default:
                    return "type: enum value expected";
                case 0:
                case 1:
                case 2:
                case 3:
                case 4:
                case 5:
                case 6:
                case 7:
                case 8:
                case 9:
                    break;
                }
            if (message.timestamp != null && message.hasOwnProperty("timestamp"))
                if (!$util.isInteger(message.timestamp) && !(message.timestamp && $util.isInteger(message.timestamp.low) && $util.isInteger(message.timestamp.high)))
                    return "timestamp: integer|Long expected";
            if (message.payload != null && message.hasOwnProperty("payload"))
                if (!(message.payload && typeof message.payload.length === "number" || $util.isString(message.payload)))
                    return "payload: buffer expected";
            if (message.sessionId != null && message.hasOwnProperty("sessionId"))
                if (!$util.isString(message.sessionId))
                    return "sessionId: string expected";
            return null;
        };

        /**
         * Creates a Message message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof proto.Message
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {proto.Message} Message
         */
        Message.fromObject = function fromObject(object) {
            if (object instanceof $root.proto.Message)
                return object;
            var message = new $root.proto.Message();
            switch (object.type) {
            default:
                if (typeof object.type === "number") {
                    message.type = object.type;
                    break;
                }
                break;
            case "UNKNOWN":
            case 0:
                message.type = 0;
                break;
            case "HEARTBEAT":
            case 1:
                message.type = 1;
                break;
            case "AUTH":
            case 2:
                message.type = 2;
                break;
            case "CHAT":
            case 3:
                message.type = 3;
                break;
            case "SYSTEM":
            case 4:
                message.type = 4;
                break;
            case "STUDY_ROOM":
            case 5:
                message.type = 5;
                break;
            case "FRIEND":
            case 6:
                message.type = 6;
                break;
            case "PROFILE":
            case 7:
                message.type = 7;
                break;
            case "ERROR":
            case 8:
                message.type = 8;
                break;
            case "REGISTER":
            case 9:
                message.type = 9;
                break;
            }
            if (object.timestamp != null)
                if ($util.Long)
                    (message.timestamp = $util.Long.fromValue(object.timestamp)).unsigned = false;
                else if (typeof object.timestamp === "string")
                    message.timestamp = parseInt(object.timestamp, 10);
                else if (typeof object.timestamp === "number")
                    message.timestamp = object.timestamp;
                else if (typeof object.timestamp === "object")
                    message.timestamp = new $util.LongBits(object.timestamp.low >>> 0, object.timestamp.high >>> 0).toNumber();
            if (object.payload != null)
                if (typeof object.payload === "string")
                    $util.base64.decode(object.payload, message.payload = $util.newBuffer($util.base64.length(object.payload)), 0);
                else if (object.payload.length >= 0)
                    message.payload = object.payload;
            if (object.sessionId != null)
                message.sessionId = String(object.sessionId);
            return message;
        };

        /**
         * Creates a plain object from a Message message. Also converts values to other types if specified.
         * @function toObject
         * @memberof proto.Message
         * @static
         * @param {proto.Message} message Message
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Message.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.type = options.enums === String ? "UNKNOWN" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.timestamp = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.timestamp = options.longs === String ? "0" : 0;
                if (options.bytes === String)
                    object.payload = "";
                else {
                    object.payload = [];
                    if (options.bytes !== Array)
                        object.payload = $util.newBuffer(object.payload);
                }
                object.sessionId = "";
            }
            if (message.type != null && message.hasOwnProperty("type"))
                object.type = options.enums === String ? $root.proto.MessageType[message.type] === undefined ? message.type : $root.proto.MessageType[message.type] : message.type;
            if (message.timestamp != null && message.hasOwnProperty("timestamp"))
                if (typeof message.timestamp === "number")
                    object.timestamp = options.longs === String ? String(message.timestamp) : message.timestamp;
                else
                    object.timestamp = options.longs === String ? $util.Long.prototype.toString.call(message.timestamp) : options.longs === Number ? new $util.LongBits(message.timestamp.low >>> 0, message.timestamp.high >>> 0).toNumber() : message.timestamp;
            if (message.payload != null && message.hasOwnProperty("payload"))
                object.payload = options.bytes === String ? $util.base64.encode(message.payload, 0, message.payload.length) : options.bytes === Array ? Array.prototype.slice.call(message.payload) : message.payload;
            if (message.sessionId != null && message.hasOwnProperty("sessionId"))
                object.sessionId = message.sessionId;
            return object;
        };

        /**
         * Converts this Message to JSON.
         * @function toJSON
         * @memberof proto.Message
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Message.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Message
         * @function getTypeUrl
         * @memberof proto.Message
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Message.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/proto.Message";
        };

        return Message;
    })();

    proto.HeartbeatMessage = (function() {

        /**
         * Properties of a HeartbeatMessage.
         * @memberof proto
         * @interface IHeartbeatMessage
         * @property {number|Long|null} [timestamp] HeartbeatMessage timestamp
         */

        /**
         * Constructs a new HeartbeatMessage.
         * @memberof proto
         * @classdesc Represents a HeartbeatMessage.
         * @implements IHeartbeatMessage
         * @constructor
         * @param {proto.IHeartbeatMessage=} [properties] Properties to set
         */
        function HeartbeatMessage(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * HeartbeatMessage timestamp.
         * @member {number|Long} timestamp
         * @memberof proto.HeartbeatMessage
         * @instance
         */
        HeartbeatMessage.prototype.timestamp = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Creates a new HeartbeatMessage instance using the specified properties.
         * @function create
         * @memberof proto.HeartbeatMessage
         * @static
         * @param {proto.IHeartbeatMessage=} [properties] Properties to set
         * @returns {proto.HeartbeatMessage} HeartbeatMessage instance
         */
        HeartbeatMessage.create = function create(properties) {
            return new HeartbeatMessage(properties);
        };

        /**
         * Encodes the specified HeartbeatMessage message. Does not implicitly {@link proto.HeartbeatMessage.verify|verify} messages.
         * @function encode
         * @memberof proto.HeartbeatMessage
         * @static
         * @param {proto.IHeartbeatMessage} message HeartbeatMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        HeartbeatMessage.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.timestamp != null && Object.hasOwnProperty.call(message, "timestamp"))
                writer.uint32(/* id 1, wireType 0 =*/8).int64(message.timestamp);
            return writer;
        };

        /**
         * Encodes the specified HeartbeatMessage message, length delimited. Does not implicitly {@link proto.HeartbeatMessage.verify|verify} messages.
         * @function encodeDelimited
         * @memberof proto.HeartbeatMessage
         * @static
         * @param {proto.IHeartbeatMessage} message HeartbeatMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        HeartbeatMessage.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a HeartbeatMessage message from the specified reader or buffer.
         * @function decode
         * @memberof proto.HeartbeatMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {proto.HeartbeatMessage} HeartbeatMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        HeartbeatMessage.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.proto.HeartbeatMessage();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.timestamp = reader.int64();
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a HeartbeatMessage message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof proto.HeartbeatMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {proto.HeartbeatMessage} HeartbeatMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        HeartbeatMessage.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a HeartbeatMessage message.
         * @function verify
         * @memberof proto.HeartbeatMessage
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        HeartbeatMessage.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.timestamp != null && message.hasOwnProperty("timestamp"))
                if (!$util.isInteger(message.timestamp) && !(message.timestamp && $util.isInteger(message.timestamp.low) && $util.isInteger(message.timestamp.high)))
                    return "timestamp: integer|Long expected";
            return null;
        };

        /**
         * Creates a HeartbeatMessage message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof proto.HeartbeatMessage
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {proto.HeartbeatMessage} HeartbeatMessage
         */
        HeartbeatMessage.fromObject = function fromObject(object) {
            if (object instanceof $root.proto.HeartbeatMessage)
                return object;
            var message = new $root.proto.HeartbeatMessage();
            if (object.timestamp != null)
                if ($util.Long)
                    (message.timestamp = $util.Long.fromValue(object.timestamp)).unsigned = false;
                else if (typeof object.timestamp === "string")
                    message.timestamp = parseInt(object.timestamp, 10);
                else if (typeof object.timestamp === "number")
                    message.timestamp = object.timestamp;
                else if (typeof object.timestamp === "object")
                    message.timestamp = new $util.LongBits(object.timestamp.low >>> 0, object.timestamp.high >>> 0).toNumber();
            return message;
        };

        /**
         * Creates a plain object from a HeartbeatMessage message. Also converts values to other types if specified.
         * @function toObject
         * @memberof proto.HeartbeatMessage
         * @static
         * @param {proto.HeartbeatMessage} message HeartbeatMessage
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        HeartbeatMessage.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.timestamp = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.timestamp = options.longs === String ? "0" : 0;
            if (message.timestamp != null && message.hasOwnProperty("timestamp"))
                if (typeof message.timestamp === "number")
                    object.timestamp = options.longs === String ? String(message.timestamp) : message.timestamp;
                else
                    object.timestamp = options.longs === String ? $util.Long.prototype.toString.call(message.timestamp) : options.longs === Number ? new $util.LongBits(message.timestamp.low >>> 0, message.timestamp.high >>> 0).toNumber() : message.timestamp;
            return object;
        };

        /**
         * Converts this HeartbeatMessage to JSON.
         * @function toJSON
         * @memberof proto.HeartbeatMessage
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        HeartbeatMessage.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for HeartbeatMessage
         * @function getTypeUrl
         * @memberof proto.HeartbeatMessage
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        HeartbeatMessage.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/proto.HeartbeatMessage";
        };

        return HeartbeatMessage;
    })();

    proto.AuthMessage = (function() {

        /**
         * Properties of an AuthMessage.
         * @memberof proto
         * @interface IAuthMessage
         * @property {string|null} [token] AuthMessage token
         * @property {string|null} [deviceId] AuthMessage deviceId
         * @property {string|null} [username] AuthMessage username
         * @property {string|null} [email] AuthMessage email
         * @property {string|null} [password] AuthMessage password
         */

        /**
         * Constructs a new AuthMessage.
         * @memberof proto
         * @classdesc Represents an AuthMessage.
         * @implements IAuthMessage
         * @constructor
         * @param {proto.IAuthMessage=} [properties] Properties to set
         */
        function AuthMessage(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * AuthMessage token.
         * @member {string} token
         * @memberof proto.AuthMessage
         * @instance
         */
        AuthMessage.prototype.token = "";

        /**
         * AuthMessage deviceId.
         * @member {string} deviceId
         * @memberof proto.AuthMessage
         * @instance
         */
        AuthMessage.prototype.deviceId = "";

        /**
         * AuthMessage username.
         * @member {string} username
         * @memberof proto.AuthMessage
         * @instance
         */
        AuthMessage.prototype.username = "";

        /**
         * AuthMessage email.
         * @member {string} email
         * @memberof proto.AuthMessage
         * @instance
         */
        AuthMessage.prototype.email = "";

        /**
         * AuthMessage password.
         * @member {string} password
         * @memberof proto.AuthMessage
         * @instance
         */
        AuthMessage.prototype.password = "";

        /**
         * Creates a new AuthMessage instance using the specified properties.
         * @function create
         * @memberof proto.AuthMessage
         * @static
         * @param {proto.IAuthMessage=} [properties] Properties to set
         * @returns {proto.AuthMessage} AuthMessage instance
         */
        AuthMessage.create = function create(properties) {
            return new AuthMessage(properties);
        };

        /**
         * Encodes the specified AuthMessage message. Does not implicitly {@link proto.AuthMessage.verify|verify} messages.
         * @function encode
         * @memberof proto.AuthMessage
         * @static
         * @param {proto.IAuthMessage} message AuthMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        AuthMessage.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.token != null && Object.hasOwnProperty.call(message, "token"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.token);
            if (message.deviceId != null && Object.hasOwnProperty.call(message, "deviceId"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.deviceId);
            if (message.username != null && Object.hasOwnProperty.call(message, "username"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.username);
            if (message.email != null && Object.hasOwnProperty.call(message, "email"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.email);
            if (message.password != null && Object.hasOwnProperty.call(message, "password"))
                writer.uint32(/* id 5, wireType 2 =*/42).string(message.password);
            return writer;
        };

        /**
         * Encodes the specified AuthMessage message, length delimited. Does not implicitly {@link proto.AuthMessage.verify|verify} messages.
         * @function encodeDelimited
         * @memberof proto.AuthMessage
         * @static
         * @param {proto.IAuthMessage} message AuthMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        AuthMessage.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an AuthMessage message from the specified reader or buffer.
         * @function decode
         * @memberof proto.AuthMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {proto.AuthMessage} AuthMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        AuthMessage.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.proto.AuthMessage();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.token = reader.string();
                        break;
                    }
                case 2: {
                        message.deviceId = reader.string();
                        break;
                    }
                case 3: {
                        message.username = reader.string();
                        break;
                    }
                case 4: {
                        message.email = reader.string();
                        break;
                    }
                case 5: {
                        message.password = reader.string();
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an AuthMessage message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof proto.AuthMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {proto.AuthMessage} AuthMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        AuthMessage.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an AuthMessage message.
         * @function verify
         * @memberof proto.AuthMessage
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        AuthMessage.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.token != null && message.hasOwnProperty("token"))
                if (!$util.isString(message.token))
                    return "token: string expected";
            if (message.deviceId != null && message.hasOwnProperty("deviceId"))
                if (!$util.isString(message.deviceId))
                    return "deviceId: string expected";
            if (message.username != null && message.hasOwnProperty("username"))
                if (!$util.isString(message.username))
                    return "username: string expected";
            if (message.email != null && message.hasOwnProperty("email"))
                if (!$util.isString(message.email))
                    return "email: string expected";
            if (message.password != null && message.hasOwnProperty("password"))
                if (!$util.isString(message.password))
                    return "password: string expected";
            return null;
        };

        /**
         * Creates an AuthMessage message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof proto.AuthMessage
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {proto.AuthMessage} AuthMessage
         */
        AuthMessage.fromObject = function fromObject(object) {
            if (object instanceof $root.proto.AuthMessage)
                return object;
            var message = new $root.proto.AuthMessage();
            if (object.token != null)
                message.token = String(object.token);
            if (object.deviceId != null)
                message.deviceId = String(object.deviceId);
            if (object.username != null)
                message.username = String(object.username);
            if (object.email != null)
                message.email = String(object.email);
            if (object.password != null)
                message.password = String(object.password);
            return message;
        };

        /**
         * Creates a plain object from an AuthMessage message. Also converts values to other types if specified.
         * @function toObject
         * @memberof proto.AuthMessage
         * @static
         * @param {proto.AuthMessage} message AuthMessage
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        AuthMessage.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.token = "";
                object.deviceId = "";
                object.username = "";
                object.email = "";
                object.password = "";
            }
            if (message.token != null && message.hasOwnProperty("token"))
                object.token = message.token;
            if (message.deviceId != null && message.hasOwnProperty("deviceId"))
                object.deviceId = message.deviceId;
            if (message.username != null && message.hasOwnProperty("username"))
                object.username = message.username;
            if (message.email != null && message.hasOwnProperty("email"))
                object.email = message.email;
            if (message.password != null && message.hasOwnProperty("password"))
                object.password = message.password;
            return object;
        };

        /**
         * Converts this AuthMessage to JSON.
         * @function toJSON
         * @memberof proto.AuthMessage
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        AuthMessage.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for AuthMessage
         * @function getTypeUrl
         * @memberof proto.AuthMessage
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        AuthMessage.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/proto.AuthMessage";
        };

        return AuthMessage;
    })();

    proto.RegisterMessage = (function() {

        /**
         * Properties of a RegisterMessage.
         * @memberof proto
         * @interface IRegisterMessage
         * @property {string|null} [username] RegisterMessage username
         * @property {string|null} [password] RegisterMessage password
         * @property {string|null} [email] RegisterMessage email
         */

        /**
         * Constructs a new RegisterMessage.
         * @memberof proto
         * @classdesc Represents a RegisterMessage.
         * @implements IRegisterMessage
         * @constructor
         * @param {proto.IRegisterMessage=} [properties] Properties to set
         */
        function RegisterMessage(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * RegisterMessage username.
         * @member {string} username
         * @memberof proto.RegisterMessage
         * @instance
         */
        RegisterMessage.prototype.username = "";

        /**
         * RegisterMessage password.
         * @member {string} password
         * @memberof proto.RegisterMessage
         * @instance
         */
        RegisterMessage.prototype.password = "";

        /**
         * RegisterMessage email.
         * @member {string} email
         * @memberof proto.RegisterMessage
         * @instance
         */
        RegisterMessage.prototype.email = "";

        /**
         * Creates a new RegisterMessage instance using the specified properties.
         * @function create
         * @memberof proto.RegisterMessage
         * @static
         * @param {proto.IRegisterMessage=} [properties] Properties to set
         * @returns {proto.RegisterMessage} RegisterMessage instance
         */
        RegisterMessage.create = function create(properties) {
            return new RegisterMessage(properties);
        };

        /**
         * Encodes the specified RegisterMessage message. Does not implicitly {@link proto.RegisterMessage.verify|verify} messages.
         * @function encode
         * @memberof proto.RegisterMessage
         * @static
         * @param {proto.IRegisterMessage} message RegisterMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegisterMessage.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.username != null && Object.hasOwnProperty.call(message, "username"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.username);
            if (message.password != null && Object.hasOwnProperty.call(message, "password"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.password);
            if (message.email != null && Object.hasOwnProperty.call(message, "email"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.email);
            return writer;
        };

        /**
         * Encodes the specified RegisterMessage message, length delimited. Does not implicitly {@link proto.RegisterMessage.verify|verify} messages.
         * @function encodeDelimited
         * @memberof proto.RegisterMessage
         * @static
         * @param {proto.IRegisterMessage} message RegisterMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegisterMessage.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RegisterMessage message from the specified reader or buffer.
         * @function decode
         * @memberof proto.RegisterMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {proto.RegisterMessage} RegisterMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegisterMessage.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.proto.RegisterMessage();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.username = reader.string();
                        break;
                    }
                case 2: {
                        message.password = reader.string();
                        break;
                    }
                case 3: {
                        message.email = reader.string();
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RegisterMessage message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof proto.RegisterMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {proto.RegisterMessage} RegisterMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegisterMessage.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RegisterMessage message.
         * @function verify
         * @memberof proto.RegisterMessage
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        RegisterMessage.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.username != null && message.hasOwnProperty("username"))
                if (!$util.isString(message.username))
                    return "username: string expected";
            if (message.password != null && message.hasOwnProperty("password"))
                if (!$util.isString(message.password))
                    return "password: string expected";
            if (message.email != null && message.hasOwnProperty("email"))
                if (!$util.isString(message.email))
                    return "email: string expected";
            return null;
        };

        /**
         * Creates a RegisterMessage message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof proto.RegisterMessage
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {proto.RegisterMessage} RegisterMessage
         */
        RegisterMessage.fromObject = function fromObject(object) {
            if (object instanceof $root.proto.RegisterMessage)
                return object;
            var message = new $root.proto.RegisterMessage();
            if (object.username != null)
                message.username = String(object.username);
            if (object.password != null)
                message.password = String(object.password);
            if (object.email != null)
                message.email = String(object.email);
            return message;
        };

        /**
         * Creates a plain object from a RegisterMessage message. Also converts values to other types if specified.
         * @function toObject
         * @memberof proto.RegisterMessage
         * @static
         * @param {proto.RegisterMessage} message RegisterMessage
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RegisterMessage.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.username = "";
                object.password = "";
                object.email = "";
            }
            if (message.username != null && message.hasOwnProperty("username"))
                object.username = message.username;
            if (message.password != null && message.hasOwnProperty("password"))
                object.password = message.password;
            if (message.email != null && message.hasOwnProperty("email"))
                object.email = message.email;
            return object;
        };

        /**
         * Converts this RegisterMessage to JSON.
         * @function toJSON
         * @memberof proto.RegisterMessage
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        RegisterMessage.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for RegisterMessage
         * @function getTypeUrl
         * @memberof proto.RegisterMessage
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        RegisterMessage.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/proto.RegisterMessage";
        };

        return RegisterMessage;
    })();

    proto.ChatMessage = (function() {

        /**
         * Properties of a ChatMessage.
         * @memberof proto
         * @interface IChatMessage
         * @property {number|null} [senderId] ChatMessage senderId
         * @property {number|null} [receiverId] ChatMessage receiverId
         * @property {string|null} [content] ChatMessage content
         * @property {string|null} [messageType] ChatMessage messageType
         */

        /**
         * Constructs a new ChatMessage.
         * @memberof proto
         * @classdesc Represents a ChatMessage.
         * @implements IChatMessage
         * @constructor
         * @param {proto.IChatMessage=} [properties] Properties to set
         */
        function ChatMessage(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * ChatMessage senderId.
         * @member {number} senderId
         * @memberof proto.ChatMessage
         * @instance
         */
        ChatMessage.prototype.senderId = 0;

        /**
         * ChatMessage receiverId.
         * @member {number} receiverId
         * @memberof proto.ChatMessage
         * @instance
         */
        ChatMessage.prototype.receiverId = 0;

        /**
         * ChatMessage content.
         * @member {string} content
         * @memberof proto.ChatMessage
         * @instance
         */
        ChatMessage.prototype.content = "";

        /**
         * ChatMessage messageType.
         * @member {string} messageType
         * @memberof proto.ChatMessage
         * @instance
         */
        ChatMessage.prototype.messageType = "";

        /**
         * Creates a new ChatMessage instance using the specified properties.
         * @function create
         * @memberof proto.ChatMessage
         * @static
         * @param {proto.IChatMessage=} [properties] Properties to set
         * @returns {proto.ChatMessage} ChatMessage instance
         */
        ChatMessage.create = function create(properties) {
            return new ChatMessage(properties);
        };

        /**
         * Encodes the specified ChatMessage message. Does not implicitly {@link proto.ChatMessage.verify|verify} messages.
         * @function encode
         * @memberof proto.ChatMessage
         * @static
         * @param {proto.IChatMessage} message ChatMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ChatMessage.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.senderId != null && Object.hasOwnProperty.call(message, "senderId"))
                writer.uint32(/* id 1, wireType 0 =*/8).int32(message.senderId);
            if (message.receiverId != null && Object.hasOwnProperty.call(message, "receiverId"))
                writer.uint32(/* id 2, wireType 0 =*/16).int32(message.receiverId);
            if (message.content != null && Object.hasOwnProperty.call(message, "content"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.content);
            if (message.messageType != null && Object.hasOwnProperty.call(message, "messageType"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.messageType);
            return writer;
        };

        /**
         * Encodes the specified ChatMessage message, length delimited. Does not implicitly {@link proto.ChatMessage.verify|verify} messages.
         * @function encodeDelimited
         * @memberof proto.ChatMessage
         * @static
         * @param {proto.IChatMessage} message ChatMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ChatMessage.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a ChatMessage message from the specified reader or buffer.
         * @function decode
         * @memberof proto.ChatMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {proto.ChatMessage} ChatMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        ChatMessage.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.proto.ChatMessage();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.senderId = reader.int32();
                        break;
                    }
                case 2: {
                        message.receiverId = reader.int32();
                        break;
                    }
                case 3: {
                        message.content = reader.string();
                        break;
                    }
                case 4: {
                        message.messageType = reader.string();
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a ChatMessage message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof proto.ChatMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {proto.ChatMessage} ChatMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        ChatMessage.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a ChatMessage message.
         * @function verify
         * @memberof proto.ChatMessage
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        ChatMessage.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.senderId != null && message.hasOwnProperty("senderId"))
                if (!$util.isInteger(message.senderId))
                    return "senderId: integer expected";
            if (message.receiverId != null && message.hasOwnProperty("receiverId"))
                if (!$util.isInteger(message.receiverId))
                    return "receiverId: integer expected";
            if (message.content != null && message.hasOwnProperty("content"))
                if (!$util.isString(message.content))
                    return "content: string expected";
            if (message.messageType != null && message.hasOwnProperty("messageType"))
                if (!$util.isString(message.messageType))
                    return "messageType: string expected";
            return null;
        };

        /**
         * Creates a ChatMessage message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof proto.ChatMessage
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {proto.ChatMessage} ChatMessage
         */
        ChatMessage.fromObject = function fromObject(object) {
            if (object instanceof $root.proto.ChatMessage)
                return object;
            var message = new $root.proto.ChatMessage();
            if (object.senderId != null)
                message.senderId = object.senderId | 0;
            if (object.receiverId != null)
                message.receiverId = object.receiverId | 0;
            if (object.content != null)
                message.content = String(object.content);
            if (object.messageType != null)
                message.messageType = String(object.messageType);
            return message;
        };

        /**
         * Creates a plain object from a ChatMessage message. Also converts values to other types if specified.
         * @function toObject
         * @memberof proto.ChatMessage
         * @static
         * @param {proto.ChatMessage} message ChatMessage
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        ChatMessage.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.senderId = 0;
                object.receiverId = 0;
                object.content = "";
                object.messageType = "";
            }
            if (message.senderId != null && message.hasOwnProperty("senderId"))
                object.senderId = message.senderId;
            if (message.receiverId != null && message.hasOwnProperty("receiverId"))
                object.receiverId = message.receiverId;
            if (message.content != null && message.hasOwnProperty("content"))
                object.content = message.content;
            if (message.messageType != null && message.hasOwnProperty("messageType"))
                object.messageType = message.messageType;
            return object;
        };

        /**
         * Converts this ChatMessage to JSON.
         * @function toJSON
         * @memberof proto.ChatMessage
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        ChatMessage.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for ChatMessage
         * @function getTypeUrl
         * @memberof proto.ChatMessage
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        ChatMessage.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/proto.ChatMessage";
        };

        return ChatMessage;
    })();

    proto.SystemMessage = (function() {

        /**
         * Properties of a SystemMessage.
         * @memberof proto
         * @interface ISystemMessage
         * @property {string|null} [type] SystemMessage type
         * @property {string|null} [content] SystemMessage content
         */

        /**
         * Constructs a new SystemMessage.
         * @memberof proto
         * @classdesc Represents a SystemMessage.
         * @implements ISystemMessage
         * @constructor
         * @param {proto.ISystemMessage=} [properties] Properties to set
         */
        function SystemMessage(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * SystemMessage type.
         * @member {string} type
         * @memberof proto.SystemMessage
         * @instance
         */
        SystemMessage.prototype.type = "";

        /**
         * SystemMessage content.
         * @member {string} content
         * @memberof proto.SystemMessage
         * @instance
         */
        SystemMessage.prototype.content = "";

        /**
         * Creates a new SystemMessage instance using the specified properties.
         * @function create
         * @memberof proto.SystemMessage
         * @static
         * @param {proto.ISystemMessage=} [properties] Properties to set
         * @returns {proto.SystemMessage} SystemMessage instance
         */
        SystemMessage.create = function create(properties) {
            return new SystemMessage(properties);
        };

        /**
         * Encodes the specified SystemMessage message. Does not implicitly {@link proto.SystemMessage.verify|verify} messages.
         * @function encode
         * @memberof proto.SystemMessage
         * @static
         * @param {proto.ISystemMessage} message SystemMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        SystemMessage.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.type != null && Object.hasOwnProperty.call(message, "type"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.type);
            if (message.content != null && Object.hasOwnProperty.call(message, "content"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.content);
            return writer;
        };

        /**
         * Encodes the specified SystemMessage message, length delimited. Does not implicitly {@link proto.SystemMessage.verify|verify} messages.
         * @function encodeDelimited
         * @memberof proto.SystemMessage
         * @static
         * @param {proto.ISystemMessage} message SystemMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        SystemMessage.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a SystemMessage message from the specified reader or buffer.
         * @function decode
         * @memberof proto.SystemMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {proto.SystemMessage} SystemMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        SystemMessage.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.proto.SystemMessage();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.type = reader.string();
                        break;
                    }
                case 2: {
                        message.content = reader.string();
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a SystemMessage message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof proto.SystemMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {proto.SystemMessage} SystemMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        SystemMessage.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a SystemMessage message.
         * @function verify
         * @memberof proto.SystemMessage
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        SystemMessage.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.type != null && message.hasOwnProperty("type"))
                if (!$util.isString(message.type))
                    return "type: string expected";
            if (message.content != null && message.hasOwnProperty("content"))
                if (!$util.isString(message.content))
                    return "content: string expected";
            return null;
        };

        /**
         * Creates a SystemMessage message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof proto.SystemMessage
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {proto.SystemMessage} SystemMessage
         */
        SystemMessage.fromObject = function fromObject(object) {
            if (object instanceof $root.proto.SystemMessage)
                return object;
            var message = new $root.proto.SystemMessage();
            if (object.type != null)
                message.type = String(object.type);
            if (object.content != null)
                message.content = String(object.content);
            return message;
        };

        /**
         * Creates a plain object from a SystemMessage message. Also converts values to other types if specified.
         * @function toObject
         * @memberof proto.SystemMessage
         * @static
         * @param {proto.SystemMessage} message SystemMessage
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        SystemMessage.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.type = "";
                object.content = "";
            }
            if (message.type != null && message.hasOwnProperty("type"))
                object.type = message.type;
            if (message.content != null && message.hasOwnProperty("content"))
                object.content = message.content;
            return object;
        };

        /**
         * Converts this SystemMessage to JSON.
         * @function toJSON
         * @memberof proto.SystemMessage
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        SystemMessage.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for SystemMessage
         * @function getTypeUrl
         * @memberof proto.SystemMessage
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        SystemMessage.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/proto.SystemMessage";
        };

        return SystemMessage;
    })();

    proto.StudyRoomMessage = (function() {

        /**
         * Properties of a StudyRoomMessage.
         * @memberof proto
         * @interface IStudyRoomMessage
         * @property {proto.StudyRoomMessage.Operation|null} [operation] StudyRoomMessage operation
         * @property {number|null} [roomId] StudyRoomMessage roomId
         * @property {string|null} [name] StudyRoomMessage name
         * @property {number|null} [maxMembers] StudyRoomMessage maxMembers
         * @property {boolean|null} [isPrivate] StudyRoomMessage isPrivate
         * @property {string|null} [duration] StudyRoomMessage duration
         * @property {string|null} [shareLink] StudyRoomMessage shareLink
         */

        /**
         * Constructs a new StudyRoomMessage.
         * @memberof proto
         * @classdesc Represents a StudyRoomMessage.
         * @implements IStudyRoomMessage
         * @constructor
         * @param {proto.IStudyRoomMessage=} [properties] Properties to set
         */
        function StudyRoomMessage(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * StudyRoomMessage operation.
         * @member {proto.StudyRoomMessage.Operation} operation
         * @memberof proto.StudyRoomMessage
         * @instance
         */
        StudyRoomMessage.prototype.operation = 0;

        /**
         * StudyRoomMessage roomId.
         * @member {number} roomId
         * @memberof proto.StudyRoomMessage
         * @instance
         */
        StudyRoomMessage.prototype.roomId = 0;

        /**
         * StudyRoomMessage name.
         * @member {string} name
         * @memberof proto.StudyRoomMessage
         * @instance
         */
        StudyRoomMessage.prototype.name = "";

        /**
         * StudyRoomMessage maxMembers.
         * @member {number} maxMembers
         * @memberof proto.StudyRoomMessage
         * @instance
         */
        StudyRoomMessage.prototype.maxMembers = 0;

        /**
         * StudyRoomMessage isPrivate.
         * @member {boolean} isPrivate
         * @memberof proto.StudyRoomMessage
         * @instance
         */
        StudyRoomMessage.prototype.isPrivate = false;

        /**
         * StudyRoomMessage duration.
         * @member {string} duration
         * @memberof proto.StudyRoomMessage
         * @instance
         */
        StudyRoomMessage.prototype.duration = "";

        /**
         * StudyRoomMessage shareLink.
         * @member {string} shareLink
         * @memberof proto.StudyRoomMessage
         * @instance
         */
        StudyRoomMessage.prototype.shareLink = "";

        /**
         * Creates a new StudyRoomMessage instance using the specified properties.
         * @function create
         * @memberof proto.StudyRoomMessage
         * @static
         * @param {proto.IStudyRoomMessage=} [properties] Properties to set
         * @returns {proto.StudyRoomMessage} StudyRoomMessage instance
         */
        StudyRoomMessage.create = function create(properties) {
            return new StudyRoomMessage(properties);
        };

        /**
         * Encodes the specified StudyRoomMessage message. Does not implicitly {@link proto.StudyRoomMessage.verify|verify} messages.
         * @function encode
         * @memberof proto.StudyRoomMessage
         * @static
         * @param {proto.IStudyRoomMessage} message StudyRoomMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        StudyRoomMessage.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.operation != null && Object.hasOwnProperty.call(message, "operation"))
                writer.uint32(/* id 1, wireType 0 =*/8).int32(message.operation);
            if (message.roomId != null && Object.hasOwnProperty.call(message, "roomId"))
                writer.uint32(/* id 2, wireType 0 =*/16).int32(message.roomId);
            if (message.name != null && Object.hasOwnProperty.call(message, "name"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.name);
            if (message.maxMembers != null && Object.hasOwnProperty.call(message, "maxMembers"))
                writer.uint32(/* id 4, wireType 0 =*/32).int32(message.maxMembers);
            if (message.isPrivate != null && Object.hasOwnProperty.call(message, "isPrivate"))
                writer.uint32(/* id 5, wireType 0 =*/40).bool(message.isPrivate);
            if (message.duration != null && Object.hasOwnProperty.call(message, "duration"))
                writer.uint32(/* id 6, wireType 2 =*/50).string(message.duration);
            if (message.shareLink != null && Object.hasOwnProperty.call(message, "shareLink"))
                writer.uint32(/* id 7, wireType 2 =*/58).string(message.shareLink);
            return writer;
        };

        /**
         * Encodes the specified StudyRoomMessage message, length delimited. Does not implicitly {@link proto.StudyRoomMessage.verify|verify} messages.
         * @function encodeDelimited
         * @memberof proto.StudyRoomMessage
         * @static
         * @param {proto.IStudyRoomMessage} message StudyRoomMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        StudyRoomMessage.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a StudyRoomMessage message from the specified reader or buffer.
         * @function decode
         * @memberof proto.StudyRoomMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {proto.StudyRoomMessage} StudyRoomMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        StudyRoomMessage.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.proto.StudyRoomMessage();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.operation = reader.int32();
                        break;
                    }
                case 2: {
                        message.roomId = reader.int32();
                        break;
                    }
                case 3: {
                        message.name = reader.string();
                        break;
                    }
                case 4: {
                        message.maxMembers = reader.int32();
                        break;
                    }
                case 5: {
                        message.isPrivate = reader.bool();
                        break;
                    }
                case 6: {
                        message.duration = reader.string();
                        break;
                    }
                case 7: {
                        message.shareLink = reader.string();
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a StudyRoomMessage message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof proto.StudyRoomMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {proto.StudyRoomMessage} StudyRoomMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        StudyRoomMessage.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a StudyRoomMessage message.
         * @function verify
         * @memberof proto.StudyRoomMessage
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        StudyRoomMessage.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.operation != null && message.hasOwnProperty("operation"))
                switch (message.operation) {
                default:
                    return "operation: enum value expected";
                case 0:
                case 1:
                case 2:
                case 3:
                case 4:
                    break;
                }
            if (message.roomId != null && message.hasOwnProperty("roomId"))
                if (!$util.isInteger(message.roomId))
                    return "roomId: integer expected";
            if (message.name != null && message.hasOwnProperty("name"))
                if (!$util.isString(message.name))
                    return "name: string expected";
            if (message.maxMembers != null && message.hasOwnProperty("maxMembers"))
                if (!$util.isInteger(message.maxMembers))
                    return "maxMembers: integer expected";
            if (message.isPrivate != null && message.hasOwnProperty("isPrivate"))
                if (typeof message.isPrivate !== "boolean")
                    return "isPrivate: boolean expected";
            if (message.duration != null && message.hasOwnProperty("duration"))
                if (!$util.isString(message.duration))
                    return "duration: string expected";
            if (message.shareLink != null && message.hasOwnProperty("shareLink"))
                if (!$util.isString(message.shareLink))
                    return "shareLink: string expected";
            return null;
        };

        /**
         * Creates a StudyRoomMessage message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof proto.StudyRoomMessage
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {proto.StudyRoomMessage} StudyRoomMessage
         */
        StudyRoomMessage.fromObject = function fromObject(object) {
            if (object instanceof $root.proto.StudyRoomMessage)
                return object;
            var message = new $root.proto.StudyRoomMessage();
            switch (object.operation) {
            default:
                if (typeof object.operation === "number") {
                    message.operation = object.operation;
                    break;
                }
                break;
            case "CREATE":
            case 0:
                message.operation = 0;
                break;
            case "JOIN":
            case 1:
                message.operation = 1;
                break;
            case "LEAVE":
            case 2:
                message.operation = 2;
                break;
            case "DESTROY":
            case 3:
                message.operation = 3;
                break;
            case "GET_DETAIL":
            case 4:
                message.operation = 4;
                break;
            }
            if (object.roomId != null)
                message.roomId = object.roomId | 0;
            if (object.name != null)
                message.name = String(object.name);
            if (object.maxMembers != null)
                message.maxMembers = object.maxMembers | 0;
            if (object.isPrivate != null)
                message.isPrivate = Boolean(object.isPrivate);
            if (object.duration != null)
                message.duration = String(object.duration);
            if (object.shareLink != null)
                message.shareLink = String(object.shareLink);
            return message;
        };

        /**
         * Creates a plain object from a StudyRoomMessage message. Also converts values to other types if specified.
         * @function toObject
         * @memberof proto.StudyRoomMessage
         * @static
         * @param {proto.StudyRoomMessage} message StudyRoomMessage
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        StudyRoomMessage.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.operation = options.enums === String ? "CREATE" : 0;
                object.roomId = 0;
                object.name = "";
                object.maxMembers = 0;
                object.isPrivate = false;
                object.duration = "";
                object.shareLink = "";
            }
            if (message.operation != null && message.hasOwnProperty("operation"))
                object.operation = options.enums === String ? $root.proto.StudyRoomMessage.Operation[message.operation] === undefined ? message.operation : $root.proto.StudyRoomMessage.Operation[message.operation] : message.operation;
            if (message.roomId != null && message.hasOwnProperty("roomId"))
                object.roomId = message.roomId;
            if (message.name != null && message.hasOwnProperty("name"))
                object.name = message.name;
            if (message.maxMembers != null && message.hasOwnProperty("maxMembers"))
                object.maxMembers = message.maxMembers;
            if (message.isPrivate != null && message.hasOwnProperty("isPrivate"))
                object.isPrivate = message.isPrivate;
            if (message.duration != null && message.hasOwnProperty("duration"))
                object.duration = message.duration;
            if (message.shareLink != null && message.hasOwnProperty("shareLink"))
                object.shareLink = message.shareLink;
            return object;
        };

        /**
         * Converts this StudyRoomMessage to JSON.
         * @function toJSON
         * @memberof proto.StudyRoomMessage
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        StudyRoomMessage.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for StudyRoomMessage
         * @function getTypeUrl
         * @memberof proto.StudyRoomMessage
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        StudyRoomMessage.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/proto.StudyRoomMessage";
        };

        /**
         * Operation enum.
         * @name proto.StudyRoomMessage.Operation
         * @enum {number}
         * @property {number} CREATE=0 CREATE value
         * @property {number} JOIN=1 JOIN value
         * @property {number} LEAVE=2 LEAVE value
         * @property {number} DESTROY=3 DESTROY value
         * @property {number} GET_DETAIL=4 GET_DETAIL value
         */
        StudyRoomMessage.Operation = (function() {
            var valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "CREATE"] = 0;
            values[valuesById[1] = "JOIN"] = 1;
            values[valuesById[2] = "LEAVE"] = 2;
            values[valuesById[3] = "DESTROY"] = 3;
            values[valuesById[4] = "GET_DETAIL"] = 4;
            return values;
        })();

        return StudyRoomMessage;
    })();

    proto.FriendMessage = (function() {

        /**
         * Properties of a FriendMessage.
         * @memberof proto
         * @interface IFriendMessage
         * @property {proto.FriendMessage.Operation|null} [operation] FriendMessage operation
         * @property {number|null} [friendId] FriendMessage friendId
         * @property {string|null} [action] FriendMessage action
         * @property {string|null} [contractType] FriendMessage contractType
         * @property {string|null} [contractTerms] FriendMessage contractTerms
         * @property {number|null} [contractId] FriendMessage contractId
         */

        /**
         * Constructs a new FriendMessage.
         * @memberof proto
         * @classdesc Represents a FriendMessage.
         * @implements IFriendMessage
         * @constructor
         * @param {proto.IFriendMessage=} [properties] Properties to set
         */
        function FriendMessage(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * FriendMessage operation.
         * @member {proto.FriendMessage.Operation} operation
         * @memberof proto.FriendMessage
         * @instance
         */
        FriendMessage.prototype.operation = 0;

        /**
         * FriendMessage friendId.
         * @member {number} friendId
         * @memberof proto.FriendMessage
         * @instance
         */
        FriendMessage.prototype.friendId = 0;

        /**
         * FriendMessage action.
         * @member {string} action
         * @memberof proto.FriendMessage
         * @instance
         */
        FriendMessage.prototype.action = "";

        /**
         * FriendMessage contractType.
         * @member {string} contractType
         * @memberof proto.FriendMessage
         * @instance
         */
        FriendMessage.prototype.contractType = "";

        /**
         * FriendMessage contractTerms.
         * @member {string} contractTerms
         * @memberof proto.FriendMessage
         * @instance
         */
        FriendMessage.prototype.contractTerms = "";

        /**
         * FriendMessage contractId.
         * @member {number} contractId
         * @memberof proto.FriendMessage
         * @instance
         */
        FriendMessage.prototype.contractId = 0;

        /**
         * Creates a new FriendMessage instance using the specified properties.
         * @function create
         * @memberof proto.FriendMessage
         * @static
         * @param {proto.IFriendMessage=} [properties] Properties to set
         * @returns {proto.FriendMessage} FriendMessage instance
         */
        FriendMessage.create = function create(properties) {
            return new FriendMessage(properties);
        };

        /**
         * Encodes the specified FriendMessage message. Does not implicitly {@link proto.FriendMessage.verify|verify} messages.
         * @function encode
         * @memberof proto.FriendMessage
         * @static
         * @param {proto.IFriendMessage} message FriendMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        FriendMessage.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.operation != null && Object.hasOwnProperty.call(message, "operation"))
                writer.uint32(/* id 1, wireType 0 =*/8).int32(message.operation);
            if (message.friendId != null && Object.hasOwnProperty.call(message, "friendId"))
                writer.uint32(/* id 2, wireType 0 =*/16).int32(message.friendId);
            if (message.action != null && Object.hasOwnProperty.call(message, "action"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.action);
            if (message.contractType != null && Object.hasOwnProperty.call(message, "contractType"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.contractType);
            if (message.contractTerms != null && Object.hasOwnProperty.call(message, "contractTerms"))
                writer.uint32(/* id 5, wireType 2 =*/42).string(message.contractTerms);
            if (message.contractId != null && Object.hasOwnProperty.call(message, "contractId"))
                writer.uint32(/* id 6, wireType 0 =*/48).int32(message.contractId);
            return writer;
        };

        /**
         * Encodes the specified FriendMessage message, length delimited. Does not implicitly {@link proto.FriendMessage.verify|verify} messages.
         * @function encodeDelimited
         * @memberof proto.FriendMessage
         * @static
         * @param {proto.IFriendMessage} message FriendMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        FriendMessage.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a FriendMessage message from the specified reader or buffer.
         * @function decode
         * @memberof proto.FriendMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {proto.FriendMessage} FriendMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        FriendMessage.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.proto.FriendMessage();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.operation = reader.int32();
                        break;
                    }
                case 2: {
                        message.friendId = reader.int32();
                        break;
                    }
                case 3: {
                        message.action = reader.string();
                        break;
                    }
                case 4: {
                        message.contractType = reader.string();
                        break;
                    }
                case 5: {
                        message.contractTerms = reader.string();
                        break;
                    }
                case 6: {
                        message.contractId = reader.int32();
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a FriendMessage message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof proto.FriendMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {proto.FriendMessage} FriendMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        FriendMessage.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a FriendMessage message.
         * @function verify
         * @memberof proto.FriendMessage
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        FriendMessage.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.operation != null && message.hasOwnProperty("operation"))
                switch (message.operation) {
                default:
                    return "operation: enum value expected";
                case 0:
                case 1:
                case 2:
                case 3:
                case 4:
                case 5:
                case 6:
                    break;
                }
            if (message.friendId != null && message.hasOwnProperty("friendId"))
                if (!$util.isInteger(message.friendId))
                    return "friendId: integer expected";
            if (message.action != null && message.hasOwnProperty("action"))
                if (!$util.isString(message.action))
                    return "action: string expected";
            if (message.contractType != null && message.hasOwnProperty("contractType"))
                if (!$util.isString(message.contractType))
                    return "contractType: string expected";
            if (message.contractTerms != null && message.hasOwnProperty("contractTerms"))
                if (!$util.isString(message.contractTerms))
                    return "contractTerms: string expected";
            if (message.contractId != null && message.hasOwnProperty("contractId"))
                if (!$util.isInteger(message.contractId))
                    return "contractId: integer expected";
            return null;
        };

        /**
         * Creates a FriendMessage message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof proto.FriendMessage
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {proto.FriendMessage} FriendMessage
         */
        FriendMessage.fromObject = function fromObject(object) {
            if (object instanceof $root.proto.FriendMessage)
                return object;
            var message = new $root.proto.FriendMessage();
            switch (object.operation) {
            default:
                if (typeof object.operation === "number") {
                    message.operation = object.operation;
                    break;
                }
                break;
            case "SEND_REQUEST":
            case 0:
                message.operation = 0;
                break;
            case "HANDLE_REQUEST":
            case 1:
                message.operation = 1;
                break;
            case "GET_LIST":
            case 2:
                message.operation = 2;
                break;
            case "DELETE":
            case 3:
                message.operation = 3;
                break;
            case "CREATE_CONTRACT":
            case 4:
                message.operation = 4;
                break;
            case "TERMINATE_CONTRACT":
            case 5:
                message.operation = 5;
                break;
            case "GET_CONTRACTS":
            case 6:
                message.operation = 6;
                break;
            }
            if (object.friendId != null)
                message.friendId = object.friendId | 0;
            if (object.action != null)
                message.action = String(object.action);
            if (object.contractType != null)
                message.contractType = String(object.contractType);
            if (object.contractTerms != null)
                message.contractTerms = String(object.contractTerms);
            if (object.contractId != null)
                message.contractId = object.contractId | 0;
            return message;
        };

        /**
         * Creates a plain object from a FriendMessage message. Also converts values to other types if specified.
         * @function toObject
         * @memberof proto.FriendMessage
         * @static
         * @param {proto.FriendMessage} message FriendMessage
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        FriendMessage.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.operation = options.enums === String ? "SEND_REQUEST" : 0;
                object.friendId = 0;
                object.action = "";
                object.contractType = "";
                object.contractTerms = "";
                object.contractId = 0;
            }
            if (message.operation != null && message.hasOwnProperty("operation"))
                object.operation = options.enums === String ? $root.proto.FriendMessage.Operation[message.operation] === undefined ? message.operation : $root.proto.FriendMessage.Operation[message.operation] : message.operation;
            if (message.friendId != null && message.hasOwnProperty("friendId"))
                object.friendId = message.friendId;
            if (message.action != null && message.hasOwnProperty("action"))
                object.action = message.action;
            if (message.contractType != null && message.hasOwnProperty("contractType"))
                object.contractType = message.contractType;
            if (message.contractTerms != null && message.hasOwnProperty("contractTerms"))
                object.contractTerms = message.contractTerms;
            if (message.contractId != null && message.hasOwnProperty("contractId"))
                object.contractId = message.contractId;
            return object;
        };

        /**
         * Converts this FriendMessage to JSON.
         * @function toJSON
         * @memberof proto.FriendMessage
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        FriendMessage.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for FriendMessage
         * @function getTypeUrl
         * @memberof proto.FriendMessage
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        FriendMessage.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/proto.FriendMessage";
        };

        /**
         * Operation enum.
         * @name proto.FriendMessage.Operation
         * @enum {number}
         * @property {number} SEND_REQUEST=0 SEND_REQUEST value
         * @property {number} HANDLE_REQUEST=1 HANDLE_REQUEST value
         * @property {number} GET_LIST=2 GET_LIST value
         * @property {number} DELETE=3 DELETE value
         * @property {number} CREATE_CONTRACT=4 CREATE_CONTRACT value
         * @property {number} TERMINATE_CONTRACT=5 TERMINATE_CONTRACT value
         * @property {number} GET_CONTRACTS=6 GET_CONTRACTS value
         */
        FriendMessage.Operation = (function() {
            var valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "SEND_REQUEST"] = 0;
            values[valuesById[1] = "HANDLE_REQUEST"] = 1;
            values[valuesById[2] = "GET_LIST"] = 2;
            values[valuesById[3] = "DELETE"] = 3;
            values[valuesById[4] = "CREATE_CONTRACT"] = 4;
            values[valuesById[5] = "TERMINATE_CONTRACT"] = 5;
            values[valuesById[6] = "GET_CONTRACTS"] = 6;
            return values;
        })();

        return FriendMessage;
    })();

    proto.ProfileMessage = (function() {

        /**
         * Properties of a ProfileMessage.
         * @memberof proto
         * @interface IProfileMessage
         * @property {proto.ProfileMessage.Operation|null} [operation] ProfileMessage operation
         * @property {string|null} [username] ProfileMessage username
         * @property {string|null} [signature] ProfileMessage signature
         * @property {string|null} [studyDirection] ProfileMessage studyDirection
         */

        /**
         * Constructs a new ProfileMessage.
         * @memberof proto
         * @classdesc Represents a ProfileMessage.
         * @implements IProfileMessage
         * @constructor
         * @param {proto.IProfileMessage=} [properties] Properties to set
         */
        function ProfileMessage(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * ProfileMessage operation.
         * @member {proto.ProfileMessage.Operation} operation
         * @memberof proto.ProfileMessage
         * @instance
         */
        ProfileMessage.prototype.operation = 0;

        /**
         * ProfileMessage username.
         * @member {string} username
         * @memberof proto.ProfileMessage
         * @instance
         */
        ProfileMessage.prototype.username = "";

        /**
         * ProfileMessage signature.
         * @member {string} signature
         * @memberof proto.ProfileMessage
         * @instance
         */
        ProfileMessage.prototype.signature = "";

        /**
         * ProfileMessage studyDirection.
         * @member {string} studyDirection
         * @memberof proto.ProfileMessage
         * @instance
         */
        ProfileMessage.prototype.studyDirection = "";

        /**
         * Creates a new ProfileMessage instance using the specified properties.
         * @function create
         * @memberof proto.ProfileMessage
         * @static
         * @param {proto.IProfileMessage=} [properties] Properties to set
         * @returns {proto.ProfileMessage} ProfileMessage instance
         */
        ProfileMessage.create = function create(properties) {
            return new ProfileMessage(properties);
        };

        /**
         * Encodes the specified ProfileMessage message. Does not implicitly {@link proto.ProfileMessage.verify|verify} messages.
         * @function encode
         * @memberof proto.ProfileMessage
         * @static
         * @param {proto.IProfileMessage} message ProfileMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ProfileMessage.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.operation != null && Object.hasOwnProperty.call(message, "operation"))
                writer.uint32(/* id 1, wireType 0 =*/8).int32(message.operation);
            if (message.username != null && Object.hasOwnProperty.call(message, "username"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.username);
            if (message.signature != null && Object.hasOwnProperty.call(message, "signature"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.signature);
            if (message.studyDirection != null && Object.hasOwnProperty.call(message, "studyDirection"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.studyDirection);
            return writer;
        };

        /**
         * Encodes the specified ProfileMessage message, length delimited. Does not implicitly {@link proto.ProfileMessage.verify|verify} messages.
         * @function encodeDelimited
         * @memberof proto.ProfileMessage
         * @static
         * @param {proto.IProfileMessage} message ProfileMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ProfileMessage.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a ProfileMessage message from the specified reader or buffer.
         * @function decode
         * @memberof proto.ProfileMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {proto.ProfileMessage} ProfileMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        ProfileMessage.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.proto.ProfileMessage();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.operation = reader.int32();
                        break;
                    }
                case 2: {
                        message.username = reader.string();
                        break;
                    }
                case 3: {
                        message.signature = reader.string();
                        break;
                    }
                case 4: {
                        message.studyDirection = reader.string();
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a ProfileMessage message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof proto.ProfileMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {proto.ProfileMessage} ProfileMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        ProfileMessage.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a ProfileMessage message.
         * @function verify
         * @memberof proto.ProfileMessage
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        ProfileMessage.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.operation != null && message.hasOwnProperty("operation"))
                switch (message.operation) {
                default:
                    return "operation: enum value expected";
                case 0:
                case 1:
                    break;
                }
            if (message.username != null && message.hasOwnProperty("username"))
                if (!$util.isString(message.username))
                    return "username: string expected";
            if (message.signature != null && message.hasOwnProperty("signature"))
                if (!$util.isString(message.signature))
                    return "signature: string expected";
            if (message.studyDirection != null && message.hasOwnProperty("studyDirection"))
                if (!$util.isString(message.studyDirection))
                    return "studyDirection: string expected";
            return null;
        };

        /**
         * Creates a ProfileMessage message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof proto.ProfileMessage
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {proto.ProfileMessage} ProfileMessage
         */
        ProfileMessage.fromObject = function fromObject(object) {
            if (object instanceof $root.proto.ProfileMessage)
                return object;
            var message = new $root.proto.ProfileMessage();
            switch (object.operation) {
            default:
                if (typeof object.operation === "number") {
                    message.operation = object.operation;
                    break;
                }
                break;
            case "UPDATE":
            case 0:
                message.operation = 0;
                break;
            case "GET":
            case 1:
                message.operation = 1;
                break;
            }
            if (object.username != null)
                message.username = String(object.username);
            if (object.signature != null)
                message.signature = String(object.signature);
            if (object.studyDirection != null)
                message.studyDirection = String(object.studyDirection);
            return message;
        };

        /**
         * Creates a plain object from a ProfileMessage message. Also converts values to other types if specified.
         * @function toObject
         * @memberof proto.ProfileMessage
         * @static
         * @param {proto.ProfileMessage} message ProfileMessage
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        ProfileMessage.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.operation = options.enums === String ? "UPDATE" : 0;
                object.username = "";
                object.signature = "";
                object.studyDirection = "";
            }
            if (message.operation != null && message.hasOwnProperty("operation"))
                object.operation = options.enums === String ? $root.proto.ProfileMessage.Operation[message.operation] === undefined ? message.operation : $root.proto.ProfileMessage.Operation[message.operation] : message.operation;
            if (message.username != null && message.hasOwnProperty("username"))
                object.username = message.username;
            if (message.signature != null && message.hasOwnProperty("signature"))
                object.signature = message.signature;
            if (message.studyDirection != null && message.hasOwnProperty("studyDirection"))
                object.studyDirection = message.studyDirection;
            return object;
        };

        /**
         * Converts this ProfileMessage to JSON.
         * @function toJSON
         * @memberof proto.ProfileMessage
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        ProfileMessage.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for ProfileMessage
         * @function getTypeUrl
         * @memberof proto.ProfileMessage
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        ProfileMessage.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/proto.ProfileMessage";
        };

        /**
         * Operation enum.
         * @name proto.ProfileMessage.Operation
         * @enum {number}
         * @property {number} UPDATE=0 UPDATE value
         * @property {number} GET=1 GET value
         */
        ProfileMessage.Operation = (function() {
            var valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "UPDATE"] = 0;
            values[valuesById[1] = "GET"] = 1;
            return values;
        })();

        return ProfileMessage;
    })();

    proto.ErrorMessage = (function() {

        /**
         * Properties of an ErrorMessage.
         * @memberof proto
         * @interface IErrorMessage
         * @property {number|null} [code] ErrorMessage code
         * @property {string|null} [message] ErrorMessage message
         * @property {string|null} [detail] ErrorMessage detail
         * @property {string|null} [requestId] ErrorMessage requestId
         */

        /**
         * Constructs a new ErrorMessage.
         * @memberof proto
         * @classdesc Represents an ErrorMessage.
         * @implements IErrorMessage
         * @constructor
         * @param {proto.IErrorMessage=} [properties] Properties to set
         */
        function ErrorMessage(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * ErrorMessage code.
         * @member {number} code
         * @memberof proto.ErrorMessage
         * @instance
         */
        ErrorMessage.prototype.code = 0;

        /**
         * ErrorMessage message.
         * @member {string} message
         * @memberof proto.ErrorMessage
         * @instance
         */
        ErrorMessage.prototype.message = "";

        /**
         * ErrorMessage detail.
         * @member {string} detail
         * @memberof proto.ErrorMessage
         * @instance
         */
        ErrorMessage.prototype.detail = "";

        /**
         * ErrorMessage requestId.
         * @member {string} requestId
         * @memberof proto.ErrorMessage
         * @instance
         */
        ErrorMessage.prototype.requestId = "";

        /**
         * Creates a new ErrorMessage instance using the specified properties.
         * @function create
         * @memberof proto.ErrorMessage
         * @static
         * @param {proto.IErrorMessage=} [properties] Properties to set
         * @returns {proto.ErrorMessage} ErrorMessage instance
         */
        ErrorMessage.create = function create(properties) {
            return new ErrorMessage(properties);
        };

        /**
         * Encodes the specified ErrorMessage message. Does not implicitly {@link proto.ErrorMessage.verify|verify} messages.
         * @function encode
         * @memberof proto.ErrorMessage
         * @static
         * @param {proto.IErrorMessage} message ErrorMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ErrorMessage.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.code != null && Object.hasOwnProperty.call(message, "code"))
                writer.uint32(/* id 1, wireType 0 =*/8).int32(message.code);
            if (message.message != null && Object.hasOwnProperty.call(message, "message"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.message);
            if (message.detail != null && Object.hasOwnProperty.call(message, "detail"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.detail);
            if (message.requestId != null && Object.hasOwnProperty.call(message, "requestId"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.requestId);
            return writer;
        };

        /**
         * Encodes the specified ErrorMessage message, length delimited. Does not implicitly {@link proto.ErrorMessage.verify|verify} messages.
         * @function encodeDelimited
         * @memberof proto.ErrorMessage
         * @static
         * @param {proto.IErrorMessage} message ErrorMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ErrorMessage.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an ErrorMessage message from the specified reader or buffer.
         * @function decode
         * @memberof proto.ErrorMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {proto.ErrorMessage} ErrorMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        ErrorMessage.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.proto.ErrorMessage();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.code = reader.int32();
                        break;
                    }
                case 2: {
                        message.message = reader.string();
                        break;
                    }
                case 3: {
                        message.detail = reader.string();
                        break;
                    }
                case 4: {
                        message.requestId = reader.string();
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an ErrorMessage message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof proto.ErrorMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {proto.ErrorMessage} ErrorMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        ErrorMessage.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an ErrorMessage message.
         * @function verify
         * @memberof proto.ErrorMessage
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        ErrorMessage.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.code != null && message.hasOwnProperty("code"))
                if (!$util.isInteger(message.code))
                    return "code: integer expected";
            if (message.message != null && message.hasOwnProperty("message"))
                if (!$util.isString(message.message))
                    return "message: string expected";
            if (message.detail != null && message.hasOwnProperty("detail"))
                if (!$util.isString(message.detail))
                    return "detail: string expected";
            if (message.requestId != null && message.hasOwnProperty("requestId"))
                if (!$util.isString(message.requestId))
                    return "requestId: string expected";
            return null;
        };

        /**
         * Creates an ErrorMessage message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof proto.ErrorMessage
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {proto.ErrorMessage} ErrorMessage
         */
        ErrorMessage.fromObject = function fromObject(object) {
            if (object instanceof $root.proto.ErrorMessage)
                return object;
            var message = new $root.proto.ErrorMessage();
            if (object.code != null)
                message.code = object.code | 0;
            if (object.message != null)
                message.message = String(object.message);
            if (object.detail != null)
                message.detail = String(object.detail);
            if (object.requestId != null)
                message.requestId = String(object.requestId);
            return message;
        };

        /**
         * Creates a plain object from an ErrorMessage message. Also converts values to other types if specified.
         * @function toObject
         * @memberof proto.ErrorMessage
         * @static
         * @param {proto.ErrorMessage} message ErrorMessage
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        ErrorMessage.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.code = 0;
                object.message = "";
                object.detail = "";
                object.requestId = "";
            }
            if (message.code != null && message.hasOwnProperty("code"))
                object.code = message.code;
            if (message.message != null && message.hasOwnProperty("message"))
                object.message = message.message;
            if (message.detail != null && message.hasOwnProperty("detail"))
                object.detail = message.detail;
            if (message.requestId != null && message.hasOwnProperty("requestId"))
                object.requestId = message.requestId;
            return object;
        };

        /**
         * Converts this ErrorMessage to JSON.
         * @function toJSON
         * @memberof proto.ErrorMessage
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        ErrorMessage.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for ErrorMessage
         * @function getTypeUrl
         * @memberof proto.ErrorMessage
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        ErrorMessage.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/proto.ErrorMessage";
        };

        return ErrorMessage;
    })();

    return proto;
})();

module.exports = $root;
