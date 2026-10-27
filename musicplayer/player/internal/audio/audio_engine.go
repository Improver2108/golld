package audio

import "fmt"

type audio interface {
	GetName() string
	GetArtist() string
	GetPath() string
}

type audioOutputDevice interface {
	PlayAudioFromOutput(audio audio)
}

type audioEngine struct {
	currentAudio  audio
	isSongPlaying bool
}

func newAudioEngine() *audioEngine {
	return &audioEngine{}
}

func (a *audioEngine) Play(song audio, audioOutput audioOutputDevice) {
	if a.isSongPlaying {
		fmt.Println("Song already playing")
	}
	audioOutput.PlayAudioFromOutput(song)
	a.isSongPlaying = true
}

func (a *audioEngine) Pause() {
	a.isSongPlaying = false
}
