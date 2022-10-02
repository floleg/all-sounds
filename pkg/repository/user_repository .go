package repository

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
)

type UserRepository struct {
	BaseRepo Repository
}

// Retrieve User by id, eager loading AlbumTracks associations
func (tr UserRepository) FindById(id int, user model.User) (model.User, error) {
	err := db.DBCon.Model(&model.User{}).Preload("Tracks").First(&user, id).Error

	return user, err
}
