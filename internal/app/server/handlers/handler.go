package handlers

import (
	"auth_audit/internal/app/server/middlewares"
	"auth_audit/internal/app/server/services/interfaces"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService   interfaces.AuthService
	validateToken interfaces.ValidateToken
}

func NewHandler(authService interfaces.AuthService, validateToken interfaces.ValidateToken) *Handler {
	return &Handler{
		authService:   authService,
		validateToken: validateToken,
	}
}

func (h Handler) Router() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/login", h.login)

		needToken := auth.Group("/audits")
		needToken.Use(middlewares.AuthMiddleware(h.validateToken))
		{
			needToken.GET("/", h.getAuthAuditByToken)
			needToken.DELETE("/", h.clearAuthAuditsByToken)
		}
	}

	return router
}
