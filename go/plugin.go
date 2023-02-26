package flutter_usb_go

import (
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

const channelName = "flutter_usb_go"

// FlutterUsbGoPlugin implements flutter.Plugin and handles method.
type FlutterUsbGoPlugin struct{}

var _ flutter.Plugin = &FlutterUsbGoPlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *FlutterUsbGoPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getPlatformVersion", p.handlePlatformVersion)
	return nil
}

func (p *FlutterUsbGoPlugin) handlePlatformVersion(arguments interface{}) (reply interface{}, err error) {
	return "go-flutter " + flutter.PlatformVersion, nil
}
