package handler

import (
	"encoding/json"
	"net/http"

	"github.com/depri11/lolipad/auth/src/interfaces"
	"github.com/depri11/lolipad/auth/src/models"
)

type handler struct {
	service interfaces.UserService
}

func NewHandler(service interfaces.UserService) *handler {
	return &handler{service}
}

func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input models.InputRegisterUser
	json.NewDecoder(r.Body).Decode(&input)

	result, err := h.service.Register(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input models.InputLoginUser
	json.NewDecoder(r.Body).Decode(&input)

	result, err := h.service.Login(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (h *handler) CheckToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input models.InputCheckToken
	json.NewDecoder(r.Body).Decode(&input)

	result, err := h.service.CheckToken(input.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
}
