package repository

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
)

type AlbumRepository struct {
	BaseRepo Repository
}

// Retrieve Album by id, eager loading AlbumTracks associations
func (ar AlbumRepository) FindById(id int, album model.Album) (model.Album, error) {
	err := db.DBCon.Model(&model.Album{}).Preload("Tracks").First(&album, id).Error

	return album, err
}
