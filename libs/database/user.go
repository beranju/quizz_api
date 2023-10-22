package database

import (
	"main/config"
	"main/models"

	"github.com/labstack/echo"
)

func GetUser(c echo.Context) (interface{}, error) {
	id := c.Param("id")

	var users []*models.User
	if err := config.DB.Find(&users, id).Error; err != nil {
		return nil, err
	}
	user := users[0]
	return user, nil
}

func PostUser(c echo.Context) (interface{}, error) {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return nil, err
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	userCreated := &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		ImageUrl: user.ImageUrl,
	}

	return userCreated, nil
}

func UpdateUser(c echo.Context) (interface{}, error) {
	id := c.Param("id")
	dataUser := new(models.User)

	if err := c.Bind(dataUser); err != nil {
		return nil, err
	}

	existingUser := new(models.User)
	if err := config.DB.First(&existingUser, id).Error; err != nil {
		return nil, err
	}

	existingUser.Name = dataUser.Name
	existingUser.Email = dataUser.Email
	existingUser.Password = dataUser.Password
	existingUser.ImageUrl = dataUser.ImageUrl

	if err := config.DB.Save(&existingUser).Error; err != nil {
		return nil, err
	}
	return existingUser, nil
}
