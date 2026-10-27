package device

type bluetoothAdapter struct {
	bluetooth *bluetoothAudioApi
}

func newBluetoothAdapter() *bluetoothAdapter {
	return &bluetoothAdapter{
		bluetooth: newBluetoothAudioApi(),
	}
}

func (a *bluetoothAdapter) PlayAudioFromOutput(audio audio) {
	a.bluetooth.playAudioViaBluetooth(audio.GetData())
}

type wiredAdapter struct {
	wired *wiredAudioApi
}

func newWiredAdapter() *wiredAdapter {
	return &wiredAdapter{
		wired: newWiredAudioApi(),
	}
}

func (a *wiredAdapter) PlayAudioFromOutput(audio audio) {
	a.wired.playAudioViaWired(audio.GetData())
}
