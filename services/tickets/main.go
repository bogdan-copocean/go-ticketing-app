package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// ticketsRepository := repository.ConnectToMongo()
	// ticketsService := app.NewTicketsService(ticketsRepository)
	// ticketsHandler := interfaces.NewTicketsHandler(ticketsService)

	http.ListenAndServe(":9001", r)

}
