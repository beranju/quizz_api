package models

import (
	"main/models/response"

	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	QuizId   int64  `json:"quiz_id" gorm:"quiz_id;not null"`
	Text     string `json:"text" gorm:"text;not null"`
	Score    int64  `json:"score" gorm:"score;not null"`
	Answer   int32  `json:"answer" gorm:"answer;not null"`
	Options1 string `json:"options_1" gorm:"options_1;not null"`
	Options2 string `json:"options_2" gorm:"options_2;not null"`
	Options3 string `json:"options_3" gorm:"options_3;not null"`
	Options4 string `json:"options_4" gorm:"options_4;not null"`
	ImageUrl string `json:"image_url" gorm:"image_url"`
}

func (question Question) ToQuestionResponse() response.QuestionResponse {
	questionResponse := response.QuestionResponse{
		ID:        question.ID,
		CreatedAt: question.CreatedAt,
		UpdatedAt: question.UpdatedAt,
		QuizId:    question.QuizId,
		Text:      question.Text,
		Score:     question.Score,
		Answer:    question.Answer,
		Options1:  question.Options1,
		Options2:  question.Options2,
		Options3:  question.Options3,
		Options4:  question.Options4,
		ImageUrl:  question.ImageUrl,
	}
	return questionResponse
}
