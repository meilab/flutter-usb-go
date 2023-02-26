import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_usb_go/flutter_usb_go.dart';
import 'package:flutter_usb_go/flutter_usb_go_platform_interface.dart';
import 'package:flutter_usb_go/flutter_usb_go_method_channel.dart';
import 'package:plugin_platform_interface/plugin_platform_interface.dart';

class MockFlutterUsbGoPlatform
    with MockPlatformInterfaceMixin
    implements FlutterUsbGoPlatform {

  @override
  Future<String?> getPlatformVersion() => Future.value('42');
}

void main() {
  final FlutterUsbGoPlatform initialPlatform = FlutterUsbGoPlatform.instance;

  test('$MethodChannelFlutterUsbGo is the default instance', () {
    expect(initialPlatform, isInstanceOf<MethodChannelFlutterUsbGo>());
  });

  test('getPlatformVersion', () async {
    FlutterUsbGo flutterUsbGoPlugin = FlutterUsbGo();
    MockFlutterUsbGoPlatform fakePlatform = MockFlutterUsbGoPlatform();
    FlutterUsbGoPlatform.instance = fakePlatform;

    expect(await flutterUsbGoPlugin.getPlatformVersion(), '42');
  });
}
