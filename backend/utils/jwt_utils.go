package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/DaviMF29/fennec/models"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

var SECRET_KEY []byte

func init() {
	env := os.Getenv("ENV")

	if env == "production" {
		log.Println("Running in production mode. Skipping .env loading.")
		return
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))
	if len(SECRET_KEY) == 0 {
		log.Fatal("SECRET_KEY não está definida no arquivo .env")
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
