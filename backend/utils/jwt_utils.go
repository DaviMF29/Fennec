package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/DaviMF29/fennec/models"
	"github.com/golang-jwt/jwt"
)

var SECRET_KEY []byte

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

func ExtractUserIdFromRequest(r *http.Request) (string, error) {
	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" {
		return "", errors.New("token not provided")
	}

	var tokenString string
	if strings.HasPrefix(tokenHeader, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenHeader, "Bearer ")
	} else {
		return "", errors.New("invalid token format")
	}

	userId, err := GetUserIDFromToken(tokenString)
	if err != nil {
		return "", errors.New("error getting user ID from token")
	}

	return userId, nil
}