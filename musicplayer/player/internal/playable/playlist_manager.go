package playable

import "sync"

type playlistManager struct {
	playlists []*playlist
}

var (
	instance *playlistManager
	once     sync.Once
)

func GetPlaylistManagerInstance() *playlistManager {
	once.Do(func() {
		instance = &playlistManager{}
	})
	return instance
}

func (p *playlistManager) AddPlaylist(playlist *playlist) {
	p.playlists = append(p.playlists, playlist)
}

func (p *playlistManager) GetPlaylist(idx int) *playlist {
	return p.playlists[idx]
}

func (p *playlistManager) DeletePlaylist(idx int) {
	if idx < 0 || idx >= len(p.playlists) {
		return
	}
	p.playlists = append(p.playlists[:idx], p.playlists[idx+1:]...)
}
