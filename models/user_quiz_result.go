package models

import (
	"main/models/response"

	"github.com/jinzhu/gorm"
)

type UserQuizResult struct {
	gorm.Model
	UserId    int64 `json:"user_id" gorm:"user_id;not null"`
	QuizId    int64 `json:"quiz_id" gorm:"quiz_id;not null"`
	UserScore int32 `json:"user_score" gorm:"user_score;not null"`
}

func (result UserQuizResult) ToQuizResultResponse() response.QuizResultResponse {
	quizResultResponse := response.QuizResultResponse{
		ID:        result.ID,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		UserId:    result.UserId,
		QuizId:    result.QuizId,
		UserScore: result.UserScore,
	}
	return quizResultResponse
}
