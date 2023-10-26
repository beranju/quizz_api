package controller

import (
	"main/config"
	"main/models"
	"main/models/response"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func GetAllQuestionController(ctx echo.Context) error {
	db := config.DB
	quizId := ctx.Param("quizId")
	var questions []models.Question

	if err := db.Model(&models.Question{}).Where("quiz_id = ?", quizId).Find(&questions).Error; err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve questions",
			Error:   err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response := response.SuccessResponse{
		Status:  "success",
		Message: "Questions retrieved successfully",
		Data:    questions,
	}
	return ctx.JSON(http.StatusOK, response)

}

func CreateQuestionController(ctx echo.Context) error {
	db := config.DB
	validator := validator.New()
	questions := new(models.Question)

	// bind
	if err := ctx.Bind(&questions); err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Questions add Failed",
			Error:   err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	// validate
	if err := validator.Struct(questions); err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Questions add Failed",
			Error:   err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	if err := db.Create(&questions).Error; err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Question add Failed",
			Error:   err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response := response.SuccessResponse{
		Status:  "success",
		Message: "Questions added successfully",
		Data:    questions,
	}
	return ctx.JSON(http.StatusOK, response)
}

func GetQuestionController(ctx echo.Context) error {
	db := config.DB
	quizId := ctx.Param("quizId")
	questionId := ctx.Param("questionId")
	existingQuestion := new(models.Question)

	if err := db.Where("ID = ? AND quiz_id = ?", questionId, quizId).First(&existingQuestion, &quizId, &questionId).Error; err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve Question",
			Error:   err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, response)

	}
	response := response.SuccessResponse{
		Status:  "success",
		Message: "Question retrieved successfully",
		Data:    existingQuestion,
	}
	return ctx.JSON(http.StatusOK, response)

}

func UpdateQuestionController(ctx echo.Context) error {
	db := config.DB
	validator := validator.New()
	quizId := ctx.Param("quizId")
	questionId := ctx.Param("questionId")
	question := new(models.Question)

	// bind data
	if err := ctx.Bind(question); err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Question Update Failed",
			Error:   err.Error(),
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	// validate the quiz sturct
	if err := validator.Struct(question); err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Question Update Failed",
			Error:   err.Error(),
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	existingQuestion := new(models.Question)

	if err := db.Where("ID = ? AND quiz_id = ?", questionId, quizId).First(&existingQuestion, questionId, quizId).Error; err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Question Update Failed",
			Error:   err.Error(),
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}

	existingQuestion.QuizId = question.QuizId
	existingQuestion.Text = question.Text
	existingQuestion.Score = question.Score
	existingQuestion.Answer = question.Answer
	existingQuestion.Options1 = question.Options1
	existingQuestion.Options2 = question.Options2
	existingQuestion.Options3 = question.Options3
	existingQuestion.Options4 = question.Options4
	existingQuestion.ImageUrl = question.ImageUrl

	if err := config.DB.Save(&existingQuestion).Error; err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Question Update Failed",
			Error:   err.Error(),
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}

	response := response.SuccessResponse{
		Status:  "success",
		Message: "Question Updated successfully",
		Data:    existingQuestion,
	}

	return ctx.JSON(http.StatusOK, response)
}

func DeleteQuestionController(ctx echo.Context) error {
	db := config.DB
	questionId := ctx.Param("questionId")
	quizId := ctx.Param("quizId")
	question := new(models.Question)

	if err := db.Where("ID = ? AND quiz_id = ?", questionId, quizId).Delete(&question, questionId, quizId).Error; err != nil {
		response := response.ErrorResponse{
			Status:  "error",
			Message: "Question Delete Failed",
			Error:   err.Error(),
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}
	response := response.SuccessResponse{
		Status:  "success",
		Message: "Question Deleted successfully",
		Data:    nil,
	}

	return ctx.JSON(http.StatusOK, response)

}
