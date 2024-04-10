package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/D3troit98/go/rrsagg/internal/database"
	"github.com/google/uuid"
)

func (apicfg *apiConfig)handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User){

	type parameters struct {
		Name string `json:"name"`
		URL string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err:= decoder.Decode(&params)
	if err != nil {
		respondWithError(w,400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apicfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:uuid.New() ,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url : params.URL,
		UserID: user.ID,
	})

	if err != nil{
		respondWithError(w,400, fmt.Sprintf("couldn't create user: %v", err))
		return
	}
	respondWithJSON(w , 201, databaseFeedToFeed(feed))
}


func (apicfg *apiConfig)handlerGetFeeds(w http.ResponseWriter, r *http.Request){

	
	feed, err := apicfg.DB.GetFeeds(r.Context())

	if err != nil{
		respondWithError(w,400, fmt.Sprintf("couldn't get feeds: %v", err))
		return
	}
	respondWithJSON(w , 201, databaseFeedsToFeeds(feed))
}


