package migration

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
)

func CreateTables() {
	// Create user table if not existing in current db
	if (!db.DBCon.Migrator().HasTable(&model.User{})) {
		db.DBCon.Migrator().CreateTable(&model.User{})
	}

	// Create artist table if not existing in current db
	if (!db.DBCon.Migrator().HasTable(&model.Artist{})) {
		db.DBCon.Migrator().CreateTable(&model.Artist{})
	}

	// Create album table if not existing in current db
	if (!db.DBCon.Migrator().HasTable(&model.Album{})) {
		db.DBCon.Migrator().CreateTable(&model.Album{})
	}

	// Create track table if not existing in current db
	if (!db.DBCon.Migrator().HasTable(&model.Track{})) {
		db.DBCon.Migrator().CreateTable(&model.Track{})
	}

	// Create user_track_favorite table if not existing in current db
	if (!db.DBCon.Migrator().HasTable(&model.UserTrackFavorite{})) {
		db.DBCon.Migrator().CreateTable(&model.UserTrackFavorite{})
	}
}
