package app

import (
	"github.com/bogdan-user/go-ticketing-app/services/tickets/repository"
)

type TicketsService interface {
}

type ticketsService struct {
	ticketsRepository repository.TicketsRepository
}

func NewTicketsService(ticketsRepository repository.TicketsRepository) TicketsService {
	return &ticketsService{ticketsRepository: ticketsRepository}
}
