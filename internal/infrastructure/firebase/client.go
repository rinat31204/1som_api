package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	fbAuth "firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// Client обертка для Firebase
type Client struct {
	App  *firebase.App
	Auth *fbAuth.Client
}

// NewClient создает новый Firebase клиент
func NewClient() (*fbAuth.Client, error) {
	opt := option.WithCredentialsFile("secrets/298d0-fbsvc-e11b1c9f2b.json")
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	auth, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	//return &Client{
	//	App:  app,
	//	Auth: auth,
	//}, nil

	return auth, nil
}
