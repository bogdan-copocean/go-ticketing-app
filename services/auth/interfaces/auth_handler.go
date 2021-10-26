package interfaces

import (
	"net/http"

	"github.com/labstack/echo"
)

type AuthHandler interface {
	GetCurrentUser(echo.Context) error
}

type authHandler struct{}

func NewAuthHandler() AuthHandler {
	return &authHandler{}
}

func (ah *authHandler) GetCurrentUser(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]string{"accessed": "granted"})
}
