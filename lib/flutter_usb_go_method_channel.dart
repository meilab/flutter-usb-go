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
    final version = await methodChannel.invokeMethod<String>('getPlatformVersion');
    return version;
  }
}
