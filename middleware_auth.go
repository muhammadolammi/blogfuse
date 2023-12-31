package main

import (
	"net/http"

	"github.com/muhammadolammi/blogfuse/internal/auth"
	"github.com/muhammadolammi/blogfuse/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *Config) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 500, err.Error())
			return
		}
		dbUser, err := cfg.DB.GetUserByAPIKey(r.Context(), apikey)
		if err != nil {
			respondWithError(w, 500, err.Error())
			return
		}
		handler(w, r, dbUser)
	}
}
