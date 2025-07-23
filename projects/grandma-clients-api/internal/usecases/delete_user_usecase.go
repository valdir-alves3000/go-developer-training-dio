package usecases

import "github.com/valdir-alves3000/go-developer-training-dio/projets/grandma-clients-api/internal/repositories"

func (uc *UserUsecaseImpl) Delete(id string) error {
	exists, err := uc.userRepo.Exists(id)
	if err != nil {
		return err
	}
	if !exists {
		return repositories.ErrUserNotFound
	}

	return uc.userRepo.Delete(id)
}