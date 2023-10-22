package models

import (
	"github.com/jinzhu/gorm"
)

type Quiz struct {
	gorm.Model
	QuizName    string `json:"quiz_name" form:"quiz_name"`
	Description string `json:"description" form:"description"`
	Duration    int64  `json:"duration" form:"duration"`
}
