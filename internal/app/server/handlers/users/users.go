package users

import (
	"auth_audit/internal/app/server/services/users"
	"auth_audit/internal/pkg/repository/postgres/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	Users struct {
		userService users.Service
	}
	Handler interface {
		CreateUser(ctx *gin.Context)
	}
)

func NewUsers(userService users.Service) *Users {
	return &Users{userService: userService}
}

func (u Users) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := u.userService.CreateUser(user.Login, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser.Password = ""
	ctx.JSON(http.StatusOK, createdUser)
}
