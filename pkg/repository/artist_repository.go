package repository

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
)

type ArtistRepository struct {
	BaseRepo Repository
}

// Retrieve Artist by id, eager loading AlbumTracks associations
func (ar ArtistRepository) FindById(id int, artist model.Artist) model.Artist {
	db.DBCon.First(&artist, id)

	return artist

}
