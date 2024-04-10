package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/D3troit98/go/rrsagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apicfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := apicfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't create feed follow: %v", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apicfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apicfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't get feed follows: %v", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apicfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't parse feed follow ID: %v", err))
		return
	}

	err = apicfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't delete feed follow ID: %v", err))
		return
	}

	respondWithJSON(w, 200, struct{}{})
}
