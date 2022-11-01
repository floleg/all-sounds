package repository

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
)

type User struct {
	BaseRepo Repository
}

// FindById retrieves a User by id, eager loading UserTracks associations
func (tr User) FindById(id int, user model.User) (model.User, error) {
	err := db.DBCon.Model(&model.User{}).Preload("Tracks").First(&user, id).Error

	return user, err
}
