package repository

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
)

type TrackRepository struct {
	BaseRepo Repository
}

// Retrieve Artist by id, eager loading AlbumTracks associations
func (tr TrackRepository) FindById(id int, track model.Track) (model.Track, error) {
	err := db.DBCon.Model(&model.Track{}).Preload("Albums").First(&track, id).Error

	return track, err
}
