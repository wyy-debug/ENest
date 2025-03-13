#!/bin/bash

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 日志函数
log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 检查是否以root权限运行
if [ "$(id -u)" != "0" ]; then
    log_error "请使用root权限运行此脚本"
    exit 1

fi

# 检查系统要求
if ! command -v go &> /dev/null; then
    log_error "未检测到Go环境，请先安装Go 1.21或更高版本"
    exit 1
fi

if ! command -v psql &> /dev/null; then
    log_error "未检测到PostgreSQL，请先安装PostgreSQL"
    exit 1
fi

# 创建服务用户
log_info "创建服务用户..."
id -u estudyroom &>/dev/null || useradd -r -s /bin/false estudyroom

# 创建必要的目录
log_info "创建必要的目录..."
install_dir="/opt/estudyroom"
mkdir -p "${install_dir}"
chown estudyroom:estudyroom "${install_dir}"

# 复制应用文件
log_info "复制应用文件..."
cp -r ./* "${install_dir}/"
chown -R estudyroom:estudyroom "${install_dir}"

# 初始化数据库
log_info "初始化数据库..."
sudo -u postgres psql -c "CREATE USER postgres WITH PASSWORD 'postgres';"
sudo -u postgres psql -c "CREATE DATABASE estudyroom OWNER postgres;"
sudo -u postgres psql -d estudyroom -f "${install_dir}/database/init.sql"

# 创建配置文件
log_info "创建配置文件..."
cat > "${install_dir}/config/database.go" << EOL
package config

type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
    SSLMode  string
}

func GetDatabaseConfig() *DatabaseConfig {
    return &DatabaseConfig{
        Host:     "localhost",
        Port:     5432,
        User:     "estudyroom",
        Password: "estudyroom",
        DBName:   "estudyroom",
        SSLMode:  "disable",
    }
}
EOL

# 编译应用
log_info "编译应用..."
cd "${install_dir}"
go build -o estudyroom

# 创建systemd服务文件
log_info "创建systemd服务文件..."
cat > /etc/systemd/system/estudyroom.service << EOL
[Unit]
Description=E-StudyRoom API Server
After=network.target postgresql.service

[Service]
Type=simple
User=estudyroom
WorkingDirectory=${install_dir}
ExecStart=${install_dir}/estudyroom
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
EOL

# 重新加载systemd配置
log_info "重新加载systemd配置..."
systemctl daemon-reload

log_info "安装完成！"
log_info "使用以下命令管理服务："
echo -e "启动服务：${GREEN}systemctl start estudyroom${NC}"
echo -e "停止服务：${GREEN}systemctl stop estudyroom${NC}"
echo -e "查看状态：${GREEN}systemctl status estudyroom${NC}"
echo -e "设置开机启动：${GREEN}systemctl enable estudyroom${NC}"