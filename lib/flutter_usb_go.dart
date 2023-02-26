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

  Future<int?> openDevice() {
    return FlutterUsbGoPlatform.instance.openDevice();
  }

  Future<bool?> closeDevice() {
    return FlutterUsbGoPlatform.instance.closeDevice();
  }

  Future<Uint8List?> read() {
    return FlutterUsbGoPlatform.instance.read();
  }

  Future<int?> write() {
    return FlutterUsbGoPlatform.instance.write();
  }

  Future<Uint8List?> controlRead() {
    return FlutterUsbGoPlatform.instance.controlRead();
  }

  Future<int?> controlWrite() {
    return FlutterUsbGoPlatform.instance.controlWrite();
  }
}
