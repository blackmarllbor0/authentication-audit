package handlers

import (
	"auth_audit/internal/app/server/services/interfaces"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService interfaces.AuthService
}

func NewHandler(authService interfaces.AuthService) *Handler {
	return &Handler{authService: authService}
}

func (h Handler) Router() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/login", h.login)
		auth.GET("/audit", h.getAuthAuditByToken)
	}

	return router
}
