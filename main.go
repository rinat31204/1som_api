package main

import (
	"1som/internal/di/wire"
	"log"
)

func main() {
	app, err := wire.InitializeApp()
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
