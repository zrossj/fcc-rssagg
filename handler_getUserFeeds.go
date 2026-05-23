package main

import (
	"fmt"
	"net/http"

	"gibhub.com/jzross/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerGetUserFeeds(w http.ResponseWriter, r *http.Request, u database.User) {

	userData, err := apiCfg.DB.GetUserFeeds(r.Context(), u.ID)
	if err != nil {
		respondWithError(w, 401, fmt.Sprintf("Some error while getting data from DB: %v", err))
	}
	respondWithJSON(w, 200, databaseUserFeedsToUserFeeds(userData))
}
