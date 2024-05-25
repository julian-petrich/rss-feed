package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rss-feed/internal/database"
	"time"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:name`
		URL  string `json:url`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json: %s", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed: %s", err))
		return
	}

	respondWithJSON(w, 200, feed)

}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feed, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feed: %v", err))
		return
	}

	respondWithJSON(w, 200, feed)
}
