package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       string `json:"name" form:"id_user"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Books struct {
	Id        int    `json:"id_book"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}
