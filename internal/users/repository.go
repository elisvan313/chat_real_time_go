package users

import (
	"chat_real_time_go/internal/domain"
	"database/sql"

	"github.com/gin-gonic/gin"
)

const (
	getAllUsersQuery = "SELECT * FROM users"
)

type Repository interface {
	GetAll(c *gin.Context) ([]domain.User, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(c *gin.Context) ([]domain.User, error) {
	users := []domain.User{}

	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := domain.User{}
		_ = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.FirstName, &user.LastName, &user.RegistrationDate)
		users = append(users, user)
	}

	return users, nil
}
