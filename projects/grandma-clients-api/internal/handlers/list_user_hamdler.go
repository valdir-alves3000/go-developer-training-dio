package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/valdir-alves3000/go-developer-training-dio/projets/grandma-clients-api/internal/entities"
)

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userUsecase.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	responses := append([]entities.UserResponse{}, users...)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}
