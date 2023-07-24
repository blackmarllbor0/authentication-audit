package handlers

import (
	"auth_audit/internal/app/server/handlers/users"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userHandler users.Handler
}

func NewHandler(userHandler users.Handler) *Handler {
	return &Handler{userHandler: userHandler}
}

func (h Handler) Router() *gin.Engine {
	router := gin.New()

	usersGroup := router.Group("users")
	{
		usersGroup.POST("/", h.userHandler.CreateUser)
	}

	return router
}
