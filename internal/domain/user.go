package domain

type User struct {
	Id               int
	Username         string `json:"username"`
	Password         string `json:"password"`
	Email            string `json:"email"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	RegistrationDate string `json:"registration_date"`
}
