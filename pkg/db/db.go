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
	dbCon *gorm.DB
)

func Init(config *config.Config) {
	// Build gorm connection string
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUSer, config.DBPassword, config.DBName)

	// Instantiate SQL logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// Open gorm connection
	DBCon, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		Logger: newLogger,
	})

	setDB(DBCon)

	if err != nil {
		panic(err)
	}
}

func setDB(db *gorm.DB) {
	dbCon = db
}

func GetDB() *gorm.DB {
	return dbCon
}
