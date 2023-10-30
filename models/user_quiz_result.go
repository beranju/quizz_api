package models

import (
	"main/models/response"

	"github.com/jinzhu/gorm"
)

type UserQuizResult struct {
	gorm.Model
	UserId    int64 `json:"user_id" form:"user_id"`
	QuizId    int64 `json:"quiz_id" form:"quiz_id"`
	UserScore int32 `json:"user_score" form:"user_score"`
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
