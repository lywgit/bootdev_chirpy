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
	type returnValsOK struct {
		CleanedBody string `json:"cleaned_body"`
	}
	type returnValsError struct {
		Error string `json:"error"`
	}

	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	reqData := requestBody{}
	err := decoder.Decode(&reqData)
	if err != nil {
		respData, _ := json.Marshal(returnValsError{Error: "Something went Wrong"})
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write(respData)
		return
	}
	if len(reqData.Body) > maxLength {
		respData, _ := json.Marshal(returnValsError{Error: "Chirp is too long"})
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write(respData)
		return
	}

	dat, _ := json.Marshal(returnValsOK{CleanedBody: replaceProfane(reqData.Body)})
	w.WriteHeader(http.StatusOK) // 200
	w.Write(dat)

}
