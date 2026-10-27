package playable

type Playlist interface {
	GetName() string
	GetSize() int
	GetSongs() []song
	AddSong(song song)
}

type song interface {
	GetName() string
	GetArtist() string
	GetPath() string
}
