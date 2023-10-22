package controller

import (
	"main/libs/database"
	"main/models/response"
	"net/http"

	"github.com/labstack/echo"
)

func GetQuizzesController(c echo.Context) error {
	quizzes, err := database.GetQuizzes()
	if err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Get Quizzes Failed",
			Error:   err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	response := response.SuccessResponse{
		Status:  "success",
		Message: "Get Quizzes successfully",
		Data:    quizzes,
	}
	return c.JSON(http.StatusOK, response)
}

func CreateQuizController(ctx echo.Context) error {
	quiz, err := database.CreateQuiz(ctx)
	if err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Quiz Creation Failed",
			Error:   err.Error(),
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	response := response.SuccessResponse{
		Status:  "success",
		Message: "Quiz Created successfully",
		Data:    quiz,
	}

	return ctx.JSON(http.StatusOK, response)

}

func GetQuizController(ctx echo.Context) error {
	quiz, err := database.GetQuiz(ctx)
	if err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Get Quiz Failed",
			Error:   err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}
	response := response.SuccessResponse{
		Status:  "success",
		Message: "Get Quiz successfully",
		Data:    quiz,
	}
	return ctx.JSON(http.StatusOK, response)
}

func UpdateQuizController(ctx echo.Context) error {
	quiz, err := database.UpdateQuiz(ctx)
	if err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Quiz Update Failed",
			Error:   err.Error(),
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	response := response.SuccessResponse{
		Status:  "success",
		Message: "Quiz Updated successfully",
		Data:    quiz,
	}

	return ctx.JSON(http.StatusOK, response)
}

func DeleteQuizController(ctx echo.Context) error {
	_, err := database.DeleteQuiz(ctx)
	if err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Quiz Delete Failed",
			Error:   err.Error(),
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	response := response.SuccessResponse{
		Status:  "success",
		Message: "Quiz Deleted successfully",
		Data:    nil,
	}

	return ctx.JSON(http.StatusOK, response)
}
