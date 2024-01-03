package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/muhammadolammi/blogfuse/internal/database"
)

func (cfg *Config) postFeedsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type Parameter struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	params := Parameter{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error decoding request body : %v", err.Error()))
		return
	}
	feedParam := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	}
	feed, err := cfg.DB.CreateFeed(r.Context(), feedParam)
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}
	followFeedParam := database.FollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}
	feedFollow, err := cfg.DB.FollowFeed(r.Context(), followFeedParam)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error following feed %v", err))
		return
	}
	respondBody := struct {
		Feed       Feed       `json:"feed"`
		FeedFollow FeedFollow `json:"feed_follow"`
	}{
		Feed:       databaseFeedToFeed(feed),
		FeedFollow: databaseFeedFollowToFeedFollow(feedFollow),
	}

	respondWithJson(w, 200, respondBody)
}

func (cfg *Config) getFeedsHandler(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}
	respondWithJson(w, 200, databaseFeedsToFeeds(feeds))

}
