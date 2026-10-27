package playable

import (
	"errors"
)

type sequentialStrategy struct {
	currentPlaylist Playlist
	currentIndex    int
}

func newSequentialStrategy() *sequentialStrategy {
	return &sequentialStrategy{
		currentIndex: -1,
	}
}

func (s *sequentialStrategy) SelectPlaylist(playlist Playlist) {
	s.currentPlaylist = playlist
	s.currentIndex = -1
}

func (s *sequentialStrategy) HasNext() bool {
	return s.currentIndex+1 < s.currentPlaylist.GetSize()
}

func (s *sequentialStrategy) Next() (song, error) {
	if s.currentPlaylist == nil || s.currentPlaylist.GetSize() == 0 {
		return nil, errors.New("No playlist loaded or playlist is empty.")
	}
	s.currentIndex = s.currentIndex + 1
	return s.currentPlaylist.GetSongs()[s.currentIndex], nil
}

func (s *sequentialStrategy) HasPrevious() bool {
	return s.currentIndex-1 >= 0
}

func (s *sequentialStrategy) Previous() (song, error) {
	if s.currentPlaylist == nil || s.currentPlaylist.GetSize() == 0 {
		return nil, errors.New("No playlist loaded or playlist is empty.")
	}
	s.currentIndex = s.currentIndex - 1
	return s.currentPlaylist.GetSongs()[s.currentIndex], nil
}
