package users

import (
	"chat_real_time_go/internal/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository interface {
	GetOne(c *gin.Context) domain.User
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(domain.User{})
	return &repository{
		db: db,
	}
}

func (r *repository) GetOne(c *gin.Context) domain.User {
	user := domain.User{}
	r.db.First(&user)

	return user
}
