package repository

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
)

// Artist repository exposes gorm persistence methods
// in order to interact with the artist postgresql table.
type Artist struct {
	BaseRepo Repository
}

// FindById retrieves an Artist by id, eager loading ArtistTracks associations
func (ar Artist) FindById(id int, artist model.Artist) (model.Artist, error) {
	err := db.DBCon.Model(&model.Artist{}).Preload("Tracks").First(&artist, id).Error

	return artist, err
}
