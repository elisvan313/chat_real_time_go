package domain

type Contact struct {
	Id          int
	IdUser      int    `json:"id_user"`
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	ContactDate string `json:"contact_date"`
}
