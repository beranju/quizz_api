package controller

import (
	"main/config"
	"main/models"
	"main/models/response"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func GetQuizResultController(ctx echo.Context) error {
	db := config.DB
	quizId := ctx.Param("quiz_id")
	userId := ctx.Param("user_id")
	quizResult := new(models.UserQuizResult)

	if err := db.Where("quiz_id = ? AND user_id = ?", quizId, userId).First(quizResult).Error; err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve Quiz result",
			Error:   err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response := response.SuccessResponse{
		Status:  "success",
		Message: "Quiz result retrieved successfully",
		Data:    quizResult,
	}
	return ctx.JSON(http.StatusOK, response)
}

func AddQuizResultController(ctx echo.Context) error {
	db := config.DB
	validator := validator.New()
	quisResult := new(models.UserQuizResult)
	// bind
	if err := ctx.Bind(&quisResult); err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Quiz result add Failed",
			Error:   err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	// validate
	if err := validator.Struct(quisResult); err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Quiz result add Failed",
			Error:   err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	if err := db.Create(&quisResult).Error; err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Quiz result add Failed",
			Error:   err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response := response.SuccessResponse{
		Status:  "success",
		Message: "Quiz result added successfully",
		Data:    quisResult,
	}
	return ctx.JSON(http.StatusOK, response)
}
