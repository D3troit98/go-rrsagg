package main

import (
	"fmt"
	"net/http"

	"github.com/D3troit98/go/rrsagg/internal/auth"
	"github.com/D3troit98/go/rrsagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)


func (apicfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		apikey,err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("auth error: %v", err))
		return
	}

	user, err := apicfg.DB.GetUserByAPIKey(r.Context(), apikey)
	if err != nil {
		respondWithError(w,400, fmt.Sprintf("Couldn't get user: %v",err))
		return

	}
	handler(w,r,user)
	}
}