package controller

import (
	"main/middlewares"
	"main/models"
	"main/models/response"
	"main/repositories"
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

	err := repositories.Login(&user)
	if err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Username and password are incorrect",
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}
	token, err := middlewares.GenerateToken(int(user.ID), user.Name)
	if err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Something went wrong, when generating token",
		}
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	var userResponse response.UserResponse
	userResponse.Id = int(user.ID)
	userResponse.Name = user.Name
	userResponse.Email = user.Email
	userResponse.ImageUrl = user.ImageUrl
	userResponse.Token = token
	userResponse.CreatedAt = user.CreatedAt
	userResponse.UpdatedAt = user.UpdatedAt

	response := response.BaseResponse{
		Status:  "success",
		Message: "Login successfully",
		Data:    userResponse,
	}

	return ctx.JSON(http.StatusOK, response)
}
