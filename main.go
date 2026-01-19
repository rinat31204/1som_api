package main

import (
	"1som/internal/application/firebase"
	"1som/internal/delivery/http"
	"1som/internal/delivery/http/handlers"
)

func main() {
	// Firebase
	app := firebase.NewApp()
	authClient := firebase.NewAuthClient(app)
	authProvider := firebase.NewAuthProvider(authClient)

	// Handlers
	userHandler := handlers.NewUserHandler()
	walletHandler := handlers.NewWalletHandler()
	winnersHandler := handlers.NewWinnersHandler()

	// Router
	r := http.NewRouter(http.RouterDeps{
		AuthProvider:   authProvider,
		UserHandler:    userHandler,
		WalletHandler:  walletHandler,
		WinnersHandler: winnersHandler,
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
