package wire

import (
	"1som/internal/application/auth"
	"1som/internal/application/firebase"
	"1som/internal/delivery/http"
	"1som/internal/delivery/http/handlers"

	"github.com/google/wire"
)

var FirebaseSet = wire.NewSet(
	firebase.NewApp,
	firebase.NewAuthClient,
	firebase.NewAuthProvider,
)

var HandlerSet = wire.NewSet(
	handlers.NewUserHandler,
	handlers.NewWalletHandler,
	handlers.NewWinnersHandler,
)

var RouterSet = wire.NewSet(
	NewRouterDeps,
	http.NewRouter,
)

func NewRouterDeps(
	authProvider auth.Provider,
	userHandler *handlers.UserHandler,
	walletHandler *handlers.WalletHandler,
	winnersHandler *handlers.WinnersHandler,
) http.RouterDeps {
	return http.RouterDeps{
		AuthProvider:   authProvider,
		UserHandler:    userHandler,
		WalletHandler:  walletHandler,
		WinnersHandler: winnersHandler,
	}
}
