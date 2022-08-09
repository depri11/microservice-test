package handler

import (
	"encoding/json"
	"net/http"

	"github.com/depri11/lolipad/logistic/src/interfaces"
	"github.com/depri11/lolipad/logistic/src/models"
)

type handler struct {
	service interfaces.LogisticService
}

func NewHandler(service interfaces.LogisticService) *handler {
	return &handler{service}
}

func (h *handler) GetAllData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, err := h.service.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (h *handler) GetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input models.InputLogisticData
	json.NewDecoder(r.Body).Decode(&input)

	result, err := h.service.GetData(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
}
