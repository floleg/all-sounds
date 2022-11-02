// Package model declares the gorm entities used by the API
package model

// Album data entity
type Album struct {
	Entity
	Title       string `faker:"sentence"`
	ReleaseYear uint64
	Tracks      []Track `gorm:"many2many:album_tracks;" faker:"-"`
}
