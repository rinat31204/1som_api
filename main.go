package main

import (
	"1som/internal/di/wire"
	"log"
)

func main() {
	r, err := wire.InitializeApp()
	if err != nil {
		log.Fatal(err)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
