package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/lywgit/bootdev_chirpy/internal/auth"
	"github.com/lywgit/bootdev_chirpy/internal/database"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Email        string `json:"email"`
		Password     string `json:"password"`
		ExpiresInSec *int   `json:"expires_in_seconds,omitempty"`
	}
	decoder := json.NewDecoder(r.Body)
	req := requestBody{}
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode request", err) // 400
		return
	}
	user, _ := cfg.db.GetUserByEmail(r.Context(), req.Email)
	if user == (database.User{}) {
		log.Println("Login failed: email not found:", req.Email)
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", nil)
		return
	}

	if err := auth.CheckPasswordHash(req.Password, user.HashedPassword); err != nil {
		log.Println("Login failed: incorrect password:", err)
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", nil)
		return
	}
	// create token for user
	expiresInSec := 3600 // defaut 1 hour
	if req.ExpiresInSec != nil && (*req.ExpiresInSec) <= expiresInSec {
		expiresInSec = *req.ExpiresInSec
	}
	tokenString, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Duration(expiresInSec*int(time.Second)))
	if err != nil {
		respondWithError(w, 500, "could not generate token", err)
	}
	respondWithJSON(w, http.StatusOK, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
		Token:     tokenString,
	})

}
