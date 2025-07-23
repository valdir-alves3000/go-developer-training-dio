package handlers

import (
	"github.com/valdir-alves3000/go-developer-training-dio/projets/grandma-clients-api/internal/usecases"
)

type UserHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHandler(userUsecase usecases.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}
