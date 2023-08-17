package handlers

import (
	"auth_audit/internal/app/server/DTO"

	"github.com/gin-gonic/gin"

	"net/http"
)

// @Summary		User registration
// @Description	User registration using login and password
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			request			body		DTO.RegisterUserDTO	true	"Request with user data"
// @Success		201				{object}	string				"Successful response with token"
// @Failure		400				{object}	string				"Bad request"
// @Failure		500				{object}	string				"Internal Server Error"
// @Router			/auth/register 	 [post]
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

// @Summary		User authorization
// @Description	User authorization using login and password
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			request			body		DTO.LoginUserDTO	true	"Request with user data"
// @Success		201				{object}	string				"Successful response with token"
// @Failure		400				{object}	string				"Bad request"
// @Failure		500				{object}	string				"Internal Server Error"
// @Router			/auth/login 	 [post]
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

// @Summary		Get authentication audit by token
// @Description	Get authentication audit events for the user associated with the provided token
// @Tags			auth
// @Security		ApiKeyAuth
// @Produce		json
// @Param			X-Token	header		string	true	"Authentication Token"
// @Success		200		{object}	DTO.AuthAuditDTO
// @Failure		401		{string}	string	"Unauthorized"
// @Failure		500		{string}	string	"Internal Server Error"
// @Router			/auth/audit [get]
func (h Handler) getAuthAuditByToken(ctx *gin.Context) {
	token := ctx.GetHeader("X-Token")

	audit, err := h.authService.GetAuthAuditByToken(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"audit": audit})
}

// @Summary		Clear authentication audit by token
// @Description	Clear the authentication audit events for the user associated with the provided token
// @Tags			auth
// @Security		ApiKeyAuth
// @Produce		json
// @Param			X-Token	header		string	true	"Authentication Token"
// @Success		200		{object}	string
// @Failure		401		{string}	string	"Unauthorized"
// @Failure		500		{string}	string	"Internal Server Error"
// @Router			/auth/audit/clear [delete]
func (h Handler) clearAuthAuditsByToken(ctx *gin.Context) {
	token := ctx.GetHeader("X-Token")

	if err := h.authService.ClearAuthAuditsByToken(token); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "authentication audit cleared"})
}
