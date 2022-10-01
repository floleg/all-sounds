package migration

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
)

func CreateTables() {
	// Create user table if not existing in current db
	if (!db.GetDB().Migrator().HasTable(&model.User{})) {
		db.GetDB().Migrator().CreateTable(&model.User{})
	}

	// Create artist table if not existing in current db
	if (!db.GetDB().Migrator().HasTable(&model.Artist{})) {
		db.GetDB().Migrator().CreateTable(&model.Artist{})
	}

	// Create album table if not existing in current db
	if (!db.GetDB().Migrator().HasTable(&model.Album{})) {
		db.GetDB().Migrator().CreateTable(&model.Album{})
	}

	// Create track table if not existing in current db
	if (!db.GetDB().Migrator().HasTable(&model.Track{})) {
		db.GetDB().Migrator().CreateTable(&model.Track{})
	}

	// Create user_track_favorite table if not existing in current db
	if (!db.GetDB().Migrator().HasTable(&model.UserTrackFavorite{})) {
		db.GetDB().Migrator().CreateTable(&model.UserTrackFavorite{})
	}
}
