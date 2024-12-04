package utils

import (
	"errors"
	"fmt"
	"os"
	"time"
	"github.com/DaviMF29/wombat/models"
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
		"userId":user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 168).Unix(),
	})

	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserIDFromToken(tokenString string) (string, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return []byte(SECRET_KEY), nil
    })
    if err != nil {
        return "", fmt.Errorf("failed to parse token: %w", err)
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return "", errors.New("invalid token claims")
    }

    userId, ok := claims["userId"].(string)
    if !ok {
        return "", errors.New("userId claim not found or invalid")
    }

    return userId, nil
}
