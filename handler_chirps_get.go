package main

import (
	"net/http"
	"sort"

	"github.com/google/uuid"
	"github.com/lywgit/bootdev_chirpy/internal/database"
)

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	var err error
	var chirpsData []database.Chirp
	authorIDStr := r.URL.Query().Get("author_id")
	sortOrder := r.URL.Query().Get("sort")
	if sortOrder == "" {
		sortOrder = "asc"
	}
	if sortOrder != "asc" && sortOrder != "desc" {
		respondWithError(w, http.StatusBadRequest, "invalid sort parameter", nil)
		return
	}

	if authorIDStr == "" {
		chirpsData, err = cfg.db.GetChirps(r.Context())
		if err != nil {
			respondWithError(w, 500, "could not get chirps", err)
			return
		}
	} else {
		authorID, err := uuid.Parse(authorIDStr)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "invalid author_id", err)
			return
		}
		chirpsData, err = cfg.db.GetChirpsByUserID(r.Context(), authorID)
		if err != nil {
			respondWithError(w, 500, "could not get chirps", err)
			return
		}
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
	if sortOrder == "asc" {
		sort.Slice(chirps, func(i int, j int) bool { return chirps[i].CreatedAt.Before(chirps[j].CreatedAt) })
	} else {
		sort.Slice(chirps, func(i int, j int) bool { return chirps[i].CreatedAt.After(chirps[j].CreatedAt) })
	}

	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handlerChirpsGetByID(w http.ResponseWriter, r *http.Request) {
	chirpIDStr := r.PathValue("chirpID")
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
		respondWithError(w, http.StatusNotFound, "not found", err)
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
