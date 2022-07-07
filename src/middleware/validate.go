package middleware

import (
	"net/http"
	"strings"

	"github.com/Irsad99/LectronicApp/src/helpers"
)

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		headerToken := r.Header.Get("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			helpers.New("invalid header type", 401, true).Send(w)
			return
		}

		token := strings.Replace(headerToken, "Bearer ", "", -1)
		checkToken, err := helpers.CheckToken(token)
		if err != nil {
			helpers.New("invalid token", 401, true).Send(w)
			return
		}

		r.Header.Set("user_id", checkToken.Id)
		r.Header.Set("role", checkToken.Role)

		next.ServeHTTP(w, r)
	}
}

func CheckRoleAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("role")
		if role != "admin" {
			http.Error(w, "you are not admin", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
