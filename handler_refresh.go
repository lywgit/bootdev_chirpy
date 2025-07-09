package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lywgit/bootdev_chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "could not get token from request %s", err)
		return
	}
	userID, _ := cfg.db.GetUserFromRefreshToken(r.Context(), token)
	if userID == uuid.Nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	accessToken, err := auth.MakeJWT(userID, cfg.jwtSecret, time.Duration(cfg.accessTokenExpireSec*int(time.Second)))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not make access token:", err)
		return
	}
	type refreshResp struct {
		Token string `json:"token"`
	}
	respondWithJSON(w, http.StatusOK, refreshResp{Token: accessToken})
}
