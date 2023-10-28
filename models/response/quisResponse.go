package response

import (
	"time"
)

type QuizResponse struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	QuizName    string `json:"quiz_name"`
	Description string `json:"description"`
	Duration    int64  `json:"duration"`
}
