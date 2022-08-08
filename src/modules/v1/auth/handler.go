package auth

import (
	"encoding/json"
	"net/http"

	"github.com/depri11/lolipad/src/interfaces"
	"github.com/depri11/lolipad/src/models"
)

type handler struct {
	service interfaces.UserService
}

func NewHandler(service interfaces.UserService) *handler {
	return &handler{service}
}

func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	res := h.service.Register(&user)

	json.NewEncoder(w).Encode(res)
}
