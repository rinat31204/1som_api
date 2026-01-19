package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	// сюда потом добавишь usecase
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetUserInfo(c *gin.Context) {
	uid := c.GetString("uid")

	c.JSON(http.StatusOK, gin.H{
		"uid":  uid,
		"name": "Test User",
	})
}
