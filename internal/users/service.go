package users

import (
	"chat_real_time_go/internal/domain"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetOne(c *gin.Context) domain.User
}

type service struct{ repository Repository }

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetOne(c *gin.Context) domain.User {
	return s.repository.GetOne(c)
}
