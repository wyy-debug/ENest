-- 创建用户表
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    avatar VARCHAR(255),
    signature TEXT,
    study_direction VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 创建用户表索引
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- 创建会话表
CREATE TABLE IF NOT EXISTS sessions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token VARCHAR(255) UNIQUE NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_sessions_user FOREIGN KEY (user_id) REFERENCES users(id)
);

-- 创建会话表索引
CREATE INDEX IF NOT EXISTS idx_sessions_token ON sessions(token);
CREATE INDEX IF NOT EXISTS idx_sessions_user_id ON sessions(user_id);
CREATE INDEX IF NOT EXISTS idx_sessions_expires_at ON sessions(expires_at);

-- 创建自习室表
CREATE TABLE IF NOT EXISTS study_rooms (
    id SERIAL PRIMARY KEY,
    owner_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    share_link VARCHAR(255) UNIQUE,
    max_members INTEGER NOT NULL DEFAULT 20,
    is_private BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT fk_study_rooms_owner FOREIGN KEY (owner_id) REFERENCES users(id)
);

-- 创建自习室表索引
CREATE INDEX IF NOT EXISTS idx_study_rooms_owner_id ON study_rooms(owner_id);
CREATE INDEX IF NOT EXISTS idx_study_rooms_share_link ON study_rooms(share_link);
CREATE INDEX IF NOT EXISTS idx_study_rooms_expires_at ON study_rooms(expires_at);

-- 创建自习室成员表
CREATE TABLE IF NOT EXISTS room_members (
    id SERIAL PRIMARY KEY,
    room_id INTEGER NOT NULL REFERENCES study_rooms(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    is_anonymous BOOLEAN NOT NULL DEFAULT false,
    joined_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_room_members_room FOREIGN KEY (room_id) REFERENCES study_rooms(id),
    CONSTRAINT fk_room_members_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT unique_room_member UNIQUE (room_id, user_id)
);

-- 创建自习室成员表索引
CREATE INDEX IF NOT EXISTS idx_room_members_room_id ON room_members(room_id);
CREATE INDEX IF NOT EXISTS idx_room_members_user_id ON room_members(user_id);

-- 创建好友关系表
CREATE TABLE IF NOT EXISTS friends (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    friend_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',  -- pending, accepted, rejected
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_friendship UNIQUE (user_id, friend_id),
    CONSTRAINT fk_friends_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_friends_friend FOREIGN KEY (friend_id) REFERENCES users(id)
);

-- 创建好友关系表索引
CREATE INDEX IF NOT EXISTS idx_friends_user_id ON friends(user_id);
CREATE INDEX IF NOT EXISTS idx_friends_friend_id ON friends(friend_id);
CREATE INDEX IF NOT EXISTS idx_friends_status ON friends(status);

-- 创建好友契约表
CREATE TABLE IF NOT EXISTS friend_contracts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    friend_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    contract_type VARCHAR(50) NOT NULL,  -- study_buddy, accountability_partner, etc.
    contract_terms TEXT,
    status VARCHAR(20) NOT NULL DEFAULT 'active',  -- active, terminated
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_contract UNIQUE (user_id, friend_id),
    CONSTRAINT fk_contracts_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_contracts_friend FOREIGN KEY (friend_id) REFERENCES users(id)
);

-- 创建好友契约表索引
CREATE INDEX IF NOT EXISTS idx_contracts_user_id ON friend_contracts(user_id);
CREATE INDEX IF NOT EXISTS idx_contracts_friend_id ON friend_contracts(friend_id);
CREATE INDEX IF NOT EXISTS idx_contracts_status ON friend_contracts(status);

-- 创建好友消息表
CREATE TABLE IF NOT EXISTS friend_messages (
    id SERIAL PRIMARY KEY,
    sender_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    receiver_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    message_type VARCHAR(20) NOT NULL DEFAULT 'text',  -- text, image, etc.
    content TEXT NOT NULL,
    is_read BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_messages_sender FOREIGN KEY (sender_id) REFERENCES users(id),
    CONSTRAINT fk_messages_receiver FOREIGN KEY (receiver_id) REFERENCES users(id)
);

-- 创建好友消息表索引
CREATE INDEX IF NOT EXISTS idx_messages_sender_id ON friend_messages(sender_id);
CREATE INDEX IF NOT EXISTS idx_messages_receiver_id ON friend_messages(receiver_id);
CREATE INDEX IF NOT EXISTS idx_messages_created_at ON friend_messages(created_at);