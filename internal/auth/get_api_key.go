package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("authorization header not found")
	}
	if !strings.HasPrefix(authHeader, "ApiKey ") {
		return "", fmt.Errorf("malformed authorization header")
	}
	apiString := strings.TrimPrefix(authHeader, "ApiKey ")
	return apiString, nil
}
