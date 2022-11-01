// Package cmd/server is used to run the API server as a cli tool
package main

import (
	"allsounds/internal/router"
	"allsounds/pkg/config"
	"allsounds/pkg/db"
	"allsounds/pkg/migration"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	// Retrieve configuration based on ENV system variable
	appConfig, err := config.LoadConfig(os.Getenv("ENV"), "./configs")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config file")
	}

	// Open postgres connection
	err = db.Init(&appConfig)
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
