package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WinnersHandler struct {
}

func NewWinnersHandler() *WinnersHandler {
	return &WinnersHandler{}
}

func (h *WinnersHandler) GetWinners(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"winners": []string{"user1", "user2"},
	})
}

func (h *WinnersHandler) GetWinnerDetail(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"prize":  1000,
		"status": "paid",
	})
}
