package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/bogdan-user/go-ticketing-app/pkg/crypto"
	"github.com/bogdan-user/go-ticketing-app/pkg/errors"
)

type contextKey int

const (
	CurrentUser contextKey = iota
)

func CurrentUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// extract jwt token from cookie
		cookie, cookieErr := r.Cookie("jwt")
		if cookieErr != nil {
			w.WriteHeader(http.StatusUnauthorized)
			res, _ := json.Marshal(errors.NewUnauthorizedErr("not authorized"))
			w.Write(res)
			return
		}

		// verify payload and extract claims
		claims, tokenErr := crypto.VerifyJWTToken(cookie.Value)

		if tokenErr != nil {
			w.WriteHeader(tokenErr.StatusCode)
			res, _ := json.Marshal(tokenErr)
			w.Write(res)
			return
		}

		// write claims to request's context and call the next middleware
		ctx := context.WithValue(r.Context(), CurrentUser, claims)
		next.ServeHTTP(w, r.WithContext(ctx))

	})

}
