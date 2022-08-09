package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/depri11/lolipad/logistic/src/helpers"
)

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		headerToken := r.Header.Get("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			helpers.ResponseJSON("Failed", http.StatusBadRequest, "error", "invalid header type").Send(w)
			return
		}

		token := strings.Replace(headerToken, "Bearer ", "", -1)

		urlCheckToken := "http://localhost:3000/auth/check-token"

		var jsonData = map[string]string{"token": token}
		b, _ := json.Marshal(jsonData)
		request, err := http.NewRequest("GET", urlCheckToken, bytes.NewBuffer(b))
		if err != nil {
			helpers.ResponseJSON("Internal Server Error", http.StatusInternalServerError, "error", err.Error()).Send(w)
			return
		}

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			helpers.ResponseJSON("Internal Server Error", http.StatusInternalServerError, "error", err.Error()).Send(w)
			return
		}
		// defer response.Body.Close()

		if response.StatusCode != 200 {
			helpers.ResponseJSON("Unauthorized", http.StatusUnauthorized, "error", nil).Send(w)
			return
		}

		next.ServeHTTP(w, r)
	}
}
