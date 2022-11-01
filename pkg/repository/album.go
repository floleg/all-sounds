// Package repository provides the gorm db querying methods.
// Each repository structure is associated to a single gorm entity.
package repository

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
)

type Album struct {
	BaseRepo Repository
}

// FindById retrieves an Album by id, eager loading AlbumTracks associations
func (ar Album) FindById(id int, album model.Album) (model.Album, error) {
	err := db.DBCon.Model(&model.Album{}).Preload("Tracks").First(&album, id).Error

	return album, err
}
