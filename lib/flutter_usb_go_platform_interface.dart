import 'dart:typed_data';

import 'package:plugin_platform_interface/plugin_platform_interface.dart';

import 'flutter_usb_go_method_channel.dart';

abstract class FlutterUsbGoPlatform extends PlatformInterface {
  /// Constructs a FlutterUsbGoPlatform.
  FlutterUsbGoPlatform() : super(token: _token);

  static final Object _token = Object();

  static FlutterUsbGoPlatform _instance = MethodChannelFlutterUsbGo();

  /// The default instance of [FlutterUsbGoPlatform] to use.
  ///
  /// Defaults to [MethodChannelFlutterUsbGo].
  static FlutterUsbGoPlatform get instance => _instance;

  /// Platform-specific implementations should set this with their own
  /// platform-specific class that extends [FlutterUsbGoPlatform] when
  /// they register themselves.
  static set instance(FlutterUsbGoPlatform instance) {
    PlatformInterface.verifyToken(instance, _token);
    _instance = instance;
  }

  Future<String?> getPlatformVersion() {
    throw UnimplementedError('platformVersion() has not been implemented.');
  }

  Future<String?> getUsbInfo() {
    throw UnimplementedError('platformVersion() has not been implemented.');
  }

  Future<int?> openDevice() {
    throw UnimplementedError('platformVersion() has not been implemented.');
  }

  Future<bool?> closeDevice() {
    throw UnimplementedError('platformVersion() has not been implemented.');
  }

  Future<Uint8List?> read() {
    throw UnimplementedError('platformVersion() has not been implemented.');
  }

  Future<int?> write() {
    throw UnimplementedError('platformVersion() has not been implemented.');
  }

  Future<Uint8List?> controlRead() {
    throw UnimplementedError('platformVersion() has not been implemented.');
  }

  Future<int?> controlWrite() {
    throw UnimplementedError('platformVersion() has not been implemented.');
  }
}
