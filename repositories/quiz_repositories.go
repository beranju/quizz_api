package repositories

import (
	"main/config"
	"main/models"
)

func FindQuiz(quizzes *[]models.Quiz) error {
	result := config.DB.Find(&quizzes)
	return result.Error
}

func CreateQuiz(quiz *models.Quiz) error {
	result := config.DB.Create(&quiz)
	return result.Error
}

func FindQuizById(quiz *models.Quiz, id int) error {
	result := config.DB.First(quiz, id)
	return result.Error
}

func SaveQuiz(quiz *models.Quiz) error {
	result := config.DB.Save(&quiz)
	return result.Error
}

func DeleteQuiz(quiz *models.Quiz, id int) error {
	result := config.DB.Delete(&quiz, id)
	return result.Error
}
