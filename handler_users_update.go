package main

import (
	"encoding/json"
	"net/http"

	"github.com/lywgit/bootdev_chirpy/internal/auth"
	"github.com/lywgit/bootdev_chirpy/internal/database"
)

func (cfg *apiConfig) handlerUsersUpdate(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "could not get access token", err)
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	type requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	req := requestBody{}
	err = decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode request", err)
		return
	}
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not hash password", err)
		return
	}

	user, err := cfg.db.UpdateUsersByID(r.Context(), database.UpdateUsersByIDParams{
		ID:             userID,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not update users table", err)
		return
	}
	respondWithJSON(w, http.StatusOK, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})

}
