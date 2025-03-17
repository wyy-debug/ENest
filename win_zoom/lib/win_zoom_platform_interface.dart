import 'package:plugin_platform_interface/plugin_platform_interface.dart';

import 'win_zoom_method_channel.dart';

abstract class WinZoomPlatform extends PlatformInterface {
  /// Constructs a WinZoomPlatform.
  WinZoomPlatform() : super(token: _token);

  static final Object _token = Object();

  static WinZoomPlatform _instance = MethodChannelWinZoom();

  /// The default instance of [WinZoomPlatform] to use.
  ///
  /// Defaults to [MethodChannelWinZoom].
  static WinZoomPlatform get instance => _instance;

  /// Platform-specific implementations should set this with their own
  /// platform-specific class that extends [WinZoomPlatform] when
  /// they register themselves.
  static set instance(WinZoomPlatform instance) {
    PlatformInterface.verifyToken(instance, _token);
    _instance = instance;
  }

  Future<String?> getPlatformVersion() {
    throw UnimplementedError('platformVersion() has not been implemented.');
  }

  Future<bool> initSDK() {
    throw UnimplementedError('initSDK() has not been implemented.');
  }

  Future<bool> joinSession({
    required String sessionName,
    required String sessionPassword,
    required String userName,
  }) {
    throw UnimplementedError('joinSession() has not been implemented.');
  }

  Future<bool> leaveSession() {
    throw UnimplementedError('leaveSession() has not been implemented.');
  }
}
