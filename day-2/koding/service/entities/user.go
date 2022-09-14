package entities

import (
	"day2-task1/models"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDTO struct {
	Name     string
	Email    string
	Password string
}

func ToUserDTO(user *User) *UserDTO {
	return &UserDTO{
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
