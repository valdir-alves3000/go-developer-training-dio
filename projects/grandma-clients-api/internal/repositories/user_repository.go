package repositories

import "github.com/valdir-alves3000/go-developer-training-dio/projets/grandma-clients-api/internal/entities"

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
	FindByID(id string) (*entities.User, error)
	FindAll() ([]entities.User, error)
	Update(user *entities.User) (*entities.User, error)
	Delete(id string) error
	Exists(id string) (bool, error)
}
