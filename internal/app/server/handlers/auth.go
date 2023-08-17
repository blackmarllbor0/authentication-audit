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

	session, err := h.authService.Register(body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"token": session.Token})
}

func (h Handler) login(ctx *gin.Context) {
	var body DTO.LoginUserDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	session, err := h.authService.Login(body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": session.Token})
}

func (h Handler) getAuthAuditByToken(ctx *gin.Context) {
	token := ctx.GetHeader("X-Token")

	audit, err := h.authService.GetAuthAuditByToken(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"audit": audit})
}

func (h Handler) clearAuthAuditsByToken(ctx *gin.Context) {
	token := ctx.GetHeader("X-Token")

	if err := h.authService.ClearAuthAuditsByToken(token); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "authentication audit cleared"})
}
