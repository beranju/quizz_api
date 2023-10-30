package repositories

import (
	"main/config"
	"main/models"
)

func FindAllQuestion(questions *[]models.Question, quizId int) error {
	result := config.DB.Model(&models.Question{}).Where("quiz_id = ?", quizId).Find(&questions)
	return result.Error
}

func FindQuestionById(question *models.Question, id, quiz_id int) error {
	result := config.DB.Where("quiz_id = ?", quiz_id).First(question, id)
	return result.Error
}

func CreateQuestion(question *models.Question) error {
	result := config.DB.Create(&question)
	return result.Error
}

func SaveQuestion(question *models.Question) error {
	result := config.DB.Save(&question)
	return result.Error
}

func DeleteQuestion(question *models.Question, id, quiz_id int) error {
	result := config.DB.Where("quiz_id = ?", quiz_id).Delete(&question, id)
	return result.Error
}
