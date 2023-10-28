package controller

import (
	"main/models"
	"main/models/response"
	"main/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetQuizzesController(c echo.Context) error {
	var quizzes []models.Quiz

	if err := repositories.FindQuiz(&quizzes); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Failed to retrieve quizzes",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := response.BaseResponse{
		Status:  "success",
		Message: "Quizzes retrieved successfully",
		Data:    quizzes,
	}
	return c.JSON(http.StatusOK, response)
}

func CreateQuizController(ctx echo.Context) error {
	quiz := new(models.Quiz)

	// bind data
	if err := ctx.Bind(&quiz); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Check your data",
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	if err := repositories.CreateQuiz(quiz); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Quiz Creation Failed",
		}
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response := response.BaseResponse{
		Status:  "success",
		Message: "Quiz Created successfully",
		Data:    quiz,
	}

	return ctx.JSON(http.StatusOK, response)

}

func GetQuizController(ctx echo.Context) error {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	exisitingQuiz := new(models.Quiz)

	if err := repositories.FindQuizById(exisitingQuiz, id); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Quiz not found",
		}
		return ctx.JSON(http.StatusNotFound, response)
	}

	response := response.BaseResponse{
		Status:  "success",
		Message: "Quiz retrieved successfully",
		Data:    exisitingQuiz,
	}
	return ctx.JSON(http.StatusOK, response)
}

func UpdateQuizController(ctx echo.Context) error {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	quiz := new(models.Quiz)

	// bind data
	if err := ctx.Bind(quiz); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Check you data",
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	exisitingQuiz := new(models.Quiz)

	if err := repositories.FindQuizById(exisitingQuiz, id); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Quiz not found",
		}
		return echo.NewHTTPError(http.StatusNotFound, response)
	}

	exisitingQuiz.QuizName = quiz.QuizName
	exisitingQuiz.Description = quiz.Description
	exisitingQuiz.Duration = quiz.Duration

	if err := repositories.SaveQuiz(quiz); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Quiz Update Failed",
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}

	response := response.BaseResponse{
		Status:  "success",
		Message: "Quiz Updated successfully",
		Data:    exisitingQuiz,
	}

	return ctx.JSON(http.StatusOK, response)
}

func DeleteQuizController(ctx echo.Context) error {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	quiz := new(models.Quiz)

	if err := repositories.DeleteQuiz(quiz, id); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Quiz Delete Failed",
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}
	response := response.BaseResponse{
		Status:  "success",
		Message: "Quiz Deleted successfully",
		Data:    nil,
	}

	return ctx.JSON(http.StatusOK, response)
}
