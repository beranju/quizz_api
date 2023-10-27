package controller

import (
	"main/middlewares"
	"main/models"
	"main/models/response"
	"main/repositories"
	"main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserDetailController(c echo.Context) error {
	var user models.User
	id := c.Param("id")

	err := repositories.FindUserById(&user, id)
	if err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "User not found",
		}
		return echo.NewHTTPError(http.StatusNotFound, response)
	}

	response := response.BaseResponse{
		Status:  "success",
		Message: "Get User Detail successfully",
		Data:    user,
	}

	return c.JSON(http.StatusOK, response)

}

func RegisterController(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Cek the user data",
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	hash, _ := utils.HashAndSalt([]byte(user.Password))
	user.Password = hash

	err := repositories.Register(&user)
	if err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "User Creation Failed",
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	response := response.BaseResponse{
		Status:  "success",
		Message: "User Created successfully",
		Data:    user,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	var existingUser models.User
	if err := c.Bind(&user); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Cek the user data",
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	err := repositories.FindUserById(&existingUser, id)
	if err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "User not found",
		}
		return echo.NewHTTPError(http.StatusNotFound, response)
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Password = user.Password
	existingUser.ImageUrl = user.ImageUrl

	error := repositories.SaveUser(&existingUser)
	if error != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "User update failed",
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}

	response := response.BaseResponse{
		Status:  "success",
		Message: "User Updated successfully",
	}

	return c.JSON(http.StatusOK, response)
}

func LoginController(ctx echo.Context) error {
	var user models.User
	ctx.Bind(&user)

	var existingUser models.User
	err := repositories.Login(&existingUser, user.Email)
	if err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Username and password are incorrect",
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	if passwordMatch := utils.ComparePassword(existingUser.Password, []byte(user.Password)); passwordMatch == false {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Password mismatch",
		}
		return ctx.JSON(http.StatusUnauthorized, response)
	}

	token, err := middlewares.GenerateToken(int(existingUser.ID), existingUser.Name)
	if err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Something went wrong, when generating token",
		}
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	var userResponse response.UserResponse
	userResponse.Id = int(existingUser.ID)
	userResponse.Name = existingUser.Name
	userResponse.Email = existingUser.Email
	userResponse.ImageUrl = existingUser.ImageUrl
	userResponse.Token = token
	userResponse.CreatedAt = existingUser.CreatedAt
	userResponse.UpdatedAt = existingUser.UpdatedAt

	response := response.BaseResponse{
		Status:  "success",
		Message: "Login successfully",
		Data:    userResponse,
	}

	return ctx.JSON(http.StatusOK, response)
}
