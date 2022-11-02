package main

import (
	"allsounds/pkg/config"
	"allsounds/pkg/db"
	"allsounds/pkg/migration"
	"github.com/rs/zerolog/log"
)

func main() {
	// Retrieve configuration based on ENV system variable
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	// Open postgres connection
	err = db.Init(&appConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot initiate db connection")
	}

	artists := migration.BulkInsertArtists(2)

	migration.BulkInsertAlbums(artists, 10)

	migration.BulkInsertUsers(10)
}
