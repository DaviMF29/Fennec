package utils

import (
	"errors"
	"fmt"
	"github.com/DaviMF29/wombat/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRET_KEY []byte

func init() {
	SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))
	if len(SECRET_KEY) == 0 {
		panic("SECRET_KEY is not set in environment variables")
	}
}

func GenerateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserIDFromToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return SECRET_KEY, nil
	})
	if err != nil {
		return 0, fmt.Errorf("failed to parse token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, ok := claims["id"].(float64)
		if !ok {
			return 0, errors.New("user_id claim not found or invalid")
		}
		return int(userId), nil
	} else {
		return 0, errors.New("invalid token")
	}
}
