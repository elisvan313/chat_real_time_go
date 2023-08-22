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
		user, err := u.userService.GetAll(ctx)
		if err != nil {

			ctx.JSON(http.StatusNotFound, nil)
			return
		} else {
			ctx.JSON(http.StatusOK, user)
		}
	}
}
