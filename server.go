package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func serverEntry(cfg *Config) {
	// Define CORS options
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"}, // You can customize this based on your needs
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"}, // You can customize this based on your needs
		AllowCredentials: true,
		MaxAge:           300, // Maximum age for cache, in seconds
	}

	mainRouter := chi.NewRouter()
	v1Router := chi.NewRouter()
	mainRouter.Use(cors.Handler(corsOptions))
	v1Router.Get("/err", readinessErr)
	v1Router.Get("/readiness", readinessSuccess)
	v1Router.Post("/users", cfg.postUsersHandler)
	v1Router.Get("/users", cfg.middlewareAuth(cfg.getUsersHandler))
	v1Router.Post("/feeds", cfg.middlewareAuth(cfg.postFeedsHandler))
	v1Router.Post("/feed_follows", cfg.middlewareAuth(cfg.postFeedFollowsHandler))
	v1Router.Get("/feed_follows/{feedFollowID}", cfg.middlewareAuth(cfg.getFeedFollowsHandler))

	v1Router.Delete("/feed_follows", cfg.middlewareAuth(cfg.deleteFeedFollowsHandler))
	v1Router.Get("/feeds", cfg.getFeedsHandler)
	v1Router.Get("/posts", cfg.middlewareAuth(cfg.getPostsHandler))
	mainRouter.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr:    ":" + cfg.PORT,
		Handler: mainRouter,
	}
	log.Printf("Serving on port: %s\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())

}
