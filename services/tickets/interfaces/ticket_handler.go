package interfaces

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/bogdan-user/go-ticketing-app/pkg/errors"
	"github.com/bogdan-user/go-ticketing-app/services/tickets/app"
	"github.com/bogdan-user/go-ticketing-app/services/tickets/domain"
	"github.com/bogdan-user/go-ticketing-app/services/tickets/middlewares"
	"github.com/go-chi/chi"
)

type TicketsHandler interface {
	CreateTicket(http.ResponseWriter, *http.Request)
	GetTicketById(http.ResponseWriter, *http.Request)
	GetAllTickets(w http.ResponseWriter, r *http.Request)
	UpdateTicket(w http.ResponseWriter, r *http.Request)
}

type ticketsHandler struct {
	authService app.TicketsService
}

func NewTicketsHandler(authService app.TicketsService) TicketsHandler {
	return &ticketsHandler{authService: authService}
}

func (th *ticketsHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ticket := &domain.Ticket{}
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err := json.Unmarshal(body, ticket); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(errors.NewBadRequestErr("bad json"))
		w.Write(res)
		return
	}

	reqCtx := r.Context()
	claims := reqCtx.Value(middlewares.CurrentUser)

	claimsMap := claims.(map[string]interface{})
	ticket.UserId = claimsMap["id"].(string)

	createdTicket, createErr := th.authService.CreateTicket(ticket)
	if createErr != nil {
		w.WriteHeader(createErr.StatusCode)
		res, _ := json.Marshal(createErr)
		w.Write(res)
		return
	}

	w.WriteHeader(http.StatusCreated)
	res, _ := json.Marshal(createdTicket)
	w.Write(res)

}

func (th *ticketsHandler) GetTicketById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ticketId := chi.URLParam(r, "ticketId")

	foundTicket, getErr := th.authService.GetTicketById(ticketId)
	if getErr != nil {
		w.WriteHeader(getErr.StatusCode)
		res, _ := json.Marshal(getErr)
		w.Write(res)
		return
	}

	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(foundTicket)
	w.Write(res)

}

func (th *ticketsHandler) GetAllTickets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	foundTicket, getErr := th.authService.GetAllTickets()
	if getErr != nil {
		w.WriteHeader(getErr.StatusCode)
		res, _ := json.Marshal(getErr)
		w.Write(res)
		return
	}

	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(foundTicket)
	w.Write(res)

}

func (th *ticketsHandler) UpdateTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ticketDetail := &domain.Ticket{}
	ticketId := chi.URLParam(r, "ticketId")

	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err := json.Unmarshal(body, ticketDetail); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(errors.NewBadRequestErr("bad json"))
		w.Write(res)
		return
	}

	reqCtx := r.Context()
	claims := reqCtx.Value(middlewares.CurrentUser)

	claimsMap := claims.(map[string]interface{})
	userRequestId := claimsMap["id"].(string)

	updatedTicket, updateErr := th.authService.UpdateTicket(userRequestId, ticketId, ticketDetail)

	if updateErr != nil {
		w.WriteHeader(updateErr.StatusCode)
		res, _ := json.Marshal(updateErr)
		w.Write(res)
		return
	}

	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(updatedTicket)
	w.Write(res)

}
