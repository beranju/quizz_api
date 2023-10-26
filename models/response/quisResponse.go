package response

import (
	"main/models"
)

type QuizResponse struct {
	QuizName    string `json:"quiz_name" form:"quiz_name"`
	Description string `json:"description" form:"description"`
	Duration    int64  `json:"duration" form:"duration"`
	Questions   []models.Question
}
