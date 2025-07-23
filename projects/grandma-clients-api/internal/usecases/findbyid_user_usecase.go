package usecases

import "github.com/valdir-alves3000/go-developer-training-dio/projets/grandma-clients-api/internal/entities"

func (uc *UserUsecaseImpl) FindByID(id string) (*entities.UserResponse, error) {
	user, err := uc.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := user.ToResponse()
	return &response, nil
}
