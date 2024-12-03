package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/DaviMF29/wombat/models"
	"github.com/go-chi/chi"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		log.Println("ID não fornecido")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user, err := models.GetUserById(id)
	if err != nil {
		log.Printf("Erro ao buscar usuário: %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("Erro ao codificar resposta: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func InsertUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Erro ao decodificar JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	userEmail, err := models.GetUserByEmail(user.Email)
	if err == nil && userEmail.Email != "" {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	id, err := models.InsertUser(user)
	if err != nil {
		log.Printf("Erro ao inserir usuário: %v", err)
		resp := map[string]interface{}{
			"error":   true,
			"message": fmt.Sprintf("Erro ao inserir usuário: %v", err),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := map[string]interface{}{
		"error":   false,
		"message": fmt.Sprintf("Usuário inserido com sucesso. ID: %s", id),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
