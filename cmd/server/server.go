// Package cmd/server is used to run the API server as a cli tool
package main

import (
	"allsounds/internal/router"
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
	err = db.InitPostgresDB(&appConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot initiate db connection")
	}

	// Run initial SQL migrations
	migration.CreateTables()

	// Instantiate gin router
	appRouter := router.NewRouter()
	err = appRouter.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot instantiate http server")
	}
}
