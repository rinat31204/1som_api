package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func NewApp() *firebase.App {
	ctx := context.Background()

	opt := option.WithCredentialsFile(
		"secrets/298d0-fbsvc-e11b1c9f2b.json", // путь из контейнера
	)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("firebase init error: %v", err)
	}

	return app
}
