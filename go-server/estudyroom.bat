@echo off
setlocal enabledelayedexpansion

:: 设置颜色代码
for /f %%a in ('echo prompt $E ^| cmd') do set "ESC=%%a"

:: 设置服务名称和描述
set "SERVICE_NAME=E-StudyRoom"
set "DISPLAY_NAME=E-StudyRoom API Server"
set "DESCRIPTION=E-StudyRoom API Server for study room management"
set "BINARY_PATH=%~dp0estudyroom.exe"

:: 日志函数
:log_info
echo %ESC%[32m[INFO]%ESC%[0m %~1
goto :eof

:log_warn
echo %ESC%[33m[WARN]%ESC%[0m %~1
goto :eof

:log_error
echo %ESC%[31m[ERROR]%ESC%[0m %~1
goto :eof

:: 检查管理员权限
net session >nul 2>&1
if %errorLevel% neq 0 (
    call :log_error "请以管理员权限运行此脚本"
    goto :eof
)

:: 检查Go环境
go version >nul 2>&1
if %errorLevel% neq 0 (
    call :log_error "未检测到Go环境，请先安装Go 1.21或更高版本"
    goto :eof
)

:: 检查PostgreSQL
psql --version >nul 2>&1
if %errorLevel% neq 0 (
    call :log_error "未检测到PostgreSQL，请先安装PostgreSQL"
    goto :eof
)

:: 编译应用
if not exist "%BINARY_PATH%" (
    call :log_info "正在编译应用..."
    go build -o estudyroom.exe
    if %errorLevel% neq 0 (
        call :log_error "编译应用失败"
        goto :eof
    )
)

:: 处理命令行参数
if "%1"=="install" (
    call :install_service
) else if "%1"=="uninstall" (
    call :uninstall_service
) else if "%1"=="start" (
    call :start_service
) else if "%1"=="stop" (
    call :stop_service
) else if "%1"=="restart" (
    call :restart_service
) else if "%1"=="status" (
    call :check_status
) else (
    call :show_usage
)

goto :eof

:install_service
call :log_info "正在安装服务..."
sc create %SERVICE_NAME% binPath= "%BINARY_PATH%" DisplayName= "%DISPLAY_NAME%" start= auto
sc description %SERVICE_NAME% "%DESCRIPTION%"
if %errorLevel% neq 0 (
    call :log_error "安装服务失败"
    goto :eof
)
call :log_info "服务安装成功"
goto :eof

:uninstall_service
call :log_info "正在卸载服务..."
sc stop %SERVICE_NAME% >nul 2>&1
sc delete %SERVICE_NAME%
if %errorLevel% neq 0 (
    call :log_error "卸载服务失败"
    goto :eof
)
call :log_info "服务卸载成功"
goto :eof

:start_service
call :log_info "正在启动服务..."
sc start %SERVICE_NAME%
if %errorLevel% neq 0 (
    call :log_error "启动服务失败"
    goto :eof
)
call :log_info "服务启动成功"
goto :eof

:stop_service
call :log_info "正在停止服务..."
sc stop %SERVICE_NAME%
if %errorLevel% neq 0 (
    call :log_error "停止服务失败"
    goto :eof
)
call :log_info "服务停止成功"
goto :eof

:restart_service
call :stop_service
timeout /t 2 /nobreak >nul
call :start_service
goto :eof

:check_status
sc query %SERVICE_NAME%
goto :eof

:show_usage
echo 使用方法：
echo %~nx0 ^<command^>
echo.
echo 可用命令：
echo   install    - 安装服务
echo   uninstall  - 卸载服务
echo   start      - 启动服务
echo   stop       - 停止服务
echo   restart    - 重启服务
echo   status     - 查看服务状态
goto :eof