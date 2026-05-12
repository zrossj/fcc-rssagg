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
