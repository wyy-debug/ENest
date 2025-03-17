import 'package:flutter_test/flutter_test.dart';
import 'package:win_zoom/win_zoom.dart';
import 'package:win_zoom/win_zoom_platform_interface.dart';
import 'package:win_zoom/win_zoom_method_channel.dart';
import 'package:plugin_platform_interface/plugin_platform_interface.dart';

class MockWinZoomPlatform
    with MockPlatformInterfaceMixin
    implements WinZoomPlatform {

  @override
  Future<String?> getPlatformVersion() => Future.value('42');
}

void main() {
  final WinZoomPlatform initialPlatform = WinZoomPlatform.instance;

  test('$MethodChannelWinZoom is the default instance', () {
    expect(initialPlatform, isInstanceOf<MethodChannelWinZoom>());
  });

  test('getPlatformVersion', () async {
    WinZoom winZoomPlugin = WinZoom();
    MockWinZoomPlatform fakePlatform = MockWinZoomPlatform();
    WinZoomPlatform.instance = fakePlatform;

    expect(await winZoomPlugin.getPlatformVersion(), '42');
  });
}
