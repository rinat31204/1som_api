//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeApp() (*gin.Engine, error) {
	wire.Build(
		InfraSet,
		HandlerSet,
		AuthSet,
		DeliverySet,
	)
	return nil, nil
}
