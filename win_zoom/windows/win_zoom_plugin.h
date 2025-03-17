#ifndef FLUTTER_PLUGIN_WIN_ZOOM_PLUGIN_H_
#define FLUTTER_PLUGIN_WIN_ZOOM_PLUGIN_H_

#include <flutter/method_channel.h>
#include <flutter/plugin_registrar_windows.h>

#include <memory>

namespace win_zoom {

    class WinZoomPlugin : public flutter::Plugin {
    public:
        static void RegisterWithRegistrar(flutter::PluginRegistrarWindows *registrar);

        WinZoomPlugin();
        
        virtual ~WinZoomPlugin();
        
        // Disallow copy and assign.
        WinZoomPlugin(const WinZoomPlugin&) = delete;
        WinZoomPlugin& operator=(const WinZoomPlugin&) = delete;
        
        // Called when a method is called on this plugin's channel from Dart.
        void HandleMethodCall(
            const flutter::MethodCall<flutter::EncodableValue> &method_call,
            std::unique_ptr<flutter::MethodResult<flutter::EncodableValue>> result);


        // Logic code
        void initSDK(std::unique_ptr<flutter::MethodResult<flutter::EncodableValue>> result);
        void joinSession(const flutter::MethodCall<flutter::EncodableValue> &method_call, std::unique_ptr<flutter::MethodResult<flutter::EncodableValue>> result);
        void leaveSession(std::unique_ptr<flutter::MethodResult<flutter::EncodableValue>> result);
    };

}  // namespace win_zoom

#endif  // FLUTTER_PLUGIN_WIN_ZOOM_PLUGIN_H_
