package main

import (
	"fmt"
	"net/http"

	"gibhub.com/jzross/rssagg/internal/tools"
)

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {

	apikey, err := tools.GetUserAPIKEY(r.Header)

	if err != nil {
		respondWithError(w, 401, fmt.Sprintf("Error in authentication: %v", err))
		return
	}

	userData, err := apiCfg.DB.GetUser(r.Context(), apikey)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could't get user: %v", err))
		return
	}

	respondWithJSON(w, 200, returnPOST(userData))
}
