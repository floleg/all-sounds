package migration

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
)

func CreateTables() {
	db.GetDB().AutoMigrate(
		&model.Album{},
		&model.Artist{},
		&model.Track{},
		&model.User{},
		&model.UserTrackFavorite{},
	)
}
