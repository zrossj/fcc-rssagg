package main

import (
	"fmt"
	"net/http"

	"gibhub.com/jzross/rssagg/internal/database"
	"gibhub.com/jzross/rssagg/internal/tools"
)

type handlerAuth func(w http.ResponseWriter, r *http.Request, u database.User)

func (apiCfg *apiConfig) middleware(handler handlerAuth) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := tools.GetUserAPIKEY(r.Header)

		if err != nil {
			respondWithError(w, 401, fmt.Sprintf("Error in authentication: %v", err))
		}

		userData, err := apiCfg.DB.GetUser(r.Context(), apikey)

		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Could't get user: %v", err))
		}

		handler(w, r, userData)
	}
}
