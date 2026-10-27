package audio

type song struct {
	name   string
	artist string
	path   string
}

func NewSong(name, artist, path string) *song {
	return &song{
		name:   name,
		artist: artist,
		path:   path,
	}
}

func (s *song) GetName() string {
	return s.name
}

func (s *song) GetArtist() string {
	return s.name
}

func (s *song) GetPath() string {
	return s.name
}
