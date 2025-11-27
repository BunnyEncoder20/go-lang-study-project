package middlewares

import (
	"errors"
	"net/http"
)

var UnAuthrizedError = errors.New("invalid username or token")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc()
}
