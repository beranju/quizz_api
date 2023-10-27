package database

import (
	"main/config"
	"main/models"

	"github.com/labstack/echo/v4"
)

func GetQuizzes() (interface{}, error) {
	var quizzes []models.Quiz
	if err := config.DB.Find(&quizzes).Error; err != nil {
		return nil, err
	}
	return quizzes, nil
}

func GetQuiz(ctx echo.Context) (interface{}, error) {
	id := ctx.Param("id")

	var quizzes []*models.Quiz
	if err := config.DB.Find(&quizzes, id).Error; err != nil {
		return nil, err
	}
	quiz := quizzes[0]
	return quiz, nil
}

func CreateQuiz(ctx echo.Context) (interface{}, error) {
	quiz := new(models.Quiz)

	if err := ctx.Bind(quiz); err != nil {
		return nil, err
	}

	if err := config.DB.Create(&quiz).Error; err != nil {
		return nil, err
	}

	quizCreated := &models.Quiz{
		QuizName:    quiz.QuizName,
		Description: quiz.Description,
		Duration:    quiz.Duration,
	}
	return quizCreated, nil
}

func UpdateQuiz(ctx echo.Context) (interface{}, error) {
	id := ctx.Param("id")
	dataQuiz := new(models.Quiz)

	if err := ctx.Bind(dataQuiz); err != nil {
		return nil, err
	}

	existingQuiz := new(models.Quiz)
	if err := config.DB.First(&existingQuiz, id).Error; err != nil {
		return nil, err
	}

	existingQuiz.QuizName = dataQuiz.QuizName
	existingQuiz.Description = dataQuiz.Description
	existingQuiz.Duration = dataQuiz.Duration

	if err := config.DB.Save(&existingQuiz).Error; err != nil {
		return nil, err
	}
	return existingQuiz, nil
}

func DeleteQuiz(ctx echo.Context) (interface{}, error) {
	id := ctx.Param("id")
	existingQuiz := new(models.Quiz)

	if err := config.DB.Delete(&existingQuiz, id).Error; err != nil {
		return nil, err
	}
	return existingQuiz, nil
}
