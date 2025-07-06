package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/lywgit/bootdev_chirpy/internal/database"
)

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	chirpsData, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, 500, "Create chirp failed", err)
		return
	}
	var chirps []Chirp
	for _, chirp := range chirpsData {
		chirps = append(chirps, Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		})
	}
	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handlerChirpsGetByID(w http.ResponseWriter, r *http.Request) {
	chirpIDStr := r.PathValue("chirpID")
	log.Println("[get chirp by id]:", chirpIDStr)
	if chirpIDStr == "" {
		respondWithError(w, http.StatusBadRequest, "chirp id is empty", nil)
		return
	}
	chirpID, err := uuid.Parse(chirpIDStr)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "not a valid uuid", err)
		return
	}
	chirp, err := cfg.db.GetChirpByID(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not get chirp by id", err)
		return
	}
	if chirp == (database.Chirp{}) {
		respondWithError(w, http.StatusNotFound, "not found", nil)
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}
