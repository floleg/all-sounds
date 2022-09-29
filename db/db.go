package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() {
	// conn := fmt.Sprintf("host=db port=5432 user=%s password=%s dbname=%s sslmode=disable",
	// os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	conn := fmt.Sprintf("host=db port=5432 user=%s password=%s dbname=%s sslmode=disable",
		"postgres", "-NQI2tIM?|G>B@A2", "all-sounds")

	_, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
