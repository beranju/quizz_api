package response

import (
	"time"
)

type UserResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	ImageUrl  string `json:"image_url"`
	Token     string `json:"token"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
