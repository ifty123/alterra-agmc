package repositories

import "day2-task1/models"

type UserDatabase interface {
	GetUsers() (interface{}, error)
	GetUserById(id string) (interface{}, error)
	SaveUser(user *models.User) error
	UpdateUsers(user *models.User, id int) error
	DeleteUser(id int) error
	LoginUser(user *models.User) (interface{}, error)
}
