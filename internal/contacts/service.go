package contacts

import (
	"chat_real_time_go/internal/domain"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetAll(c *gin.Context) ([]domain.Contact, error)
}

type service struct{ repository Repository }

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAll(c *gin.Context) ([]domain.Contact, error) {
	return s.repository.GetAll(c)
}
