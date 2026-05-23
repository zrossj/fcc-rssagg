package main

import (
	"fmt"
	"net/http"
)

func (apiCfg apiConfig) getFeeds(w http.ResponseWriter, r *http.Request) {

	feedsData, err := apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		respondWithError(w, 401, fmt.Sprintf("Some error gathering data from DB: %v", err))
		return
	}
	respondWithJSON(w, 200, databaseFeedsToFeeds(feedsData))
}
