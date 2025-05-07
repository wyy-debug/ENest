-- 记录迁移版本
INSERT INTO migrations (version) VALUES ('002_sample_data');

-- 创建示例用户
-- 密码都是 'password123'，使用bcrypt进行哈希处理
INSERT INTO users (username, email, password_hash, avatar, signature, study_direction, total_study_time, achievement_points)
VALUES 
    ('test_user', 'test@example.com', '$2a$10$LmhJeYeGCj/rCFJ2q1CtWeMQgbepthCKHO4GxLEVX6dUgT4vc4NL.', 'https://api.dicebear.com/7.x/avataaars/svg?seed=test_user', '每天进步一点点', '计算机科学', 1200, 50),
    ('study_buddy', 'buddy@example.com', '$2a$10$LmhJeYeGCj/rCFJ2q1CtWeMQgbepthCKHO4GxLEVX6dUgT4vc4NL.', 'https://api.dicebear.com/7.x/avataaars/svg?seed=study_buddy', '一起学习吧！', '数学', 900, 30),
    ('focusmaster', 'focus@example.com', '$2a$10$LmhJeYeGCj/rCFJ2q1CtWeMQgbepthCKHO4GxLEVX6dUgT4vc4NL.', 'https://api.dicebear.com/7.x/avataaars/svg?seed=focusmaster', '专注力就是生产力', '物理', 2400, 80);

-- 创建示例自习室
INSERT INTO study_rooms (owner_id, name, description, share_link, max_members, is_private, theme, background_image, expires_at)
VALUES 
    (1, '编程自习室', '专注于编程学习的自习室', 'code_room_123', 20, false, 'coding', 'https://images.unsplash.com/photo-1605379399642-870262d3d051', CURRENT_TIMESTAMP + INTERVAL '7 days'),
    (2, '数学研讨室', '一起解决数学难题', 'math_room_456', 10, false, 'math', 'https://images.unsplash.com/photo-1635070041078-e363dbe005cb', CURRENT_TIMESTAMP + INTERVAL '7 days'),
    (3, '专注学习间', '安静的学习环境', 'focus_room_789', 5, true, 'minimal', 'https://images.unsplash.com/photo-1497032628192-86f99bcd76bc', CURRENT_TIMESTAMP + INTERVAL '7 days');

-- 添加用户到自习室
INSERT INTO room_members (room_id, user_id, is_anonymous, role, status)
VALUES 
    (1, 1, false, 'owner', 'online'),
    (1, 2, false, 'member', 'online'),
    (2, 2, false, 'owner', 'online'),
    (2, 3, false, 'member', 'away'),
    (3, 3, false, 'owner', 'online'),
    (3, 1, true, 'member', 'online');

-- 创建好友关系
INSERT INTO friends (user_id, friend_id, status)
VALUES 
    (1, 2, 'accepted'),
    (1, 3, 'accepted'),
    (2, 3, 'pending');

-- 创建好友契约
INSERT INTO friend_contracts (user_id, friend_id, contract_type, contract_terms, start_date, end_date, goal_type, goal_value, status)
VALUES 
    (1, 2, 'study_buddy', '每天至少学习2小时', CURRENT_DATE, CURRENT_DATE + INTERVAL '30 days', 'daily_time', 120, 'active');

-- 创建好友消息
INSERT INTO friend_messages (sender_id, receiver_id, message_type, content, is_read)
VALUES 
    (1, 2, 'text', '嗨，今天要一起学习吗？', true),
    (2, 1, 'text', '好的，3点开始吧！', true),
    (3, 1, 'text', '你解决那个编程问题了吗？', false);

-- 创建学习记录
INSERT INTO study_records (user_id, room_id, start_time, end_time, duration, interruptions, notes)
VALUES 
    (1, 1, CURRENT_TIMESTAMP - INTERVAL '1 day', CURRENT_TIMESTAMP - INTERVAL '22 hours', 120, 2, '学习了React Hooks'),
    (2, 2, CURRENT_TIMESTAMP - INTERVAL '2 days', CURRENT_TIMESTAMP - INTERVAL '1 day 23 hours', 60, 1, '复习了线性代数'),
    (3, 3, CURRENT_TIMESTAMP - INTERVAL '12 hours', CURRENT_TIMESTAMP - INTERVAL '10 hours', 120, 0, '学习了量子力学基础');

-- 创建成就
INSERT INTO achievements (name, description, icon, points, requirement_type, requirement_value)
VALUES 
    ('学习新手', '累计学习时间达到10小时', 'achievement_rookie', 10, 'total_study_time', 600),
    ('专注大师', '连续学习2小时无中断', 'achievement_focus', 20, 'continuous_focus', 120),
    ('学习达人', '累计学习时间达到50小时', 'achievement_expert', 50, 'total_study_time', 3000);

-- 分配成就给用户
INSERT INTO user_achievements (user_id, achievement_id)
VALUES 
    (1, 1),
    (2, 1),
    (3, 1),
    (3, 2);

-- 创建通知
INSERT INTO notifications (user_id, type, title, content, is_read, related_entity_type, related_entity_id)
VALUES 
    (1, 'friend_request', '新的好友请求', 'focusmaster想添加你为好友', true, 'friend', 3),
    (2, 'achievement', '获得新成就', '恭喜你获得了"学习新手"成就！', false, 'achievement', 1),
    (3, 'contract', '学习契约邀请', 'test_user邀请你加入学习契约', false, 'contract', 1); 