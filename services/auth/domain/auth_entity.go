package domain

type User struct {
	Id       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
