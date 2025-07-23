package usecases

import "github.com/valdir-alves3000/go-developer-training-dio/projets/grandma-clients-api/internal/entities"


func (uc *UserUsecaseImpl) Update(id string, req *entities.UpdateUserRequest) (*entities.UserResponse, error) {
	user, err := uc.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Age != 0 {
		user.Age = req.Age
	}
	if req.Address.City != "" {
		user.Address.City = req.Address.City
	}
	if req.Address.State != "" {
		user.Address.State = req.Address.State
	}

	updatedUser, err := uc.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	response := updatedUser.ToResponse()
	return &response, nil
}