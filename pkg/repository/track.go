package repository

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
)

// Track repository exposes gorm persistence methods
// in order to interact with the track postgresql table.
type Track struct {
	BaseRepo Repository
}

// FindById retrieves a Track by id, eager loading AlbumTracks associations
func (tr Track) FindById(id int, track model.Track) (model.Track, error) {
	err := db.DBCon.Model(&model.Track{}).Preload("Albums").First(&track, id).Error

	return track, err
}
