package artist

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
)

// FindById retrieves an Artist by id, eager loading ArtistTracks associations
func FindById(id int, artist model.Artist) (model.Artist, error) {
	err := db.DBCon.Model(&model.Artist{}).Preload("Tracks").First(&artist, id).Error

	return artist, err
}
