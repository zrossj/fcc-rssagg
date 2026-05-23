package main

import (
	"fmt"
	"net/http"

	"gibhub.com/jzross/rssagg/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

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
