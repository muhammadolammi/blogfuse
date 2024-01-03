package main

import (
	"fmt"
	"net/http"

	"github.com/muhammadolammi/blogfuse/internal/database"
)

func (cfg *Config) getPostsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := cfg.DB.GetUserPosts(r.Context(), database.GetUserPostsParams{
		UserID: user.ID,
		Limit:  20,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error gettings posts %v", err))
	}
	respondWithJson(w, 200, databasePostsToPosts(posts))
}
