package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/DaviMF29/fennec/models"
	"github.com/DaviMF29/fennec/repository"
	"github.com/DaviMF29/fennec/utils"
	"golang.org/x/crypto/bcrypt"
)

//	@Summary		User login
//	@Description	Authenticates a user using their email and password, returning a JWT token upon success.
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			loginData	body		models.LoginData	true	"Login data"
//	@Success		200			{object}	map[string]string	"JWT token"
//	@Failure		400			{object}	map[string]string	"Bad request"
//	@Failure		401			{object}	map[string]string	"Unauthorized"
//	@Router			/api/login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginData models.LoginData

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