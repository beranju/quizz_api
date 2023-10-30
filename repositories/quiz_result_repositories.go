package repositories

import (
	"main/config"
	"main/models"
)

func FindResult(quizResult *models.UserQuizResult, quizId, userId int) error {
	result := config.DB.Where("quiz_id = ? AND user_id = ?", quizId, userId).First(&quizResult)
	return result.Error
}

func CreateResult(quizResult *models.UserQuizResult) error {
	result := config.DB.Create(&quizResult)
	return result.Error
}
