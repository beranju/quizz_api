package models

import (
	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	QuizId        int64    `jss:"quiz_id" form:"quiz_id"`
	Text          string   `json:"text" form:"text"`
	Score         int64    `json:"score" form:"score"`
	Answer        int32    `json:"answer" form:"answer"`
	AnswerOptions []string `json:"answer_options" form:"answer_options"`
	ImageUrl      string   `json:"image_url" form:"image_url"`
}
