import 'package:flutter/foundation.dart';
import 'package:flutter/services.dart';

import 'flutter_usb_go_platform_interface.dart';

/// An implementation of [FlutterUsbGoPlatform] that uses method channels.
class MethodChannelFlutterUsbGo extends FlutterUsbGoPlatform {
  /// The method channel used to interact with the native platform.
  @visibleForTesting
  final methodChannel = const MethodChannel('flutter_usb_go');

  @override
  Future<String?> getPlatformVersion() async {
    final version =
        await methodChannel.invokeMethod<String>('getPlatformVersion');
    return version;
  }

  @override
  Future<String?> getUsbInfo() async {
    final version = await methodChannel.invokeMethod<String>('getUsbInfo');
    return version;
  }

  @override
  Future<int?> openDevice() async {
    final version = await methodChannel.invokeMethod<int>('openDevice');
    return version;
  }

  @override
  Future<bool?> closeDevice() async {
    return await methodChannel.invokeMethod<bool>('closeDevice');
  }

  @override
  Future<Uint8List?> read() async {
    final version = await methodChannel.invokeMethod<Uint8List>('read');
    return version;
  }

  @override
  Future<int?> write() async {
    final version = await methodChannel.invokeMethod<int>('write');
    return version;
  }

  @override
  Future<Uint8List?> controlRead() async {
    final version = await methodChannel.invokeMethod<Uint8List>('controlRead');
    return version;
  }

  @override
  Future<int?> controlWrite() async {
    final version = await methodChannel.invokeMethod<int>('controlWrite');
    return version;
  }
}
