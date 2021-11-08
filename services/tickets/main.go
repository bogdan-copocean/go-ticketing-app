package main

import (
	"net/http"

	"github.com/bogdan-user/go-ticketing-app/services/tickets/app"
	"github.com/bogdan-user/go-ticketing-app/services/tickets/interfaces"
	"github.com/bogdan-user/go-ticketing-app/services/tickets/middlewares"
	"github.com/bogdan-user/go-ticketing-app/services/tickets/repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	ticketsRepository := repository.ConnectToMongo()
	ticketsService := app.NewTicketsService(ticketsRepository)
	ticketsHandler := interfaces.NewTicketsHandler(ticketsService)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.CurrentUserMiddleware)
		r.Post("/api/tickets", ticketsHandler.CreateTicket)
		r.Put("/api/tickets/{ticketId}", ticketsHandler.UpdateTicket)
	})

	r.Get("/api/tickets", ticketsHandler.GetAllTickets)
	r.Get("/api/tickets/{ticketId}", ticketsHandler.GetTicketById)
	http.ListenAndServe(":9001", r)

}
