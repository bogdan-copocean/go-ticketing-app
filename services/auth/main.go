package main

import (
	"github.com/bogdan-user/go-ticketing-app/services/auth/app"
	"github.com/bogdan-user/go-ticketing-app/services/auth/interfaces"
	"github.com/bogdan-user/go-ticketing-app/services/auth/repository"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

var mongoCli *mongo.Client

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}    ${method}   ${uri} \t ${status} -> ${latency_human}\n",
	}))

	authRepository := repository.ConnectToMongo()
	authService := app.NewAuthService(authRepository)
	authHandler := interfaces.NewAuthHandler(authService)

	e.GET("/api/users/currentuser", authHandler.GetCurrentUser)
	e.POST("/api/users/signin", authHandler.SignIn)
	e.POST("/api/users/signout", authHandler.SignOut)
	e.POST("/api/users/signup", authHandler.SignUp)

	e.Start(":3000")
}
