//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
)

func InitializeApp() (*App, error) {
	wire.Build(
		InfraSet,
		DBSet,
		UseCaseSet,
		HandlerSet,
		ServiceSet,
		RepositorySet,
		AuthSet,
		DeliverySet,
		wire.Struct(new(App), "*"),
	)
	return nil, nil
}
