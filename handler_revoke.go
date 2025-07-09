package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/lywgit/bootdev_chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "could not get token from request", err)
		return
	}
	userID, _ := cfg.db.GetUserFromRefreshToken(r.Context(), token)
	if userID == uuid.Nil {
		respondWithError(w, http.StatusBadRequest, "Not an active refresh token", nil)
		return
	}
	err = cfg.db.RevokeRefreshToken(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error revoking token", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
