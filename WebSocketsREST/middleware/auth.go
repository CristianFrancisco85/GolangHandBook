package middleware

import (
	"net/http"
	"rest_ws/server"
	"rest_ws/utils"
)

var (
	NO_AUTH_NEEDED = []string{
		"login",
		"signup",
	}
)

func CheckAuthMiddleware(s server.Server) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			_, err := utils.GetTokenFromHeader(r, s.Config().JWTSecret)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
