package device

import (
	"fmt"
	"sync"
)

type audio interface {
	GetData() string
}

type audioOutputDevice interface {
	PlayAudioFromOutput(audio audio)
}

type deviceManager struct {
	currentAudioOutputDevice audioOutputDevice
}

var (
	instance *deviceManager
	once     sync.Once
)

func GetDeviceManagerInstance() *deviceManager {
	once.Do(func() {
		instance = &deviceManager{}
	})
	return instance
}

func (d *deviceManager) createDevice(deviceType DeviceType) audioOutputDevice {
	switch deviceType {
	case Bluetooth:
		return newBluetoothAdapter()
	case Wired:
		return newWiredAdapter()
	}
	return nil
}

func (d *deviceManager) Connect(deviceType DeviceType) {
	d.currentAudioOutputDevice = d.createDevice(deviceType)
	switch deviceType {
	case Bluetooth:
		fmt.Println("Bluetooth device connected")
	case Wired:
		fmt.Println("Wired device connected")
	default:
		fmt.Println("Wrong device given")
	}
}

func (d *deviceManager) GetAudioDevice() audioOutputDevice {
	return d.currentAudioOutputDevice
}
