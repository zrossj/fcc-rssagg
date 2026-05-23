package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gibhub.com/jzross/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, u database.User) {

	decoder := json.NewDecoder(r.Body)

	type parameters struct {
		Name string
		Url  string
	}

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error in JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		ModifiedAt: time.Now().UTC(),
		Name:       params.Name,
		Url:        params.Url,
		UserID:     u.ID,
	})

	if err != nil {
		respondWithError(
			w, 401, fmt.Sprintf("Some error while creating the feed on DB: %v", err))
	}

	respondWithJSON(w, 200, databaseFeedToFeed(feed))
}
