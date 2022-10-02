package migration

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
)

func CreateTables() {
	db.DBCon.AutoMigrate(
		&model.Album{},
		&model.Artist{},
		&model.Track{},
		&model.User{},
	)
}
