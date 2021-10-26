package main

import (
	"github.com/bogdan-user/go-ticketing-app/services/auth/interfaces"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}    ${method}   ${uri} \t ${status} -> ${latency_human}\n",
	}))

	authHandler := interfaces.NewAuthHandler()

	e.GET("/api/users/currentuser", authHandler.GetCurrentUser)

	e.Logger.Fatal(e.Start(":3000"))
}
