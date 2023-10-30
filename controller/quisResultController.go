package controller

import (
	"errors"
	"main/models"
	"main/models/response"
	"main/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetQuizResultController(ctx echo.Context) error {
	quizIdString := ctx.Param("quiz_id")
	quizId, _ := strconv.Atoi(quizIdString)
	// userIdString := ctx.Param("user_id")
	// userId, _ := strconv.Atoi(userIdString)
	quizResult := new(models.UserQuizResult)

	userId := restricted(ctx)

	if err := repositories.FindResult(quizResult, quizId, userId); err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			response := response.BaseResponse{
				Status:  "error",
				Message: "Result not found",
			}
			return ctx.JSON(http.StatusNotFound, response)
		}

		response := response.BaseResponse{
			Status:  "error",
			Message: "Failed to retrieve Quiz result",
		}
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	quizResultResponse := quizResult.ToQuizResultResponse()

	response := response.BaseResponse{
		Status:  "success",
		Message: "Quiz result retrieved successfully",
		Data:    quizResultResponse,
	}
	return ctx.JSON(http.StatusOK, response)
}

func restricted(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)
	id := int(userId)
	return id
}

func AddQuizResultController(ctx echo.Context) error {
	validator := validator.New()
	quisResult := new(models.UserQuizResult)
	// bind
	if err := ctx.Bind(&quisResult); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Check your data",
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	// validate
	if err := validator.Struct(quisResult); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Data validation failed",
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	if err := repositories.CreateResult(quisResult); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Quiz result add Failed",
		}
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	quizResultResponse := quisResult.ToQuizResultResponse()

	response := response.BaseResponse{
		Status:  "success",
		Message: "Quiz result added successfully",
		Data:    quizResultResponse,
	}
	return ctx.JSON(http.StatusOK, response)
}
