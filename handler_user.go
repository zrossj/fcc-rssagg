package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gibhub.com/jzross/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	type parameters struct {
		Name string
	}

	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error in decoding JSON: %v", err))
		return
	}

	// se nao houve erro em decodar o post do usuario, podemos usar os dados recebidos para realizar input no db.

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		ModifiedAt: time.Now().UTC(),
		Name:       params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error in creating user on DB: %v", err))
		return
	}

	respondWithJSON(w, 200, returnPOST(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, u database.User) {
	respondWithJSON(w, 200, returnPOST(u))
}

func (apiCfg *apiConfig) handleGetRecentPosts(w http.ResponseWriter, r *http.Request, u database.User) {

	recentPosts, err := apiCfg.DB.GetLastPosts(
		r.Context(), database.GetLastPostsParams{
			ID:    u.ID,
			Limit: 10,
		})
	if err != nil {
		respondWithError(w, 401, fmt.Sprintf("Errro while trying to get last posts for user. %v", err))
	}

	respondWithJSON(w, 200, databasePostToPosts(recentPosts))
}

func (apiCfg *apiConfig) handlerGetUserFeeds(w http.ResponseWriter, r *http.Request, u database.User) {

	userData, err := apiCfg.DB.GetUserFeeds(r.Context(), u.ID)
	if err != nil {
		respondWithError(w, 401, fmt.Sprintf("Some error while getting data from DB: %v", err))
	}
	respondWithJSON(w, 200, databaseUserFeedsToUserFeeds(userData))
}
