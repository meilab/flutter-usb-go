// You have generated a new plugin project without specifying the `--platforms`
// flag. A plugin project with no platform support was generated. To add a
// platform, run `flutter create -t plugin --platforms <platforms> .` under the
// same directory. You can also find a detailed instruction on how to add
// platforms in the `pubspec.yaml` at
// https://flutter.dev/docs/development/packages-and-plugins/developing-packages#plugin-platforms.

import 'dart:typed_data';

import 'flutter_usb_go_platform_interface.dart';

class FlutterUsbGo {
  Future<String?> getPlatformVersion() {
    return FlutterUsbGoPlatform.instance.getPlatformVersion();
  }

  Future<String?> getUsbInfo() {
    return FlutterUsbGoPlatform.instance.getUsbInfo();
  }

  Future<int?> openDevice(Map<String, int> arguments) {
    return FlutterUsbGoPlatform.instance.openDevice(arguments);
  }

  Future<bool?> closeDevice() {
    return FlutterUsbGoPlatform.instance.closeDevice();
  }

  Future<Uint8List?> read(int argument) {
    return FlutterUsbGoPlatform.instance.read(argument);
  }

  Future<int?> write(Uint8List argument) {
    return FlutterUsbGoPlatform.instance.write(argument);
  }

  Future<Uint8List?> controlRead(Map<String, int> arguments) {
    return FlutterUsbGoPlatform.instance.controlRead(arguments);
  }

  Future<int?> controlWrite(Map<String, dynamic> arguments) {
    return FlutterUsbGoPlatform.instance.controlWrite(arguments);
  }
}
