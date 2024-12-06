package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/DaviMF29/fennec/models"
	"github.com/DaviMF29/fennec/repository"
	"github.com/DaviMF29/fennec/utils"
	"github.com/go-chi/chi"
)

//	@Summary		Create a new post
//	@Description	Creates a new post associated with the authenticated user.
//	@Tags			Posts
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Bearer token"
//	@Param			post			body		models.Post			true	"Post data"
//	@Success		201				{string}	string				"Post inserted with ID"
//	@Failure		400				{object}	map[string]string	"Bad request - Missing fields or invalid JSON"
//	@Failure		401				{object}	map[string]string	"Unauthorized - Token not provided or invalid"
//	@Failure		500				{object}	map[string]string	"Internal server error - Unable to insert post"
//	@Router			/api/post [post]
func InsertPostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Error decoding JSON")
		return
	}

	if post.Content == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Missing fields in request")
		return
	}

	if len(post.Content) > 300 {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Content too long")
		return
	}

	tokenHeader := r.Header.Get("Authorization")
	var tokenString string

	if strings.HasPrefix(tokenHeader, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenHeader, "Bearer ")
	} else {
		utils.SendErrorResponse(w, http.StatusUnauthorized, "Token not provided")
		return
	}

	userId, err := utils.GetUserIDFromToken(tokenString)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusUnauthorized, "Error getting user ID from token")
		return
	}

	post.UserID = userId

	id, err := repository.InsertPost(post)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Error inserting post")
		return
	}

	utils.SendSuccessResponse(w, http.StatusCreated, fmt.Sprintf("Post inserted with ID: %s", id))
}


//	@Summary		Get a post by ID
//	@Description	Retrieves a specific post by its unique ID.
//	@Tags			Posts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string				true	"Post ID"
//	@Success		200	{object}	models.Post			"Post retrieved successfully"
//	@Failure		400	{object}	map[string]string	"Missing ID in request"
//	@Failure		500	{object}	map[string]string	"Error getting post"
//	@Router			/api/post/{id} [get]
func GetPostByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Missing ID in request")
		return
	}

	post, err := repository.GetPostById(id)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusNotFound, "Error getting post")
		return
	}

	utils.SendSuccessResponse(w, http.StatusOK, fmt.Sprintf("Post: %v", post))
}

//	@Summary		Delete a post by ID
//	@Description	Deletes a post based on the provided post ID.
//	@Tags			Posts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string				true	"Post ID"
//	@Success		200	{string}	string				"Post deleted successfully"
//	@Failure		400	{object}	map[string]string	"Bad request - Missing ID"
//	@Failure		404	{object}	map[string]string	"Not found - Post not found"
//	@Failure		500	{object}	map[string]string	"Internal server error - Unable to delete post"
//	@Router			/api/post/{id} [delete]
func DeletePostByIdHandler(w http.ResponseWriter, r *http.Request) {
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

	post, err := repository.GetPostById(id)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusNotFound, "Error getting post")
		return
	}

	if userID != post.UserID {
		utils.SendErrorResponse(w, http.StatusForbidden, "Unauthorized - User does not own post")
		return
	}

	err = repository.DeletePostById(id)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Error deleting post")
		return
	}

	utils.SendSuccessResponse(w, http.StatusOK, "Post deleted successfully")
}

//	@Summary		Update a post by ID
//	@Description	Updates a post based on the provided post ID.
//	@Tags			Posts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string				true	"Post ID"
//	@Param			Authorization	header		string				true	"Bearer token
//	@Param			post			body		models.Post			true	"Post data"
//	@Success		200	{string}	string				"Post updated successfully"
//	@Failure		400	{object}	map[string]string	"Bad request - Missing ID or fields in request"
//	@Failure		403	{object}	map[string]string	"Forbidden - User does not own post"
//	@Failure		404	{object}	map[string]string	"Not found - Post not found"
//	@Failure		500	{object}	map[string]string	"Internal server error - Unable to update post"
//	@Router			/api/post/{id} [put]
func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Missing ID in request")
		return
	}

	var newDataPost models.Post

	err := json.NewDecoder(r.Body).Decode(&newDataPost)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Error decoding JSON")
		return
	}

	if newDataPost.Content == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Missing fields in request")
		return
	}

	userID, err := utils.ExtractUserIdFromRequest(r)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusUnauthorized, "Error getting user ID from token")
		return
	}

	existingPost, err := repository.GetPostById(id)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusNotFound, "Post not found")
		return
	}

	if userID != existingPost.UserID {
		utils.SendErrorResponse(w, http.StatusForbidden, "Unauthorized - User does not own post")
		return
	}

	existingPost.Content = newDataPost.Content

	err = repository.UpdatePost(id, existingPost)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Error updating post")
		return
	}

	utils.SendSuccessResponse(w, http.StatusOK, "Post updated successfully")
}
