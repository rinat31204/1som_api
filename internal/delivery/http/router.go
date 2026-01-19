package http

import (
	"1som/internal/application/auth"
	"1som/internal/delivery/http/handlers"
	"1som/internal/delivery/http/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouterDeps struct {
	AuthProvider auth.Provider

	UserHandler    *handlers.UserHandler
	WalletHandler  *handlers.WalletHandler
	WinnersHandler *handlers.WinnersHandler
}

func NewRouter(deps RouterDeps) *gin.Engine {
	r := gin.New()

	// глобальные middleware
	r.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// healthcheck
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// =========================
	// API
	// =========================
	api := r.Group("/api")

	// ---------- PUBLIC ----------
	api.GET("/winners", deps.WinnersHandler.GetWinners)
	api.GET("/winner-detail/:id", deps.WinnersHandler.GetWinnerDetail)

	// ---------- PROTECTED ----------
	authMiddleware := middleware.NewAuthMiddleware(deps.AuthProvider)

	protected := api.Group("/")
	protected.Use(authMiddleware)

	{
		// USER
		protected.GET("/user-info", deps.UserHandler.GetUserInfo)

		// WALLET
		protected.GET("/wallet", deps.WalletHandler.GetWallet)
		protected.GET("/wallet/balance", deps.WalletHandler.GetBalance)
	}

	return r
}
