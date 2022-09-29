package main

import (
	"allsounds/db"
	"fmt"
)

func main() {
	db.Init()
	db.GetDB()

	fmt.Println("Hello world")
}
