package main

import (
	"allsounds/pkg/config"
	"allsounds/pkg/db"
	"fmt"
	"log"
	"os"
)

func main() {
	config, err := config.LoadConfig(os.Getenv("ENV"), "./configs")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db.Init(&config)
	db.GetDB()

	fmt.Println("Hello world")
}
