package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"github.com/DaviMF29/wombat/models"
	"github.com/DaviMF29/wombat/repository"
)
func InsertPostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		utils.SendErrorResponse(w, "Error decoding JSON")
	}
	
	if post.Content == "" {
        utils.SendErrorResponse(w, "Missing fields in request")
		return
    }

	tokenHeader := r.Header.Get("Authorization")
	var tokenString string

	if strings.HasPrefix(tokenHeader, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenHeader, "Bearer ")
	} else {
		utils.SendErrorResponse(w, "Token not provided")
		return
	}

	userId, err := utils.GetUserIDFromToken(tokenString)
	if err != nil {
		utils.SendErrorResponse(w, "Error getting user ID from token")
		return
	}

	post.UserID = userId

	id, err := repository.InsertPost(post)

	if err != nil {
		utils.SendErrorResponse(w, "Error inserting post")
		return
	}

	utils.SendSuccessResponse(w, fmt.Sprintf("Post inserted with ID: %s", id))
}

func GetPostByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.SendErrorResponse(w, "Missing ID in request")
		return
	}


	post, err := repository.GetPostById(id)

	if err != nil {
		utils.SendErrorResponse(w, "Error getting post")
		return
	}

	utils.SendSuccessResponse(w, fmt.Sprintf("Post: %v", post))
}