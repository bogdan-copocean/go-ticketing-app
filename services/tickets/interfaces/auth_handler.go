package interfaces

import (
	"github.com/bogdan-user/go-ticketing-app/services/tickets/app"
)

type TicketsHandler interface {
}

type ticketsHandler struct {
	authService app.TicketsService
}

func NewTicketsHandler(authService app.TicketsService) TicketsHandler {
	return &ticketsHandler{authService: authService}
}
