package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	Id   int    `json:"user_id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int, name string) (string, error) {
	claims := &jwtCustomClaims{
		userId,
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	return t, err
}
