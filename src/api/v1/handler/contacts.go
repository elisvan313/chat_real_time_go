package handler

import (
	"chat_real_time_go/internal/contacts"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Contact struct {
	contactService contacts.Service
}

func NewContact(c contacts.Service) *Contact {
	return &Contact{
		contactService: c,
	}
}

func (c *Contact) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		contact, err := c.contactService.GetAll(ctx)
		if err != nil {

			ctx.JSON(http.StatusNotFound, nil)
			return
		} else {
			ctx.JSON(http.StatusOK, contact)
		}
	}
}
