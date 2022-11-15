package album

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
)

type Repository interface {
	FindById(int, *model.Album) error
}

type Album struct{}

// FindById retrieves an Album by id, eager loading AlbumTracks associations
func (a Album) FindById(id int, album *model.Album) error {
	return db.DBCon.Model(&model.Album{}).Preload("Tracks").First(album, id).Error
}
