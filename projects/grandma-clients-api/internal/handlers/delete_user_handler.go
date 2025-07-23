package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/valdir-alves3000/go-developer-training-dio/projets/grandma-clients-api/internal/repositories"
)

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.userUsecase.Delete(id)
	if err != nil {
		if err == repositories.ErrUserNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
