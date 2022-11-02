package track

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
)

// FindById retrieves a Track by id, eager loading AlbumTracks associations
func FindById(id int, track model.Track) (model.Track, error) {
	err := db.DBCon.Model(&model.Track{}).Preload("Albums").First(&track, id).Error

	return track, err
}
