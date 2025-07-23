package usecases

import "github.com/valdir-alves3000/go-developer-training-dio/projets/grandma-clients-api/internal/entities"

func (uc *UserUsecaseImpl) FindAll() ([]entities.UserResponse, error) {
	users, err := uc.userRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []entities.UserResponse
	for _, user := range users {
		responses = append(responses, user.ToResponse())
	}

	return responses, nil
}
