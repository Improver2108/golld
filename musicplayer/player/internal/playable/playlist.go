package playable

type playlist struct {
	name  string
	size  int
	songs []song
}

func NewPlaylist(name string) *playlist {
	return &playlist{name: name}
}

func (p *playlist) GetName() string {
	return p.name
}

func (p *playlist) GetSize() int {
	return p.size
}

func (p *playlist) GetSongs() []song {
	return p.songs
}

func (p *playlist) AddSong(song song) {
	p.songs = append(p.songs, song)
}
