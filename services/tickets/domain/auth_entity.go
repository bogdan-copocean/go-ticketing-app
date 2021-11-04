package domain

import (
	"fmt"
)

type User struct {
	Id       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func (u *User) ValidateFields() error {
	if u.Email == "" {
		return fmt.Errorf("email not provided")
	}

	if u.Password == "" {
		return fmt.Errorf("password not provided")
	}

	return nil
}
