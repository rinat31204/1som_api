package wire

import (
	"1som/internal/auth/application"
	authFirebase "1som/internal/auth/infrastructure/firebase"
	"1som/internal/delivery/http"
	"1som/internal/delivery/http/handlers"
	infraFirebase "1som/internal/infrastructure/firebase"

	"github.com/google/wire"
)

var InfraSet = wire.NewSet(
	infraFirebase.NewClient,
)

var AuthSet = wire.NewSet(
	authFirebase.NewFirebaseProvider,
	application.NewAuthService,
	//wire.Bind(new(authDomain.Provider), new(*authFirebase.Provider)),
)

var HandlerSet = wire.NewSet(
	handlers.NewUserHandler,
	handlers.NewWalletHandler,
	handlers.NewWinnersHandler,
)

var DeliverySet = wire.NewSet(
	http.NewRouter,
	provideRouterDeps,
)

func provideRouterDeps(
	authService *application.AuthService,
	userHandler *handlers.UserHandler,
	walletHandler *handlers.WalletHandler,
	winnersHandler *handlers.WinnersHandler,
) http.RouterDeps {
	return http.RouterDeps{
		AuthService:    authService,
		UserHandler:    userHandler,
		WalletHandler:  walletHandler,
		WinnersHandler: winnersHandler,
	}
}
