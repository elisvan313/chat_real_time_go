package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Id               int    `gorm:"primaryKey"`
	Username         string `json:"username" gorm:"unique_index;not null"`
	Password         string `json:"password" gorm:"not null"`
	Email            string `json:"email" gorm:"unique_index;not null"`
	FirstName        string `json:"first_name" gorm:"not null"`
	LastName         string `json:"last_name" gorm:"not null"`
	RegistrationDate string `json:"registration_date" gorm:"not null"`
}
