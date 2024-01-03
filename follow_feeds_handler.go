package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/muhammadolammi/blogfuse/internal/database"
)

func (cfg *Config) postFeedFollowsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type Parameter struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	params := Parameter{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error decoding request body : %v", err.Error()))
		return
	}
	followFeedParam := database.FollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedId,
		UserID:    user.ID,
	}
	feedFollow, err := cfg.DB.FollowFeed(r.Context(), followFeedParam)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error following feed %v", err))
		return
	}

	respondWithJson(w, 200, databaseFeedFollowToFeedFollow(feedFollow))
}

func (cfg *Config) deleteFeedFollowsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type Parameter struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	params := Parameter{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error decoding request body : %v", err.Error()))
		return
	}

	unfollowParameters := database.UnfollowParams{
		UserID: user.ID,
		FeedID: params.FeedId,
	}
	cfg.DB.Unfollow(r.Context(), unfollowParameters)

	respondWithJson(w, 200, "unfollowed")

}

func (cfg *Config) getFeedFollowsHandler(w http.ResponseWriter, r *http.Request, user database.User) {

	feeds, err := cfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error getting feeds: %v", err))
		return
	}
	respondWithJson(w, 200, databaseFeedFollowsToFeedFollows(feeds))
}
