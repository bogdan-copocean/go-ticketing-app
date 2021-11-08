package app

import (
	"github.com/bogdan-user/go-ticketing-app/pkg/errors"
	"github.com/bogdan-user/go-ticketing-app/services/tickets/domain"
	"github.com/bogdan-user/go-ticketing-app/services/tickets/repository"
)

type TicketsService interface {
	CreateTicket(*domain.Ticket) (*domain.Ticket, *errors.CustomErr)
	GetTicketById(string) (*domain.Ticket, *errors.CustomErr)
	GetAllTickets() ([]*domain.Ticket, *errors.CustomErr)
}

type ticketsService struct {
	ticketsRepository repository.TicketsRepository
}

func NewTicketsService(ticketsRepository repository.TicketsRepository) TicketsService {
	return &ticketsService{ticketsRepository: ticketsRepository}
}

func (ts *ticketsService) CreateTicket(ticket *domain.Ticket) (*domain.Ticket, *errors.CustomErr) {
	if errValidate := ticket.ValidateTicket(); errValidate != nil {
		return nil, errValidate
	}

	ticket, createErr := ts.ticketsRepository.CreateTicket(ticket)
	if createErr != nil {
		return nil, createErr
	}

	return ticket, nil

}

func (ts *ticketsService) GetTicketById(id string) (*domain.Ticket, *errors.CustomErr) {

	ticket, getErr := ts.ticketsRepository.GetTicketById(id)
	if getErr != nil {
		return nil, getErr
	}

	ticket.UserId = ""

	return ticket, nil
}

func (ts *ticketsService) GetAllTickets() ([]*domain.Ticket, *errors.CustomErr) {

	tickets, getErr := ts.ticketsRepository.GetAllTickets()
	if getErr != nil {
		return nil, getErr
	}

	for _, ticket := range tickets {
		ticket.Id = ""
	}

	return tickets, nil
}
