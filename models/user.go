package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"name;not null"`
	Email    string `json:"email" gorm:"email;unique;not null"`
	Password string `json:"password" gorm:"password;not null"`
	ImageUrl string `json:"image_url" gorm:"image_url"`
}
