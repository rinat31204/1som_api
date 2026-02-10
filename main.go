package main

import (
	"1som/internal/wire"
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
