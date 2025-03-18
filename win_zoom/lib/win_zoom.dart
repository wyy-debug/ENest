
import 'win_zoom_platform_interface.dart';

class WinZoom {
  Future<String?> getPlatformVersion() {
    return WinZoomPlatform.instance.getPlatformVersion();
  }

  Future<String?> initSDK() {
    return WinZoomPlatform.instance.initSDK();
  }

  Future<bool> joinSession({
    required String sessionName,
    required String sessionPassword,
    required String userName,
  }) {
    return WinZoomPlatform.instance.joinSession(
      sessionName: sessionName,
      sessionPassword: sessionPassword,
      userName: userName,
    );
  }

  Future<bool> leaveSession() {
    return WinZoomPlatform.instance.leaveSession();
  }
}
