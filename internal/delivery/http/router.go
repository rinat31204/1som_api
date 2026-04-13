package http

import (
	"1som/internal/auth/application"
	"1som/internal/auth/delivery/http/middleware"
	"1som/internal/delivery/http/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouterDeps struct {
	AuthService *application.AuthService

	UserHandler    *handlers.UserHandler
	WalletHandler  *handlers.WalletHandler
	WinnersHandler *handlers.WinnersHandler
}

func NewRouter(deps RouterDeps) *gin.Engine {
	r := gin.New()

	r.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := r.Group("/api")

	authMiddleware := middleware.NewAuthMiddleware(*deps.AuthService)

	protected := api.Group("/")
	protected.Use(authMiddleware)
	{
		protected.GET("/user-info", deps.UserHandler.GetUserInfo)

		protected.GET("/wallet", deps.WalletHandler.GetWallet)
		protected.GET("/wallet/balance", deps.WalletHandler.GetBalance)

		protected.GET("/winners", deps.WinnersHandler.GetWinners)
		protected.GET("/winner-detail/:id", deps.WinnersHandler.GetWinnerDetail)
	}
	return r
}
