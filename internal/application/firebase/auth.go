package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

func NewAuthClient(app *firebase.App) *auth.Client {
	ctx := context.Background()

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("firebase auth init error: %v", err)
	}

	return client
}
