package models

import (
	"github.com/jinzhu/gorm"
)

type UserQuizResult struct {
	gorm.Model
	UserId    int64 `json:"user_id" form:"user_id"`
	QuizId    int64 `json:"quiz_id" form:"quiz_id"`
	UserScore int32 `json:"user_score" form:"user_score"`
}
