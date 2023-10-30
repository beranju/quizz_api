package models

import (
	"main/models/response"

	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	QuizId   int64  `json:"quiz_id" form:"quiz_id"`
	Text     string `json:"text" form:"text"`
	Score    int64  `json:"score" form:"score"`
	Answer   int32  `json:"answer" form:"answer"`
	Options1 string `json:"options_1" form:"options_1"`
	Options2 string `json:"options_2" form:"options_2"`
	Options3 string `json:"options_3" form:"options_3"`
	Options4 string `json:"options_4" form:"options_4"`
	ImageUrl string `json:"image_url" form:"image_url"`
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
