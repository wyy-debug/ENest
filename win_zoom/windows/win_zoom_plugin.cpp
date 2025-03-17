#include <windows.h>

#include "win_zoom_plugin.h"
#include "include/zoom_video_sdk_api.h"
#include "include/zoom_video_sdk_interface.h"
#include "include/zoom_video_sdk_def.h"


// For getPlatformVersion; remove unless needed for your plugin implementation.
#include <VersionHelpers.h>

#include <flutter/method_channel.h>
#include <flutter/plugin_registrar_windows.h>
#include <flutter/standard_method_codec.h>

#include <memory>
#include <sstream>


USING_ZOOM_VIDEO_SDK_NAMESPACE

namespace win_zoom {

IZoomVideoSDK* video_sdk_ = nullptr;

void WinZoomPlugin::RegisterWithRegistrar(flutter::PluginRegistrarWindows *registrar) {
  auto channel = std::make_unique<flutter::MethodChannel<flutter::EncodableValue>>(
      registrar->messenger(), "win_zoom",
      &flutter::StandardMethodCodec::GetInstance());

  auto plugin = std::make_unique<WinZoomPlugin>();
  channel->SetMethodCallHandler(
      [plugin_pointer = plugin.get()](const auto &call, auto result) {
        plugin_pointer->HandleMethodCall(call, std::move(result));
      });
  registrar->AddPlugin(std::move(plugin));
}

WinZoomPlugin::WinZoomPlugin() {}

WinZoomPlugin::~WinZoomPlugin() {
  if (video_sdk_) {
    video_sdk_->cleanup();
    video_sdk_ = nullptr;
  }
}

void WinZoomPlugin::initSDK(std::unique_ptr<flutter::MethodResult<flutter::EncodableValue>> result)
{
  if (video_sdk_ == nullptr) {
    video_sdk_ = CreateZoomVideoSDKObj();
    if (video_sdk_) {
      ZOOM_VIDEO_SDK_NAMESPACE::ZoomVideoSDKInitParams init_params;
      init_params.domain = L"https://zoom.us";
      init_params.enableIndirectRawdata = false;
      init_params.enableLog = true;
      init_params.logFilePrefix = L"prefix";
      init_params.audioRawDataMemoryMode = ZoomVideoSDKRawDataMemoryModeHeap;
      init_params.videoRawDataMemoryMode = ZoomVideoSDKRawDataMemoryModeHeap;
      init_params.shareRawDataMemoryMode = ZoomVideoSDKRawDataMemoryModeHeap;
      
      ZoomVideoSDKErrors err = video_sdk_->initialize(init_params);
      if (err == ZoomVideoSDKErrors_Success) {
        result->Success("init success");
        return;
      }
      else {
        std::string error_message = "初始化失败，错误代码: " + std::to_string(static_cast<int>(err));
        OutputDebugStringA(error_message.c_str());
        
        switch(err) {
          case ZoomVideoSDKErrors_Auth_Error:
            result->Error("AUTH_FAILED", "认证失败");
            break;
          case ZoomVideoSDKErrors_Uninitialize:
            result->Error("CONFIG_ERROR", "SDK配置错误");
            break;
          default:
            result->Error("UNKNOWN_ERROR", error_message);
        }
      }
    }
  }
}

void WinZoomPlugin::joinSession(const flutter::MethodCall<flutter::EncodableValue> &method_call, std::unique_ptr<flutter::MethodResult<flutter::EncodableValue>> result) {
    const auto* arguments = std::get_if<flutter::EncodableMap>(method_call.arguments());
    if (arguments && video_sdk_) {
      auto session_name_it = arguments->find(flutter::EncodableValue("sessionName"));
      auto session_password_it = arguments->find(flutter::EncodableValue("sessionPassword"));
      auto user_name_it = arguments->find(flutter::EncodableValue("userName"));
      
      if (session_name_it != arguments->end() && 
          session_password_it != arguments->end() && 
          user_name_it != arguments->end()) {
        const std::string& session_name = std::get<std::string>(session_name_it->second);
        const std::string& session_password = std::get<std::string>(session_password_it->second);
        const std::string& user_name = std::get<std::string>(user_name_it->second);
        
        ZOOM_VIDEO_SDK_NAMESPACE::ZoomVideoSDKSessionContext session_context;
        std::wstring wsession_name(session_name.begin(), session_name.end());
        std::wstring wsession_password(session_password.begin(), session_password.end());
        std::wstring wuser_name(user_name.begin(), user_name.end());
        session_context.sessionName = wsession_name.c_str();
        session_context.sessionPassword = wsession_password.c_str();
        session_context.userName = wuser_name.c_str();
        session_context.audioOption.connect = true;
        session_context.videoOption.localVideoOn = true;
        
        IZoomVideoSDKSession* session = video_sdk_->joinSession(session_context);
        if (session != nullptr) {
          result->Success(flutter::EncodableValue(true));
          return;
        }
        result->Error("JOIN_FAILED", "Failed to join session");
      } else {
        result->Error("INVALID_ARGUMENT", "Session name, password and user name are required");
      }
    } else {
      result->Error("SDK_NOT_INITIALIZED", "Zoom Video SDK is not initialized");
    }
}

void WinZoomPlugin::leaveSession(std::unique_ptr<flutter::MethodResult<flutter::EncodableValue>> result) {
  if (video_sdk_) {
    ZoomVideoSDKErrors err = video_sdk_->leaveSession(false);
    if (err == ZoomVideoSDKErrors_Success) {
      result->Success(flutter::EncodableValue(true));
      return;
    }
    result->Error("LEAVE_FAILED", "Failed to leave session");
  } else {
    result->Error("SDK_NOT_INITIALIZED", "Zoom Video SDK is not initialized");
  }
}


void WinZoomPlugin::HandleMethodCall(
    const flutter::MethodCall<flutter::EncodableValue> &method_call,
    std::unique_ptr<flutter::MethodResult<flutter::EncodableValue>> result) {
  const std::string& method_name = method_call.method_name();
  
  switch(method_name[0]) {
    case 'i':
      if (method_name == "initSDK") {
        initSDK(std::move(result));
        return;
      }
      break;
    case 'j':
      if (method_name == "joinSession") {
        joinSession(method_call, std::move(result));
        return;
      }
      break;
    case 'l':
      if (method_name == "leaveSession") {
        leaveSession(std::move(result));
        return;
      }
      break;
  }
  
  result->NotImplemented();
}

}  // namespace win_zoom
