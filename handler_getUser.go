package main

import (
	"net/http"

	"gibhub.com/jzross/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, u database.User) {
	respondWithJSON(w, 200, returnPOST(u))
}
