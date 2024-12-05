package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/DaviMF29/fennec/models"
	"github.com/DaviMF29/fennec/repository"
	"github.com/go-chi/chi"
	"github.com/DaviMF29/fennec/utils"
)

// @Summary Create a new user
// @Description Creates a new user with the provided details if the email and username are unique.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {string} string "User inserted successfully with ID"
// @Failure 400 {object} map[string]string "Bad request - Missing fields or invalid JSON"
// @Failure 409 {object} map[string]string "Conflict - Email or username already exists"
// @Failure 500 {object} map[string]string "Internal server error - Unable to insert user"
// @Router /api/user [post]
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

	_, err = repository.InsertUser(user)

	if err != nil {
		utils.SendErrorResponse(w, "Error inserting user")
		return
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		utils.SendErrorResponse(w, "Error generating token")
		return
	}

	utils.SendSuccessResponse(w, token)

}

// @Summary Get a user by ID
// @Description Retrieves a user based on the provided user ID.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User "User details"
// @Failure 400 {object} map[string]string "Bad request - Missing ID"
// @Failure 404 {object} map[string]string "Not found - User not found"
// @Failure 500 {object} map[string]string "Internal server error - Unable to retrieve user"
// @Router /api/user/{id} [get]
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

// @Summary Delete a user by ID
// @Description Deletes a user based on the provided user ID.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} string "User deleted successfully"
// @Failure 400 {object} map[string]string "Bad request - Missing ID"
// @Failure 404 {object} map[string]string "Not found - User not found"
// @Failure 500 {object} map[string]string "Internal server error - Unable to delete user"
// @Router /api/user/{id} [delete]
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.SendErrorResponse(w, "Missing ID in request")
		return
	}

	userID, err := utils.ExtractUserIdFromRequest(r)
	if err != nil {
		utils.SendErrorResponse(w, "Error getting user ID from token")
		return
	}

	if id != userID {
		utils.SendErrorResponse(w, "Unauthorized")
		return
	}
	
	err = repository.DeleteUserById(id)
	if err != nil {
		utils.SendErrorResponse(w, "Error deleting user")
		return
	}

	utils.SendSuccessResponse(w, "User deleted successfully")
}

