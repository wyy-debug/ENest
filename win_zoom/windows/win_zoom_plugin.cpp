#include <windows.h>

#include "win_zoom_plugin.h"



// For getPlatformVersion; remove unless needed for your plugin implementation.
#include <VersionHelpers.h>

#include <flutter/method_channel.h>
#include <flutter/plugin_registrar_windows.h>
#include <flutter/standard_method_codec.h>
#include "include/win_zoom/MainZoom.h"
#include <memory>
#include <sstream>


namespace win_zoom {


// API

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

WinZoomPlugin::WinZoomPlugin() {
}

void WinZoomPlugin::initSDK(std::unique_ptr<flutter::MethodResult<flutter::EncodableValue>> result){
  MainZoom zoom = MainZoom::GetInstance();
  zoom.InitVideoSDK();
  zoom.JoinSession();
  result->Error("aaa","aaa");
}

void WinZoomPlugin::joinSession(const flutter::MethodCall<flutter::EncodableValue> &method_call, std::unique_ptr<flutter::MethodResult<flutter::EncodableValue>> result) {
  MainZoom zoom = MainZoom::GetInstance();
  zoom.JoinSession();
  result->Success();
}

WinZoomPlugin::~WinZoomPlugin() {}


void WinZoomPlugin::leaveSession(std::unique_ptr<flutter::MethodResult<flutter::EncodableValue>> result) {
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
