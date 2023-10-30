package response

import (
	"time"
)

type QuizResultResponse struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    int64 `json:"user_id"`
	QuizId    int64 `json:"quiz_id"`
	UserScore int32 `json:"user_score"`
}
