package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
}

func NewWalletHandler() *WalletHandler {
	return &WalletHandler{}
}

func (h *WalletHandler) GetWallet(c *gin.Context) {
	uid := c.GetString("uid")

	c.JSON(http.StatusOK, gin.H{
		"uid":    uid,
		"wallet": "main",
	})
}

func (h *WalletHandler) GetBalance(c *gin.Context) {
	uid := c.GetString("uid")

	c.JSON(http.StatusOK, gin.H{
		"uid":     uid,
		"balance": 1500,
	})
}
