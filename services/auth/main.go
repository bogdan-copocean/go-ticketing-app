package main

import (
	"net/http"

	"github.com/bogdan-user/go-ticketing-app/services/auth/app"
	"github.com/bogdan-user/go-ticketing-app/services/auth/interfaces"
	"github.com/bogdan-user/go-ticketing-app/services/auth/middlewares"
	"github.com/bogdan-user/go-ticketing-app/services/auth/repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	authRepository := repository.ConnectToMongo()
	authService := app.NewAuthService(authRepository)
	authHandler := interfaces.NewAuthHandler(authService)

	// protected routes
	r.Group(func(r chi.Router) {
		r.Use(middlewares.CurrentUserMiddleware)
		r.Get("/api/users/currentuser", authHandler.GetCurrentUser)
		r.Post("/api/users/signout", authHandler.SignOut)
	})

	// public routes
	r.Post("/api/users/signin", authHandler.SignIn)
	r.Post("/api/users/signup", authHandler.SignUp)

	http.ListenAndServe(":9000", r)

}
