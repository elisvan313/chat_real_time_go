package users

import (
	"chat_real_time_go/internal/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	GetOne(c *gin.Context) domain.User
	CreateUser(c *gin.Context) domain.User
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

func (r *repository) CreateUser(c *gin.Context) domain.User {
	user := domain.User{}
	c.Bind(&user)
	HashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 9)
	user.Password = string(string(HashedPassword))
	r.db.Create(&user)
	return user
}
