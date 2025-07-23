package usecases

import (
	"github.com/valdir-alves3000/go-developer-training-dio/projets/grandma-clients-api/internal/entities"
	"github.com/valdir-alves3000/go-developer-training-dio/projets/grandma-clients-api/internal/repositories"
)

type UserUsecase interface {
	Create(user *entities.CreateUserRequest) (*entities.UserResponse, error)
	FindByID(id string) (*entities.UserResponse, error)
	FindAll() ([]entities.UserResponse, error)
	Update(id string, user *entities.UpdateUserRequest) (*entities.UserResponse, error)
	Delete(id string) error
}

type UserUsecaseImpl struct {
	userRepo repositories.UserRepository
}

func NewUserUsecase(userRepo repositories.UserRepository) *UserUsecaseImpl {
	return &UserUsecaseImpl{userRepo: userRepo}
}
