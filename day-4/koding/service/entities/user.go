package entities

import (
	"day2-task1/models"
)

type User struct {
	id       int    `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserDTO struct {
	Id       int
	Name     string
	Email    string
	Password string
}

func ToUserDTO(user *User) *UserDTO {
	return &UserDTO{
		Id:       user.id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func AssembUserDTO(user *UserDTO) *models.User {
	return &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
