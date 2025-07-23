package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/valdir-alves3000/go-developer-training-dio/projets/grandma-clients-api/internal/entities"
)

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req entities.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdUser, err := h.userUsecase.Create(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}
