// Package db opens a gorm Postgres database connection
package db

import (
	"allsounds/pkg/config"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DBCon *gorm.DB
)

// InitPostgresDB method does three things:
//
//   - Build a gorm postgres connection string based on the runtime configuration
//   - Instantiate a gorm SQL logger to monitor the executed queries
//   - Open a postgres connection
func InitPostgresDB(config *config.Config) error {
	// Build gorm connection string
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUSer, config.DBPassword, config.DBName)

	// Instantiate SQL logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// Open gorm connection
	var err error
	DBCon, err = gorm.Open(postgres.Open(conn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return err
	}

	return nil
}
