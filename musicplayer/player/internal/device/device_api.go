package device

import "fmt"

type bluetoothAudioApi struct{}

func newBluetoothAudioApi() *bluetoothAudioApi {
	return &bluetoothAudioApi{}
}

func (api *bluetoothAudioApi) playAudioViaBluetooth(data string) {
	fmt.Println(data, "is played as audio over bluetooth")
}

type wiredAudioApi struct{}

func newWiredAudioApi() *wiredAudioApi {
	return &wiredAudioApi{}
}

func (api *wiredAudioApi) playAudioViaWired(data string) {
	fmt.Println(data, "is played as audio over wired")
}
