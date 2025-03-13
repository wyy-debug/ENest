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

# 服务名称
SERVICE_NAME="estudyroom"

# 检查服务状态
check_status() {
    if systemctl is-active --quiet ${SERVICE_NAME}; then
        return 0
    else
        return 1
    fi
}

# 启动服务
start_service() {
    if check_status; then
        log_warn "服务已经在运行中"
    else
        log_info "正在启动服务..."
        systemctl start ${SERVICE_NAME}
        if check_status; then
            log_info "服务启动成功"
        else
            log_error "服务启动失败"
            exit 1
        fi
    fi
}

# 停止服务
stop_service() {
    if check_status; then
        log_info "正在停止服务..."
        systemctl stop ${SERVICE_NAME}
        if ! check_status; then
            log_info "服务已停止"
        else
            log_error "服务停止失败"
            exit 1
        fi
    else
        log_warn "服务未在运行"
    fi
}

# 重启服务
restart_service() {
    log_info "正在重启服务..."
    systemctl restart ${SERVICE_NAME}
    if check_status; then
        log_info "服务重启成功"
    else
        log_error "服务重启失败"
        exit 1
    fi
}

# 查看服务状态
show_status() {
    systemctl status ${SERVICE_NAME}
}

# 命令行参数处理
case "$1" in
    start)
        start_service
        ;;
    stop)
        stop_service
        ;;
    restart)
        restart_service
        ;;
    status)
        show_status
        ;;
    *)
        echo "用法: $0 {start|stop|restart|status}"
        exit 1
        ;;
esac

exit 0