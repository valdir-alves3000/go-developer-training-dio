package usecases

import (
	"github.com/google/uuid"

	"github.com/valdir-alves3000/go-developer-training-dio/projets/grandma-clients-api/internal/entities"
)

func generateID() string {
	return uuid.New().String()
}

func (uc *UserUsecaseImpl) Create(req *entities.CreateUserRequest) (*entities.UserResponse, error) {
	user := req.ToUser()
	user.ID = generateID()

	createdUser, err := uc.userRepo.Create(&user)
	if err != nil {
		return nil, err
	}

	response := createdUser.ToResponse()
	return &response, nil
}