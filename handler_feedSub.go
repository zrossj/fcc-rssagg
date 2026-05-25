package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gibhub.com/jzross/rssagg/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleSubsToFeed(w http.ResponseWriter, r *http.Request, u database.User) {

	decoder := json.NewDecoder(r.Body)

	type parameters struct {
		FeedID uuid.UUID
	}

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error in JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.SubFeed(r.Context(), database.SubFeedParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		ModifiedAt: time.Now().UTC(),
		UserID:     u.ID,
		FeedID:     params.FeedID})

	if err != nil {
		respondWithError(
			w, 401, fmt.Sprintf("Some error while creating the feed on DB: %v", err))
	}

	respondWithJSON(w, 200, databaseFeedFollowToFeedFollow(feed))
}

func (apiCfg *apiConfig) handleDeleteSub(w http.ResponseWriter, r *http.Request, u database.User) {

	feedFollowIDStr := chi.URLParam(r, "FeedID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse FeedID: %v", err))
	}

	err = apiCfg.DB.DeleteUserSub(r.Context(), database.DeleteUserSubParams{
		UserID: u.ID,
		FeedID: feedFollowID,
	})

	if err != nil {
		respondWithError(
			w, 401, fmt.Sprintf("Error while trying to unscribe user from feed: %v", err))
	}

	respondWithJSON(w, 200, struct{}{})
}
