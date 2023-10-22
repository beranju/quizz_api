package controller

import (
	"main/libs/database"
	"main/models/response"
	"net/http"

	"github.com/labstack/echo"
)

func GetUserController(c echo.Context) error {
	user, err := database.GetUser(c)
	if err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Get User Detail Failed",
			Error:   err.Error(),
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	response := response.SuccessResponse{
		Status:  "success",
		Message: "Get User Detail successfully",
		Data:    user,
	}

	return c.JSON(http.StatusOK, response)

}

func PostUserController(c echo.Context) error {
	user, err := database.PostUser(c)
	if err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "User Creation Failed",
			Error:   err.Error(),
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	response := response.SuccessResponse{
		Status:  "success",
		Message: "User Created successfully",
		Data:    user,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateUserController(c echo.Context) error {
	user, err := database.UpdateUser(c)
	if err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "User Update Failed",
			Error:   err.Error(),
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	response := response.SuccessResponse{
		Status:  "success",
		Message: "User Updated successfully",
		Data:    user,
	}

	return c.JSON(http.StatusOK, response)
}
