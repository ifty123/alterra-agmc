package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// Id       string `json:"id_user" form:"id_user"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
