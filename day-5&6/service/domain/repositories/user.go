package repositories

import models "day2-task1/internal/model"

type UserDatabase interface {
	GetUsers() (interface{}, error)
	GetUserById(id string) (interface{}, error)
	SaveUser(user *models.User) error
	UpdateUsers(user *models.User, id int) error
	DeleteUser(id int) error
	LoginUser(user *models.User) (interface{}, error)
}
