package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        string  `gorm:"primaryKey;size:36"`
	Name      string  `gorm:"size:100;not null"`
	Age       int     `gorm:"not null"`
	AddressID uint    `gorm:"not null"`
	Address   Address `gorm:"foreignKey:AddressID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Address struct {
	gorm.Model
	City  string `gorm:"size:100;not null" json:"city"`
	State string `gorm:"size:2;not null" json:"state"` 
}

type CreateUserRequest struct {
	Name    string         `json:"name" validate:"required,min=2,max=100"`
	Age     int            `json:"age" validate:"required,min=1,max=120"`
	Address AddressRequest `json:"address" validate:"required"`
}

type AddressRequest struct {
	City  string `json:"city" validate:"required,min=2,max=100"`
	State string `json:"state" validate:"required,len=2"`
}

type UpdateUserRequest struct {
	Name    string         `json:"name" validate:"omitempty,min=2,max=100"`
	Age     int            `json:"age" validate:"omitempty,min=1,max=120"`
	Address AddressRequest `json:"address" validate:"omitempty"`
}

type UserResponse struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	Age       int             `json:"age"`
	Address   AddressResponse `json:"address"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
}

type AddressResponse struct {
	City      string `json:"city"`
	State     string `json:"state"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:   u.ID,
		Name: u.Name,
		Age:  u.Age,
		Address: AddressResponse{
			City:      u.Address.City,
			State:     u.Address.State,
			CreatedAt: u.Address.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: u.Address.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func (r *CreateUserRequest) ToUser() User {
	return User{
		Name: r.Name,
		Age:  r.Age,
		Address: Address{
			City:  r.Address.City,
			State: r.Address.State,
		},
	}
}
