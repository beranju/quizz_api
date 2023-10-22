package models

import (
	"github.com/jinzhu/gorm"
)

type UserAnswer struct {
	gorm.Model
	UserId          int64  `json:"user_id" form:"user_id"`
	QuizId          int64  `json:"quiz_id" form:"quiz_id"`
	QuestionId      int64  `json:"question_id" form:"question_id"`
	UserAnswerIndex uint16 `json:"user_answer_index" form:"user_answer_index"`
}
