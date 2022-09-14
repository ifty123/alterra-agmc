package database

import (
	"day2-task1/config"
	"day2-task1/models"
	"errors"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if e := config.DB.Where("deleted_at is NULL").Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetUserById(id int) (interface{}, error) {
	var user models.User

	if e := config.DB.Where("id = ?", id).Find(&user); e != nil {

		return nil, errors.New("User not found")
	}
	return user, nil
}

func SaveUser(user *models.User) error {
	//save user
	if result := config.DB.Select("name", "email", "password").Create(user); result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateUsers(user *models.User, id int) error {
	//save user
	if config.DB.Model(user).Select("name", "email", "password", "updated_at").Where("id = ?", id).Updates(user).RowsAffected == 0 {
		config.DB.Select("name", "email", "password").Create(user)
	}
	return nil
}

func DeleteUser(id int) error {
	var user models.User
	config.DB.Delete(&user, id)
	return nil
}
