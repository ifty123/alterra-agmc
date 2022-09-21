package repository

import (
	config "day2-task1/database"
	"day2-task1/internal/model"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type UserDatabase interface {
	GetUsers() (interface{}, error)
	GetUserById(id string) (interface{}, error)
	SaveUser(user *model.User) error
	UpdateUsers(user *model.User, id int) error
	DeleteUser(id int) error
	LoginUser(user *model.User) (interface{}, error)
}

type UsersRepository struct {
	Connection *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UserDatabase {
	return &UsersRepository{
		Connection: db,
	}
}

func (u *UsersRepository) GetUsers() (interface{}, error) {
	var users []model.User

	if e := config.DB.Where("deleted_at is NULL").Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func (u *UsersRepository) GetUserById(id string) (interface{}, error) {
	idConvert, _ := strconv.Atoi(id)
	var user []model.User

	if e := config.DB.Where("id = ? and deleted_at is NULL", idConvert).Find(&user).Error; e != nil {
		return nil, errors.New("User not found")
	}
	return user, nil
}

func (u *UsersRepository) SaveUser(user *model.User) error {
	//save user
	if result := config.DB.Select("name", "email", "password").Create(user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UsersRepository) UpdateUsers(user *model.User, id int) error {
	//save user
	if config.DB.Model(user).Select("name", "email", "password", "updated_at").Where("id = ?", id).Updates(user).RowsAffected == 0 {
		config.DB.Select("name", "email", "password").Create(user)
	}
	return nil
}

func (u *UsersRepository) DeleteUser(id int) error {
	var user model.User
	if e := config.DB.Delete(&user, id).Error; e != nil {
		return errors.New("cant delete user")
	}
	return nil
}

func (u *UsersRepository) LoginUser(user *model.User) (interface{}, error) {
	if err := config.DB.Where("email = ? and password = ? and deleted_at is NULL ", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
