package model

type Album struct {
	Entity
	Title       string `faker:"sentence"`
	ReleaseYear uint8
	AlbumTracks []AlbumTrack `faker:"-"`
}
