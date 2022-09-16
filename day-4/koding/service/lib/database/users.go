package database

import (
	"day2-task1/config"
	"day2-task1/domain/repositories"
	"day2-task1/models"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type UsersRepository struct {
	Connection *gorm.DB
}

func NewUsersRepository(db *gorm.DB) repositories.UserDatabase {
	return &UsersRepository{
		Connection: db,
	}
}

func (u *UsersRepository) GetUsers() (interface{}, error) {
	var users []models.User

	if e := config.DB.Where("deleted_at is NULL").Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func (u *UsersRepository) GetUserById(id string) (interface{}, error) {
	idConvert, _ := strconv.Atoi(id)
	var user []models.User

	if e := config.DB.Where("id = ? and deleted_at is NULL", idConvert).Find(&user).Error; e != nil {
		return nil, errors.New("User not found")
	}
	return user, nil
}

func (u *UsersRepository) SaveUser(user *models.User) error {
	//save user
	if result := config.DB.Select("name", "email", "password").Create(user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UsersRepository) UpdateUsers(user *models.User, id int) error {
	//save user
	if config.DB.Model(user).Select("name", "email", "password", "updated_at").Where("id = ?", id).Updates(user).RowsAffected == 0 {
		config.DB.Select("name", "email", "password").Create(user)
	}
	return nil
}

func (u *UsersRepository) DeleteUser(id int) error {
	var user models.User
	if e := config.DB.Delete(&user, id).Error; e != nil {
		return errors.New("cant delete user")
	}
	return nil
}

func (u *UsersRepository) LoginUser(user *models.User) (interface{}, error) {
	if err := config.DB.Where("email = ? and password = ? and deleted_at is NULL ", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
