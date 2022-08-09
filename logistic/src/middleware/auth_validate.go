package middleware

import (
	"net/http"
	"strings"

	"github.com/depri11/lolipad/logistic/src/helpers"
)

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		headerToken := r.Header.Get("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			helpers.ResponseJSON("invalid header type", 401, "error", nil)
			return
		}

		token := strings.Replace(headerToken, "Bearer ", "", -1)
		checkToken, err := helpers.CheckToken(token)
		if err != nil {
			helpers.ResponseJSON("invalid token", 401, "error", nil)
			return
		}

		r.Header.Set("user_id", checkToken.Id)

		next.ServeHTTP(w, r)
	}
}
