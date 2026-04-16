package wire

import (
	"1som/internal/application/service"
	"1som/internal/auth/application"
	authFirebase "1som/internal/auth/infrastructure/firebase"
	"1som/internal/delivery/http"
	"1som/internal/delivery/http/handlers"
	"1som/internal/infrastructure/eventbus/inmemory"
	infraFirebase "1som/internal/infrastructure/firebase"
	"1som/internal/infrastructure/postgres"
	"1som/internal/infrastructure/postgres/repository"

	"github.com/google/wire"
)

var InfraSet = wire.NewSet(
	infraFirebase.NewClient,
	inmemory.NewEventBus,
	wire.Bind(new(service.EventBus), new(*inmemory.EventBus)),
)

var RepositorySet = wire.NewSet(
	repository.NewUserRepository,
	repository.NewWalletRepository,
	repository.NewAccountRepository,
	repository.NewTransactionRepository,
)

var UnitOfWorkSet = wire.NewSet(
	postgres.NewUnitOfWorkFactory,
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

var ServiceSet = wire.NewSet(
	service.NewUserService,
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
