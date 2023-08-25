package contacts

import (
	"chat_real_time_go/internal/domain"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(c *gin.Context) ([]domain.Contact, error)
	GetById(c *gin.Context) ([]domain.Contact, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(c *gin.Context) ([]domain.Contact, error) {
	contacts := []domain.Contact{}

	rows, err := r.db.Query("select * from contacts")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		contact := domain.Contact{}
		_ = rows.Scan(&contact.Id, &contact.IdUser, &contact.UserName, &contact.Email, &contact.FirstName, &contact.LastName, &contact.ContactDate)
		contacts = append(contacts, contact)
	}

	return contacts, nil

}

// GetById implements Repository.
func (r *repository) GetById(c *gin.Context) ([]domain.Contact, error) {
	panic("unimplemented")
}
