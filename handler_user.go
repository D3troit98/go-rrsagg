package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/D3troit98/go/rrsagg/internal/database"
	"github.com/google/uuid"
)

func (apicfg *apiConfig)handlerCreateUser(w http.ResponseWriter, r *http.Request){

	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err:= decoder.Decode(&params)
	if err != nil {
		respondWithError(w,400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apicfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:uuid.New() ,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})

	if err != nil{
		respondWithError(w,400, fmt.Sprintf("couldn't create user: %v", err))
		return
	}
	respondWithJSON(w , 201, databaseUserToUser(user))
}



func (apicfg *apiConfig)handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User){
	respondWithJSON(w,200, databaseUserToUser(user))
}