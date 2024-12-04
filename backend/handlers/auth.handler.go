package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/DaviMF29/wombat/repository"
	"github.com/DaviMF29/wombat/utils"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Email    string `bson:"email" json:"email"`
		Password string `bson:"password" json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if loginData.Email == "" || loginData.Password == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	token, err := authenticate(loginData.Email, loginData.Password)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func authenticate(email string, password string) (string, error) {

	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if !verifyPassword(user.Password, password) {
		return "", errors.New("password incorrect")
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func verifyPassword(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}