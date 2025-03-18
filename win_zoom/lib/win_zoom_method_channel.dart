import 'package:flutter/foundation.dart';
import 'package:flutter/services.dart';

import 'win_zoom_platform_interface.dart';

/// An implementation of [WinZoomPlatform] that uses method channels.
class MethodChannelWinZoom extends WinZoomPlatform {
  /// The method channel used to interact with the native platform.
  @visibleForTesting
  final methodChannel = const MethodChannel('win_zoom');

  @override
  Future<String?> getPlatformVersion() async {
    final version = await methodChannel.invokeMethod<String>('getPlatformVersion');
    return version;
  }

  @override
  Future<String?> initSDK() async {
    final version = await methodChannel.invokeMethod<String>('initSDK');
    return version;
  }

  @override
  Future<bool> joinSession({
    required String sessionName,
    required String sessionPassword,
    required String userName,
  }) async {
    try {
      final result = await methodChannel.invokeMethod<bool>('joinSession', {
        'sessionName': sessionName,
        'sessionPassword': sessionPassword,
        'userName': userName,
      });
      return result ?? false;
    } catch (e) {
      debugPrint('Failed to join session: $e');
      return false;
    }
  }

  @override
  Future<bool> leaveSession() async {
    try {
      final result = await methodChannel.invokeMethod<bool>('leaveSession');
      return result ?? false;
    } catch (e) {
      debugPrint('Failed to leave session: $e');
      return false;
    }
  }
}
