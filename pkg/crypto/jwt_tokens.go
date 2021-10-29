package crypto

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/bogdan-user/go-ticketing-app/pkg/errors"
	"github.com/bogdan-user/go-ticketing-app/services/auth/domain"
	"github.com/golang-jwt/jwt"
)

func GenerateJWTToken(user *domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
		"iat":   strconv.Itoa(int(time.Now().Unix())),
	})
	accessToken, signErr := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if signErr != nil {
		return "", signErr
	}

	return accessToken, nil
}

func VerifyJWTToken(token string) *errors.CustomErr {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		if err.Error() == "Token is expired" {
			return errors.NewBadRequestErr("Token is expired")
		}
		return errors.NewBadRequestErr("Could not parse token")
	}

	_, ok := parsedToken.Claims.(jwt.MapClaims)
	if ok && parsedToken.Valid {
		fmt.Println("valid token")
		return nil

	}

	return nil
}
