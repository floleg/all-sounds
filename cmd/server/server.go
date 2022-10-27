package main

import (
	"allsounds/internal/router"
	"allsounds/pkg/config"
	"allsounds/pkg/db"
	"allsounds/pkg/migration"
	"log"
	"os"
)

func main() {
	// Retrieve configuration based on ENV system variable
	appConfig, err := config.LoadConfig(os.Getenv("ENV"), "./configs")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Open postgres connection
	err = db.Init(&appConfig)
	if err != nil {
		log.Fatal("cannot initiate db connection:", err)
	}

	// Run initial SQL migrations
	migration.CreateTables()

	// Instantiate gin router
	appRouter := router.NewRouter()
	err = appRouter.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot initiate db connection:", err)
	}
}
