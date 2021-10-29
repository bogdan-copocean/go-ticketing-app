package app

import (
	"github.com/bogdan-user/go-ticketing-app/pkg/errors"
	"github.com/bogdan-user/go-ticketing-app/services/auth/domain"
	"github.com/bogdan-user/go-ticketing-app/services/auth/repository"
)

type AuthService interface {
	SignUp(*domain.User) (*domain.User, *errors.CustomErr)
	SignIn(*domain.User) (*domain.User, *errors.CustomErr)
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(authRepository repository.AuthRepository) AuthService {
	return &authService{authRepository: authRepository}
}

func (as *authService) SignUp(user *domain.User) (*domain.User, *errors.CustomErr) {
	if validateUserErr := user.ValidateFields(); validateUserErr != nil {
		return nil, errors.NewBadRequestErr(validateUserErr.Error())
	}

	foundUser, getErr := as.authRepository.GetUserByEmail(user)
	if foundUser != nil {
		return nil, errors.NewBadRequestErr("email already exists")
	}
	if getErr != nil && getErr.StatusCode != 404 {
		return nil, getErr
	}

	createdUser, createdErr := as.authRepository.CreateUser(user)
	if createdErr != nil {
		return nil, createdErr
	}

	// omits password from being marshaled and sent to client
	createdUser.Password = ""

	return createdUser, nil

}

func (as *authService) SignIn(user *domain.User) (*domain.User, *errors.CustomErr) {
	if validateUserErr := user.ValidateFields(); validateUserErr != nil {
		return nil, errors.NewBadRequestErr(validateUserErr.Error())
	}

	foundUser, getErr := as.authRepository.GetUserByEmail(user)
	if getErr != nil {
		return nil, getErr
	}

	// omits password from being marshaled and sent to client
	foundUser.Password = ""

	return foundUser, nil

}
