package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/DaviMF29/wombat/models"
	"github.com/DaviMF29/wombat/repository"
	"github.com/go-chi/chi"
	"github.com/DaviMF29/wombat/utils"
)

func InsertUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Erro ao decodificar JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if user.Name == "" || user.Username == "" ||user.Email == "" || user.Password == "" || user.BirthDate == ""{
		http.Error(w, "Missing fields in request", http.StatusBadRequest)
        return
	}

	userDataByEmail, err := repository.GetUserByEmail(user.Email)

    if err == nil && userDataByEmail.Email != "" {
		utils.SendErrorResponse(w, "Email already exists")
        return
    }

    userDataByUsername, err := repository.GetUserByUsername(user.Username)

    if err == nil && userDataByUsername.Username != "" {
		utils.SendErrorResponse(w, "Username already exists")
        return
    }

	id, err := repository.InsertUser(user)

	if err != nil {
		utils.SendErrorResponse(w, "Error inserting user")
		return
	}

	utils.SendSuccessResponse(w, fmt.Sprintf("User inserted with ID: %s", id))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.SendErrorResponse(w, "Missing ID in request")
		return
	}


	user, err := repository.GetUserById(id)

	if err != nil {
		utils.SendErrorResponse(w, "Error getting user")
		return
	}

	utils.SendSuccessResponse(w, fmt.Sprintf("User: %v", user))
}
