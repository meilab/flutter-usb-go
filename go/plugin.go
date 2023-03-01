package flutter_usb_go

import (
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"

	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/gousb"
	"github.com/google/gousb/usbid"
)

const channelName = "flutter_usb_go"

var (
	debug = flag.Int("debug", 0, "libusb debug level (0..3)")
)

type UsbDevice struct {
	isActive bool
	// TODO: we need a status to check whether USB is opened
	// Fields for interacting with the USB connection
	context *gousb.Context
	device  *gousb.Device
	intf    *gousb.Interface
	epIn    *gousb.InEndpoint
	epOut   *gousb.OutEndpoint

	// Fields for managing async operations
	// waitGroup *sync.WaitGroup

	// Channels for error reporting and closing
	errors chan error
	done   func()
	// close  chan bool
}

var usbDevice UsbDevice

// FlutterUsbGoPlugin implements flutter.Plugin and handles method.
type FlutterUsbGoPlugin struct{}

var _ flutter.Plugin = &FlutterUsbGoPlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *FlutterUsbGoPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getPlatformVersion", p.handlePlatformVersion)
	channel.HandleFunc("getUsbInfo", handleGetUsbInfo)
	channel.HandleFunc("openDevice", handleOpenDevice)
	channel.HandleFunc("closeDevice", handleCloseDevice)
	channel.HandleFunc("read", handleRead)
	channel.HandleFunc("write", handleWrite)
	channel.HandleFunc("controlRead", handleControlRead)
	channel.HandleFunc("controlWrite", handleControlWrite)
	return nil
}

func (p *FlutterUsbGoPlugin) handlePlatformVersion(arguments interface{}) (reply interface{}, err error) {
	return "go-flutter " + flutter.PlatformVersion, nil
}

func handleOpenDevice(arguments interface{}) (reply interface{}, err error) {
	arg := arguments.(map[interface{}]interface{})
	vid := gousb.ID(arg["vid"].(int32))
	pid := gousb.ID(arg["pid"].(int32))

	// Initialize a new Context.
	ctx := gousb.NewContext()
	// TODO: do we need to close context here
	// defer ctx.Close()

	// Open any device with a given VID/PID using a convenience function.
	dev, err := ctx.OpenDeviceWithVIDPID(vid, pid)
	if err != nil {
		log.Fatalf("Could not open a device: %v", err)
		return "", err
	}
	// TODO: do we need to close it here or in the closeHandler
	// defer dev.Close()

	// Claim the default interface using a convenience function.
	// The default interface is always #0 alt #0 in the currently active
	// config.
	intf, done, err := dev.DefaultInterface()
	if err != nil {
		log.Fatalf("%s.DefaultInterface(): %v", dev, err)
		return "", err
	}
	// must call done here, otherwise we can not claim USB after we close it
	// defer done()

	// Open an OUT endpoint.
	epIn, err := intf.InEndpoint(1)
	if err != nil {
		log.Fatalf("%s.IutEndpoint(0): %v", intf, err)
		return "", err
	}

	epOut, err := intf.OutEndpoint(2)
	if err != nil {
		log.Fatalf("%s.OutEndpoint(1): %v", intf, err)
		return "", err
	}

	usbDevice = UsbDevice{
		isActive: true,
		context:  ctx,
		device:   dev,
		// waitGroup: &sync.WaitGroup{},
		intf:   intf,
		epIn:   epIn,
		epOut:  epOut,
		errors: make(chan error),
		// close:     make(chan bool),
		done: done,
	}

	// TODO: need to think what to return
	return arg["vid"], nil
	// return nil, fmt.Errorf("failed to obtain configuration for device")
}

func handleCloseDevice(arguments interface{}) (reply interface{}, err error) {
	// usbDevice.close <- true
	// usbDevice.waitGroup.Wait()

	if usbDevice.isActive == true {
		usbDevice.isActive = false
		usbDevice.epIn = nil
		usbDevice.epOut = nil
		usbDevice.intf.Close()
		usbDevice.device.Close()
		usbDevice.context.Close()
		// done need to be called here otherwise, we can not write: lib_usb not found code: -5
		usbDevice.done()

		fmt.Printf("USB close")
	}

	return true, nil
}

func handleControlRead(arguments interface{}) (reply interface{}, err error) {
	// arg := arguments.(map[string]interface{})
	arg := arguments.(map[interface{}]interface{})
	request := uint8(arg["cmd"].(int32))
	val := uint16(arg["value"].(int32))
	idx := uint16(arg["idx"].(int32))
	length := arg["len"].(int32)

	// fmt.Printf("request:%d,val:%d, idx:%d, length: %d\n", request, val, idx, length)

	// rType := gousb.ControlIn | gousb.ControlVendor | gousb.ControlInterface
	buf := make([]byte, length)
	numBytes, err := usbDevice.device.Control(
		gousb.ControlIn|gousb.ControlVendor|gousb.ControlInterface,
		request, val, idx, buf)

	// fmt.Printf("buf: %s, numBytes: %d\n", buf, numBytes)
	// just to avoid error message
	numBytes += 1

	return buf, err
}

func handleControlWrite(arguments interface{}) (reply interface{}, err error) {
	// arg := arguments.(map[string]interface{})
	arg := arguments.(map[interface{}]interface{})
	request := uint8(arg["cmd"].(int32))
	val := uint16(arg["value"].(int32))
	idx := uint16(arg["idx"].(int32))
	data := arg["data"].([]byte)

	// rType := gousb.ControlOut | gousb.ControlVendor | gousb.ControlInterface
	numBytes, err := usbDevice.device.Control(
		gousb.ControlOut|gousb.ControlVendor|gousb.ControlInterface,
		request, val, idx, data)
	return int32(numBytes), err
}

func handleRead(arguments interface{}) (reply interface{}, err error) {
	length := arguments.(int32)

	// fmt.Printf("handleRead: length = %d\n",length)

	buf := make([]byte, length)
	// fmt.Printf("handleRead: buf = %s\n",buf)
	// fmt.Printf("handleRead: length = %d\n",length)
	numBytes, err := usbDevice.epIn.Read(buf)

	// fmt.Printf("handleRead: numBytes = %d\n",numBytes)
	// fmt.Printf("handleRead: buf = %s\n",buf)

	// return buf[:numBytes], err
	// just to avoid error message
	numBytes += 1
	// fmt.Printf("handleRead: len = %d\n",numBytes)
	return buf[:length], err
}

func handleWrite(arguments interface{}) (reply interface{}, err error) {
	data := arguments.([]byte)
	// data := make([]byte, 5)
	// for i := range data {
	// 		data[i] = byte(i)
	// }

	numBytes, err := usbDevice.epOut.Write(data)

	return int32(numBytes), err
}

func findCwSca(vid, pid uint16) func(desc *gousb.DeviceDesc) bool {
	return func(desc *gousb.DeviceDesc) bool {
		return desc.Product == gousb.ID(pid) && desc.Vendor == gousb.ID(vid)
	}
}

// handleGetUsbInfo is called when the method getUsbInfo is invoked by
// the dart code.
//
// Supported return types of StandardMethodCodec codec are described in a table:
// https://godoc.org/github.com/go-flutter-desktop/go-flutter/plugin#StandardMessageCodec
func handleGetUsbInfo(arguments interface{}) (reply interface{}, err error) {
	var retVal string
	var deviceDesc []*gousb.DeviceDesc

	ctx := gousb.NewContext()
	defer ctx.Close()

	// Debugging can be turned on; this shows some of the inner workings of the libusb package.
	ctx.Debug(*debug)

	// OpenDevices is used to find the devices to open.
	devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		// The usbid package can be used to print out human readable information.
		retVal += fmt.Sprintf("%03d.%03d %s:%s %s\n", desc.Bus, desc.Address, desc.Vendor, desc.Product, usbid.Describe(desc))
		retVal += fmt.Sprintf("  Protocol: %s\n", usbid.Classify(desc))
		deviceDesc = append(deviceDesc, desc)

		// The configurations can be examined from the DeviceDesc, though they can only
		// be set once the device is opened.  All configuration references must be closed,
		// to free up the memory in libusb.
		for _, cfg := range desc.Configs {
			// This loop just uses more of the built-in and usbid pretty printing to list
			// the USB devices.
			retVal += fmt.Sprintf("  %s:\n", cfg)
			for _, intf := range cfg.Interfaces {
				retVal += fmt.Sprintf("    --------------\n")
				for _, ifSetting := range intf.AltSettings {
					retVal += fmt.Sprintf("    %s\n", ifSetting)
					retVal += fmt.Sprintf("      %s\n", usbid.Classify(ifSetting))
					for _, end := range ifSetting.Endpoints {
						retVal += fmt.Sprintf("      %s\n", end)
					}
				}
			}
			retVal += fmt.Sprintf("    --------------\n")
		}
		// fmt.Printf(retVal)

		// After inspecting the descriptor, return true or false depending on whether
		// the device is "interesting" or not.  Any descriptor for which true is returned
		// opens a Device which is retuned in a slice (and must be subsequently closed).
		return false
	})

	// All Devices returned from OpenDevices must be closed.
	defer func() {
		for _, d := range devs {
			d.Close()
		}
	}()

	// OpenDevices can occasionally fail, so be sure to check its return value.
	if err != nil {
		log.Fatalf("list: %s", err)
	}

	for _, dev := range devs {
		// Once the device has been selected from OpenDevices, it is opened
		// and can be interacted with.
		_ = dev
	}

	b, err := json.Marshal(deviceDesc)

	os.Stdout.Write(b)

	return b, err
}
