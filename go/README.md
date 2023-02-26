# flutter_usb_go

This Go package implements the host-side of the Flutter [flutter_usb_go](https://github.com/meilab/flutter_usb_go) plugin.

## Usage

Import as:

```go
import flutter_usb_go "github.com/meilab/flutter_usb_go/go"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&flutter_usb_go.FlutterUsbGoPlugin{}),
```
