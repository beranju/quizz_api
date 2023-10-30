package controller

import (
	"errors"
	"main/models"
	"main/models/response"
	"main/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAllQuestionController(ctx echo.Context) error {
	quizIdString := ctx.Param("quizId")
	quizId, _ := strconv.Atoi(quizIdString)
	var questions []models.Question

	if err := repositories.FindAllQuestion(&questions, quizId); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Failed to retrieve questions",
		}
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	var questionResponse []response.QuestionResponse
	for _, q := range questions {
		question := q.ToQuestionResponse()
		questionResponse = append(questionResponse, question)
	}

	response := response.BaseResponse{
		Status:  "success",
		Message: "Questions retrieved successfully",
		Data:    questionResponse,
	}
	return ctx.JSON(http.StatusOK, response)

}

func CreateQuestionController(ctx echo.Context) error {
	validator := validator.New()
	questions := new(models.Question)

	// bind
	if err := ctx.Bind(questions); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Request body is invalid",
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	if err := validator.Struct(questions); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Check your data",
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	if err := repositories.CreateQuestion(questions); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Question add Failed",
		}
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	// TODO: create method fo these requests
	questionResponse := questions.ToQuestionResponse()

	response := response.BaseResponse{
		Status:  "success",
		Message: "Questions added successfully",
		Data:    questionResponse,
	}
	return ctx.JSON(http.StatusOK, response)
}

func GetQuestionController(ctx echo.Context) error {
	quizIdString := ctx.Param("quizId")
	quizId, _ := strconv.Atoi(quizIdString)
	questionIdString := ctx.Param("questionId")
	questionId, _ := strconv.Atoi(questionIdString)
	existingQuestion := new(models.Question)

	if err := repositories.FindQuestionById(existingQuestion, questionId, quizId); err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			response := response.BaseResponse{
				Status:  "error",
				Message: "Question not found",
			}
			return ctx.JSON(http.StatusNotFound, response)
		}

		response := response.BaseResponse{
			Status:  "error",
			Message: "Failed to retrieve Question",
		}
		return ctx.JSON(http.StatusInternalServerError, response)

	}

	questionResponse := existingQuestion.ToQuestionResponse()

	response := response.BaseResponse{
		Status:  "success",
		Message: "Question retrieved successfully",
		Data:    questionResponse,
	}
	return ctx.JSON(http.StatusOK, response)

}

func UpdateQuestionController(ctx echo.Context) error {
	validator := validator.New()
	quizIdString := ctx.Param("quizId")
	quizId, _ := strconv.Atoi(quizIdString)
	questionIdString := ctx.Param("questionId")
	questionId, _ := strconv.Atoi(questionIdString)
	question := new(models.Question)

	// bind data
	if err := ctx.Bind(&question); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Check your data",
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	if err := validator.Struct(question); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Check your data",
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	existingQuestion := new(models.Question)

	if err := repositories.FindQuestionById(existingQuestion, questionId, quizId); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Question not found",
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

	if err := repositories.SaveQuestion(existingQuestion); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Question Update Failed",
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}

	response := response.BaseResponse{
		Status:  "success",
		Message: "Question Updated successfully",
		Data:    existingQuestion,
	}

	return ctx.JSON(http.StatusOK, response)
}

func DeleteQuestionController(ctx echo.Context) error {
	questionIdString := ctx.Param("questionId")
	questionId, _ := strconv.Atoi(questionIdString)
	quizIdString := ctx.Param("quizId")
	quizId, _ := strconv.Atoi(quizIdString)
	question := new(models.Question)

	if err := repositories.DeleteQuestion(question, questionId, quizId); err != nil {
		response := response.BaseResponse{
			Status:  "error",
			Message: "Question Delete Failed",
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}
	response := response.BaseResponse{
		Status:  "success",
		Message: "Question Deleted successfully",
	}

	return ctx.JSON(http.StatusOK, response)

}
