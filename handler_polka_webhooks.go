package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/lywgit/bootdev_chirpy/internal/auth"
)

func (cfg *apiConfig) handlerPolkaWebhooks(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized", err) // 401
		return
	}
	if apiKey != cfg.polkaKey {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized", err) // 401
		return
	}

	type requestBody struct {
		Event string `json:"event"`
		Data  struct {
			UserID uuid.UUID `json:"user_id"`
		} `json:"data"`
	}
	decoder := json.NewDecoder(r.Body)
	req := requestBody{}
	err = decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "could not decode request", err)
		return
	}
	if req.Event != "user.upgraded" { // ignore any other kind of events
		w.WriteHeader(http.StatusNoContent)
		return
	}
	userID := req.Data.UserID
	_, err = cfg.db.GetUserByID(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "user not found", err)
		return
	}
	_, err = cfg.db.UpdateUsersSetChirpyRed(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not set chirpy red", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
