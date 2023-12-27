package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	auth "github.com/muhammadolammi/blogfuse/internal"
	"github.com/muhammadolammi/blogfuse/internal/database"
)

func (apiConfig *Config) postUsersHandler(w http.ResponseWriter, r *http.Request) {
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
	dbUser, err := apiConfig.DB.CreateUser(r.Context(), userParam)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error creating user. err : %v", err.Error()))
		return
	}
	user := databaseUserToUser(dbUser)
	respondWithJson(w, 200, user)
}

func (apiConfig *Config) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	apikey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}
	dbUser, err := apiConfig.DB.GetUserByAPIKey(r.Context(), apikey)
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}
	user := databaseUserToUser(dbUser)
	respondWithJson(w, 200, user)
}
