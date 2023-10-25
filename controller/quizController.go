package controller

import (
	"main/config"
	"main/models"
	"main/models/response"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func GetQuizzesController(c echo.Context) error {
	var quizzes []models.Quiz
	if err := config.DB.Find(&quizzes).Error; err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve quizzes",
			Error:   err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := response.SuccessResponse{
		Status:  "success",
		Message: "Quizzes retrieved successfully",
		Data:    quizzes,
	}
	return c.JSON(http.StatusOK, response)
}

func CreateQuizController(ctx echo.Context) error {
	db := config.DB
	validator := validator.New()
	quiz := new(models.Quiz)

	// bind data
	if err := ctx.Bind(&quiz); err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Quiz Creation Failed",
			Error:   err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	// validate the quiz sturct
	if err := validator.Struct(quiz); err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Quiz Creation Failed",
			Error:   err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	if err := db.Create(&quiz).Error; err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Quiz Creation Failed",
			Error:   err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response := response.SuccessResponse{
		Status:  "success",
		Message: "Quiz Created successfully",
		Data:    quiz,
	}

	return ctx.JSON(http.StatusOK, response)

}

func GetQuizController(ctx echo.Context) error {
	db := config.DB
	id := ctx.Param("id")
	exisitingQuiz := new(models.Quiz)
	if err := db.First(&exisitingQuiz, id).Error; err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve quiz",
			Error:   err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response := response.SuccessResponse{
		Status:  "success",
		Message: "Quiz retrieved successfully",
		Data:    exisitingQuiz,
	}
	return ctx.JSON(http.StatusOK, response)
}

func UpdateQuizController(ctx echo.Context) error {
	db := config.DB
	validator := validator.New()
	id := ctx.Param("id")
	quiz := new(models.Quiz)

	// bind data
	if err := ctx.Bind(quiz); err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Quiz Update Failed",
			Error:   err.Error(),
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	// validate the quiz sturct
	if err := validator.Struct(quiz); err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Quiz Update Failed",
			Error:   err.Error(),
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	exisitingQuiz := new(models.Quiz)

	if err := db.First(&exisitingQuiz, id).Error; err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Quiz Update Failed",
			Error:   err.Error(),
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}

	exisitingQuiz.QuizName = quiz.QuizName
	exisitingQuiz.Description = quiz.Description
	exisitingQuiz.Duration = quiz.Duration

	if err := config.DB.Save(&exisitingQuiz).Error; err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Quiz Update Failed",
			Error:   err.Error(),
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}

	response := response.SuccessResponse{
		Status:  "success",
		Message: "Quiz Updated successfully",
		Data:    exisitingQuiz,
	}

	return ctx.JSON(http.StatusOK, response)
}

func DeleteQuizController(ctx echo.Context) error {
	db := config.DB
	id := ctx.Param("id")
	quiz := new(models.Quiz)

	if err := db.Delete(&quiz, id).Error; err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Quiz Delete Failed",
			Error:   err.Error(),
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}
	response := response.SuccessResponse{
		Status:  "success",
		Message: "Quiz Deleted successfully",
		Data:    nil,
	}

	return ctx.JSON(http.StatusOK, response)
}
