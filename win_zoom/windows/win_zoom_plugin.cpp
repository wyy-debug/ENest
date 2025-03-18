#include <windows.h>

#include "win_zoom_plugin.h"



// For getPlatformVersion; remove unless needed for your plugin implementation.
#include <VersionHelpers.h>

#include <flutter/method_channel.h>
#include <flutter/plugin_registrar_windows.h>
#include <flutter/standard_method_codec.h>

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

WinZoomPlugin::WinZoomPlugin() {}

WinZoomPlugin::~WinZoomPlugin() 
{
    UninitVideoSDK();
}

void WinZoomPlugin::InitVideoSDK()
{
    ZoomVideoSDKInitParams init_params;
    init_params.domain = _T("https://go.zoom.us");
    init_params.enableLog = true;
    init_params.logFilePrefix = _T("zoom_win_video_demo");
    init_params.videoRawDataMemoryMode = ZoomVideoSDKRawDataMemoryModeHeap;
    init_params.shareRawDataMemoryMode = ZoomVideoSDKRawDataMemoryModeHeap;
    init_params.audioRawDataMemoryMode = ZoomVideoSDKRawDataMemoryModeHeap;
    init_params.enableIndirectRawdata = false;

    //ZoomVideoSDKMgr::GetInst().Init(this, init_params);
}

void WinZoomPlugin::UninitVideoSDK()
{
  //ZoomVideoSDKMgr::GetInst().UnInit();
}

void WinZoomPlugin::initSDK(std::unique_ptr<flutter::MethodResult<flutter::EncodableValue>> result){
}

void WinZoomPlugin::joinSession(const flutter::MethodCall<flutter::EncodableValue> &method_call, std::unique_ptr<flutter::MethodResult<flutter::EncodableValue>> result) {
}

void WinZoomPlugin::leaveSession(std::unique_ptr<flutter::MethodResult<flutter::EncodableValue>> result) {
}

void WinZoomPlugin::JoinSession()
{
  // test code
  //  ZoomVideoSDKSessionContext sessionContext;
  //  sessionContext.sessionName = L"Session name";
  //  sessionContext.sessionPassword = L"Session password";
  //  sessionContext.userName = L"User name";
//
  //  sessionContext.token = L"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfa2V5IjoiSEppSnRJMjAwWEIxcDRhQ3NKV1h5Y1VzNWJMZDR4Q0NjYW9TIiwicm9sZV90eXBlIjoxLCJ0cGMiOiJURVNUIiwiaWF0IjoxNzQyMjc4Mzk2LCJleHAiOjE3NDIyODE5OTZ9.uDcFy8o4SsqTowoVbUcxqpa7MVa-QsTON8_b7stGAx0";
  //  sessionContext.videoOption.localVideoOn = true;
  //  sessionContext.audioOption.connect = true;
  //  sessionContext.audioOption.mute = false;
//
  //  IZoomVideoSDKSession* pSession = ZoomVideoSDKMgr::GetInst().JoinSession(sessionContext);
  //  IZoomVideoSDKUser* pUser;
  //  pUser = ZoomVideoSDKMgr::GetInst().GetMySelf();
//
//
  //  ZoomVideoSDKResolution resolution = ZoomVideoSDKResolution_360P;
//
  //  IZoomVideoSDKRawDataPipe* pPipe = NULL;
  //  pPipe = pUser->GetVideoPipe();
  //  if(!pPipe) return;

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
