package playable

import (
	"errors"
	"sync"
)

type PlayStrategyType string

const (
	Sequential PlayStrategyType = "sequential"
)

type PlayStrategy interface {
	SelectPlaylist(playlist Playlist)
	HasNext() bool
	Next() (song, error)
	HasPrevious() bool
	Previous() (song, error)
}

type playStrategyManager struct {
	sequential *sequentialStrategy
}

var (
	playStrategyInstance *playStrategyManager
	runPlayStrategyOnce  sync.Once
)

func newPlayStrategyManager() *playStrategyManager {
	return &playStrategyManager{
		sequential: newSequentialStrategy(),
	}
}

func GetPlayStrategyInstance() *playStrategyManager {
	runPlayStrategyOnce.Do(func() {
		playStrategyInstance = newPlayStrategyManager()
	})
	return playStrategyInstance
}

func (m *playStrategyManager) GetStrategy(strategyType PlayStrategyType) (PlayStrategy, error) {
	switch strategyType {
	case Sequential:
		return m.sequential, nil
	default:
		return nil, errors.New("Not a valid strategy name")
	}
}
