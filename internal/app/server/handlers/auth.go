package handlers

import (
	"auth_audit/internal/app/server/DTO"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) register(ctx *gin.Context) {
	var body DTO.RegisterUserDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if body.Login == "" || body.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "login and password must be provided"})
		return
	}

	session, err := h.authService.Register(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
	ctx.JSON(http.StatusOK, gin.H{"token": session.Token})
}
