package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func replaceProfane(s string) string {
	const replacement = "****"
	cleanedWords := []string{}
	for word := range strings.SplitSeq(s, " ") {
		if _, exists := profaneWords[strings.ToLower(word)]; exists {
			cleanedWords = append(cleanedWords, replacement)
		} else {
			cleanedWords = append(cleanedWords, word)
		}
	}
	return strings.Join(cleanedWords, " ")
}

func handlerValidate(w http.ResponseWriter, r *http.Request) {
	// decode request json and validate
	const maxLength = 140

	type requestBody struct {
		Body string `json:"body"`
	}
	type okResponse struct {
		CleanedBody string `json:"cleaned_body"`
	}
	decoder := json.NewDecoder(r.Body)
	reqData := requestBody{}
	err := decoder.Decode(&reqData)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "error decoding request", err) // 400
		return
	}
	if len(reqData.Body) > maxLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil) // 400
		return
	}
	respondWithJSON(w, http.StatusOK, okResponse{CleanedBody: replaceProfane(reqData.Body)})
}
