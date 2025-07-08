package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lywgit/bootdev_chirpy/internal/auth"
	"github.com/lywgit/bootdev_chirpy/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	req := requestBody{}
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode request", err) // 400
		return
	}
	if req.Email == "" {
		respondWithError(w, http.StatusBadRequest, "Email can not be empty", nil) // 400
		return
	}
	if req.Password == "" {
		respondWithError(w, http.StatusBadRequest, "Password can not be empty", nil) // 400
		return
	}
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		respondWithError(w, 500, "could not hash password", err)
	}
	param := database.CreateUserParams{
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}
	user, err := cfg.db.CreateUser(r.Context(), param)
	if err != nil {
		respondWithError(w, 500, "Create user failed", err)
		return
	}
	respondWithJSON(w, http.StatusCreated, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})

}
