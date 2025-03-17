#include "include/win_zoom/win_zoom_plugin_c_api.h"

#include <flutter/plugin_registrar_windows.h>

#include "win_zoom_plugin.h"

void WinZoomPluginCApiRegisterWithRegistrar(
    FlutterDesktopPluginRegistrarRef registrar) {
  win_zoom::WinZoomPlugin::RegisterWithRegistrar(
      flutter::PluginRegistrarManager::GetInstance()
          ->GetRegistrar<flutter::PluginRegistrarWindows>(registrar));
}
