// Package migration contains the gorm SQL structure migrations
// and test data generation.
package migration

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
	"github.com/rs/zerolog/log"
)

func CreateTables() {
	err := db.DBCon.AutoMigrate(
		&model.Album{},
		&model.Artist{},
		&model.Track{},
		&model.User{},
	)

	if err != nil {
		log.Fatal().Err(err)
	}
}
