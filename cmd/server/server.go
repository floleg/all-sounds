package main

import (
	"allsounds/pkg/config"
	"allsounds/pkg/db"
	"allsounds/pkg/migration"
	"log"
	"os"
)

func main() {
	// Retrieve configuration based on ENV system variable
	config, err := config.LoadConfig(os.Getenv("ENV"), "./configs")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Open postgres connection
	db.Init(&config)

	// Run initial SQL migrations
	migration.CreateTables()
}
