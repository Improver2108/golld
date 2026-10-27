package player

import (
	"errors"
	"fmt"

	"github.com/improver2108/golld/musicplayer/player/internal/device"
	"github.com/improver2108/golld/musicplayer/player/internal/playable"
)

type audioEngine interface {
}

type song interface {
	GetName() string
	GetArtist() string
	Play()
}

type audioDevice interface {
	PlayAudio(s song)
}

type playerFacade struct {
	audioEngine    audioEngine
	loadedPlaylist playable.Playlist
	playStrategy   playable.PlayStrategy
}

func newPlayerFacade(engine audioEngine) *playerFacade {
	return &playerFacade{
		audioEngine: engine,
	}
}

func (p *playerFacade) connectDevice(deviceType device.DeviceType) {
	device.GetDeviceManagerInstance().Connect(deviceType)
}

func (p *playerFacade) setPlayStrategy(strategyType playable.PlayStrategyType) {
	if strategy, err := playable.GetPlayStrategyInstance().GetStrategy(strategyType); err == nil {
		p.playStrategy = strategy
		return
	}
	fmt.Println("Provide a valid strategy type")
}

func (p *playerFacade) LoadPlaylist(idx int) error {
	p.loadedPlaylist = playable.GetPlaylistManagerInstance().GetPlaylist(idx)
	if p.playStrategy == nil {
		return errors.New("No audio device connected")
	}
	p.playStrategy.SelectPlaylist(p.loadedPlaylist)
	return nil
}
