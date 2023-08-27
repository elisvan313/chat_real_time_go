package handler

import (
	"chat_real_time_go/internal/users"

	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	userService users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		userService: u,
	}
}

func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := u.userService.GetOne(ctx)
		ctx.JSON(http.StatusOK, user)
	}
}

func (u *User) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := u.userService.CreateUser(ctx)
		ctx.JSON(http.StatusOK, user)
	}
}
