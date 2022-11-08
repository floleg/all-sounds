package album

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
)

// FindById retrieves an Album by id, eager loading AlbumTracks associations
func FindById(id int, album model.Album) (model.Album, error) {
	err := db.DBCon.Model(&model.Album{}).Preload("Tracks").First(&album, id).Error

	return album, err
}
