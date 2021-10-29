package interfaces

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/bogdan-user/go-ticketing-app/pkg/crypto"
	"github.com/bogdan-user/go-ticketing-app/pkg/errors"
	"github.com/bogdan-user/go-ticketing-app/services/auth/app"
	"github.com/bogdan-user/go-ticketing-app/services/auth/domain"
)

type AuthHandler interface {
	GetCurrentUser(http.ResponseWriter, *http.Request)
	SignIn(http.ResponseWriter, *http.Request)
	SignOut(http.ResponseWriter, *http.Request)
	SignUp(http.ResponseWriter, *http.Request)
}

type authHandler struct {
	authService app.AuthService
}

func NewAuthHandler(authService app.AuthService) AuthHandler {
	return &authHandler{authService: authService}
}

func (ah *authHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cookie, cookieErr := r.Cookie("jwt")
	if cookieErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(errors.NewBadRequestErr("no token provided"))
		w.Write(res)
		return
	}
	crypto.VerifyJWTToken(cookie.Value)

}

func (ah *authHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	user := &domain.User{}
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if err := json.Unmarshal(body, user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(errors.NewBadRequestErr("bad json"))
		w.Write(res)
		return
	}

	foundUser, serviceErr := ah.authService.SignIn(user)
	if serviceErr != nil {
		w.WriteHeader(serviceErr.StatusCode)
		res, _ := json.Marshal(serviceErr)
		w.Write(res)
		return
	}

	accessToken, genErr := crypto.GenerateJWTToken(foundUser)
	if genErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal("could not generate token")
		w.Write(res)
		return
	}

	cookie := http.Cookie{Name: "jwt", Value: accessToken}
	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(foundUser)
	w.Write(res)

}

func (ah *authHandler) SignOut(rw http.ResponseWriter, req *http.Request) {
}

func (ah *authHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	user := &domain.User{}
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if err := json.Unmarshal(body, user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(errors.NewBadRequestErr("bad json"))
		w.Write(res)
		return
	}

	createdUser, serviceErr := ah.authService.SignUp(user)
	if serviceErr != nil {
		w.WriteHeader(serviceErr.StatusCode)
		res, _ := json.Marshal(serviceErr)
		w.Write(res)
		return
	}

	accessToken, genErr := crypto.GenerateJWTToken(createdUser)
	if genErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal("could not generate token")
		w.Write(res)
		return
	}

	cookie := http.Cookie{Name: "jwt", Value: accessToken}
	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusCreated)
	res, _ := json.Marshal(createdUser)
	w.Write(res)

}
