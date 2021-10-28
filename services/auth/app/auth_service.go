package app

import (
	"net/http"

	"github.com/bogdan-user/go-ticketing-app/pkg/errors"
	"github.com/bogdan-user/go-ticketing-app/services/auth/domain"
	"github.com/bogdan-user/go-ticketing-app/services/auth/repository"
)

type AuthService interface {
	SignUp(*domain.User) (*domain.User, *errors.CustomErr)
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(authRepository repository.AuthRepository) AuthService {
	return &authService{authRepository: authRepository}
}

func (as *authService) SignUp(user *domain.User) (*domain.User, *errors.CustomErr) {
	foundUser, getErr := as.authRepository.GetUserByEmail(user.Email)

	if foundUser != nil {
		return nil, errors.NewBadRequestErr("email already exists")
	}
	if getErr != nil && getErr.StatusCode != http.StatusNotFound {
		return nil, getErr
	}

	createdUser, createdErr := as.authRepository.CreateUser(user)
	if createdErr != nil {
		return nil, createdErr
	}

	return createdUser, nil

}
