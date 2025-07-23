package repositories

import (
	"errors"

	"github.com/valdir-alves3000/go-developer-training-dio/projets/grandma-clients-api/internal/entities"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(user *entities.User) (*entities.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindByID(id string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Preload("Address").First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindAll() ([]entities.User, error) {
	var users []entities.User
	if err := r.db.Preload("Address").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepositoryImpl) Update(user *entities.User) (*entities.User, error) {
	// Start transaction
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Update address first
	if err := tx.Model(&user.Address).Updates(user.Address).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Then update user
	if err := tx.Model(user).Updates(map[string]interface{}{
		"name": user.Name,
		"age":  user.Age,
	}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) Delete(id string) error {
	// Find user first to get address ID
	user, err := r.FindByID(id)
	if err != nil {
		return err
	}

	// Start transaction
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Delete user
	if err := tx.Delete(&entities.User{}, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete address
	if err := tx.Delete(&entities.Address{}, "id = ?", user.AddressID).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	return tx.Commit().Error
}

func (r *UserRepositoryImpl) Exists(id string) (bool, error) {
	var count int64
	if err := r.db.Model(&entities.User{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
