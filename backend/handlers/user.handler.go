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

//	@Summary		Create a new user
//	@Description	Creates a new user with the provided details if the email and username are unique.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.User			true	"User data"
//	@Success		201		{string}	string				"User inserted successfully with ID"
//	@Failure		400		{object}	map[string]string	"Bad request - Missing fields or invalid JSON"
//	@Failure		409		{object}	map[string]string	"Conflict - Email or username already exists"
//	@Failure		500		{object}	map[string]string	"Internal server error - Unable to insert user"
//	@Router			/api/user [post]
func InsertUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Erro ao decodificar JSON: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	if user.Name == "" || user.Username == "" || user.Email == "" || user.Password == "" || user.BirthDate == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Missing fields in request")
		return
	}

	userDataByEmail, err := repository.GetUserByEmail(user.Email)
	if err == nil && userDataByEmail.Email != "" {
		utils.SendErrorResponse(w, http.StatusConflict, "Email already exists")
		return
	}

	userDataByUsername, err := repository.GetUserByUsername(user.Username)
	if err == nil && userDataByUsername.Username != "" {
		utils.SendErrorResponse(w, http.StatusConflict, "Username already exists")
		return
	}

	_, err = repository.InsertUser(user)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Error inserting user")
		return
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Error generating token")
		return
	}

	utils.SendSuccessResponse(w, http.StatusCreated, token)
}

//	@Summary		Get a user by ID
//	@Description	Retrieves a user based on the provided user ID.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string				true	"User ID"
//	@Success		200	{object}	models.User			"User details"
//	@Failure		400	{object}	map[string]string	"Bad request - Missing ID"
//	@Failure		404	{object}	map[string]string	"Not found - User not found"
//	@Failure		500	{object}	map[string]string	"Internal server error - Unable to retrieve user"
//	@Router			/api/user/{id} [get]
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Missing ID in request")
		return
	}

	user, err := repository.GetUserById(id)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusNotFound, "Error getting user")
		return
	}

	utils.SendSuccessResponse(w, http.StatusOK, fmt.Sprintf("User: %v", user))
}


//	@Summary		Delete a user by ID
//	@Description	Deletes a user based on the provided user ID.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string				true	"User ID"
//	@Success		200	{string}	string				"User deleted successfully"
//	@Failure		400	{object}	map[string]string	"Bad request - Missing ID"
//	@Failure		404	{object}	map[string]string	"Not found - User not found"
//	@Failure		500	{object}	map[string]string	"Internal server error - Unable to delete user"
//	@Router			/api/user/{id} [delete]
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Missing ID in request")
		return
	}

	userID, err := utils.ExtractUserIdFromRequest(r)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusUnauthorized, "Error getting user ID from token")
		return
	}

	if id != userID {
		utils.SendErrorResponse(w, http.StatusForbidden, "Unauthorized")
		return
	}
	
	err = repository.DeleteUserById(id)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Error deleting user")
		return
	}

	utils.SendSuccessResponse(w, http.StatusOK, "User deleted successfully")
}

//@Summary Update a user by ID
//	@Description	Updates a user based on the provided user ID.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string				true	"User ID"
//	@Param			user	body		models.User			true	"User data"
//	@Success		200		{string}	string				"User updated successfully"
//	@Failure		400		{object}	map[string]string	"Bad request - Missing ID or invalid JSON"
//	@Failure		403		{object}	map[string]string	"Forbidden - Unauthorized to update these fields"
//	@Failure		404		{object}	map[string]string	"Not found - User not found"
//	@Failure		409		{object}	map[string]string	"Conflict - Email or username already exists"
//	@Failure		500		{object}	map[string]string	"Internal server error - Unable to update user"
//	@Router			/api/user/{id} [put]
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Missing ID in request")
		return
	}

	var newDataUser models.User
	err := json.NewDecoder(r.Body).Decode(&newDataUser)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	if newDataUser.Name == "" && newDataUser.Username == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Missing fields in request")
		return
	}

	userID, err := utils.ExtractUserIdFromRequest(r)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusUnauthorized, "Error getting user ID from token")
		return
	}
	if id != userID {
		utils.SendErrorResponse(w, http.StatusForbidden, "Unauthorized")
		return
	}

	existingUser, err := repository.GetUserById(id)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}

	if newDataUser.Email != "" || newDataUser.Password != "" || newDataUser.BirthDate != "" {
		utils.SendErrorResponse(w, http.StatusForbidden, "Unauthorized to update these fields")
		return
	}

	if newDataUser.Name != "" {
		existingUser.Name = newDataUser.Name
	}
	if newDataUser.Username != "" {
		userDataByUsername, err := repository.GetUserByUsername(newDataUser.Username)
		if err == nil && userDataByUsername.Username != "" && userDataByUsername.ID.Hex() != id {
			utils.SendErrorResponse(w, http.StatusConflict, "Username already exists")
			return
		}
		existingUser.Username = newDataUser.Username
	}

	err = repository.UpdateUserById(userID, existingUser)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Error updating user")
		return
	}

	utils.SendSuccessResponse(w, http.StatusOK, "User updated successfully")
}

