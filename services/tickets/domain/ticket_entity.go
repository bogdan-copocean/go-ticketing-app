package domain

import "github.com/bogdan-user/go-ticketing-app/pkg/errors"

type Ticket struct {
	Id     string  `json:"id,omitempty"`
	Title  string  `json:"title"`
	Price  float64 `json:"price,omitempty"`
	UserId string  `json:"user_id,omitempty"`
}

func (ticket *Ticket) ValidateTicket() *errors.CustomErr {
	if ticket.Title == "" {
		return errors.NewBadRequestErr("Title not provided")
	}

	if ticket.Price < 0 {
		return errors.NewBadRequestErr("Price could not be less than 0")
	}

	return nil
}
