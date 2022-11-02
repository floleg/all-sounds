package user

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
)

// FindById retrieves a User by id, eager loading UserTracks associations
func FindById(id int, user model.User) (model.User, error) {
	err := db.DBCon.Model(&model.User{}).Preload("Tracks").First(&user, id).Error

	return user, err
}
