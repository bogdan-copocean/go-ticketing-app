package interfaces

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/bogdan-user/go-ticketing-app/pkg/errors"
	"github.com/bogdan-user/go-ticketing-app/services/tickets/app"
	"github.com/bogdan-user/go-ticketing-app/services/tickets/domain"
)

type TicketsHandler interface {
	CreateTicket(http.ResponseWriter, *http.Request)
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
