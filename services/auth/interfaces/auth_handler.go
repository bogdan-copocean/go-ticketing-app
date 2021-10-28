package interfaces

import (
	"net/http"

	"github.com/bogdan-user/go-ticketing-app/pkg/errors"
	"github.com/bogdan-user/go-ticketing-app/services/auth/app"
	"github.com/bogdan-user/go-ticketing-app/services/auth/domain"
	"github.com/labstack/echo"
)

type AuthHandler interface {
	GetCurrentUser(echo.Context) error
	SignIn(echo.Context) error
	SignOut(echo.Context) error
	SignUp(echo.Context) error
}

type authHandler struct {
	authService app.AuthService
}

func NewAuthHandler(authService app.AuthService) AuthHandler {
	return &authHandler{authService: authService}
}

func (ah *authHandler) GetCurrentUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"accessed": "granted"})
}

func (ah *authHandler) SignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"accessed": "granted"})
}

func (ah *authHandler) SignOut(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"accessed": "granted"})
}

func (ah *authHandler) SignUp(c echo.Context) error {
	user := &domain.User{}

	if err := c.Bind(user); err != nil {
		return err
	}

	if user.Email == "" {
		return c.JSON(http.StatusBadRequest, errors.NewBadRequestErr("email not provided"))
	}

	if user.Password == "" {
		return c.JSON(http.StatusBadRequest, errors.NewBadRequestErr("password not provided"))
	}

	createdUser, serviceErr := ah.authService.SignUp(user)
	if serviceErr != nil {
		return c.JSON(serviceErr.StatusCode, serviceErr)
	}

	return c.JSON(http.StatusOK, createdUser)
}
