package handlers

import (
	wallet2 "1som/internal/application/use_case/wallet"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	getWallet wallet2.GetWalletUseCase
}

func NewWalletHandler(
	getWallet wallet2.GetWalletUseCase,
) *WalletHandler {
	return &WalletHandler{
		getWallet: getWallet,
	}
}

func (h *WalletHandler) GetWallet(c *gin.Context) {
	uid := c.GetString("uid")
	wallet, err := h.getWallet.Execute(context.Background(), uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, wallet)
}

func (h *WalletHandler) GetBalance(c *gin.Context) {
	uid := c.GetString("uid")

	c.JSON(http.StatusOK, gin.H{
		"uid":     uid,
		"balance": 1500,
	})
}
