package wire

import (
	"1som/internal/application/service"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router      *gin.Engine
	UserService *service.UserService
}
