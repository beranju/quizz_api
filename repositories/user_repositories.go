package repositories

import (
	"main/config"
	"main/models"
)

func FindUserById(user *models.User, id string) error {
	result := config.DB.First(&user, id)
	return result.Error
}

func SaveUser(user *models.User) error {
	result := config.DB.Save(&user)
	return result.Error
}

func Register(user *models.User) error {
	result := config.DB.Create(&user)
	return result.Error
}

func Login(user *models.User) error {
	result := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user)
	return result.Error
}
