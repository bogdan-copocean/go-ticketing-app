package errors

import "net/http"

type CustomErr struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func NewBadRequestErr(message string) *CustomErr {
	return &CustomErr{Message: message, StatusCode: http.StatusBadRequest}
}

func NewInternalServerErr(message string) *CustomErr {
	return &CustomErr{Message: message, StatusCode: http.StatusInternalServerError}
}

func NewNotFoundErr(message string) *CustomErr {
	return &CustomErr{Message: message, StatusCode: http.StatusNotFound}
}

func NewUnauthorizedErr(message string) *CustomErr {
	return &CustomErr{Message: message, StatusCode: http.StatusUnauthorized}
}
