package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/muhammadolammi/blogfuse/internal/database"
)

func (cfg *Config) postUsersHandler(w http.ResponseWriter, r *http.Request) {
	type Parameter struct {
		Name string `json:"name"`
	}
	params := Parameter{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error decoding request body : %v", err.Error()))
		return
	}
	userParam := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	}
	dbUser, err := cfg.DB.CreateUser(r.Context(), userParam)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error creating user. err : %v", err.Error()))
		return
	}
	user := databaseUserToUser(dbUser)
	respondWithJson(w, 200, user)
}

func (cfg *Config) getUsersHandler(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJson(w, 200, databaseUserToUser(user))
}
