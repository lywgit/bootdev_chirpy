package auth

import (
	"strings"
	"testing"

	"net/http"
)

func TestGetBearerToken(t *testing.T) {
	// initialize test variables
	validHeader := http.Header{}
	validHeader.Add("Authorization", "Bearer MOCK_TOKEN_STRING")

	noAuthFieldHeader := http.Header{}

	malformedTokenHeader := http.Header{}
	malformedTokenHeader.Add("Authorization", "NotBearer MOCK_TOKEN_STRING")

	tests := []struct {
		name          string
		inputHeader   http.Header
		expectedToken string
		expectedError string // expected error message
	}{
		{
			name:          "Valid token",
			inputHeader:   validHeader,
			expectedToken: "MOCK_TOKEN_STRING",
			expectedError: "",
		},
		{
			name:          "Header without Authorization field",
			inputHeader:   noAuthFieldHeader,
			expectedToken: "",
			expectedError: "authorization header not found",
		},
		{
			name:          "Header without Authorization field",
			inputHeader:   malformedTokenHeader,
			expectedToken: "",
			expectedError: "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualToken, err := GetBearerToken(tt.inputHeader)
			if tt.expectedError != "" {
				if err == nil {
					t.Errorf("Expect error containg %s for case %s but got none", tt.expectedError, tt.name)
				}
				if !strings.Contains(err.Error(), tt.expectedError) {
					t.Errorf("Expect error containg %s for case %s but got %v", tt.expectedError, tt.name, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for %s: %v", tt.name, err)
				}
				if actualToken != tt.expectedToken {
					t.Errorf("Expected BearerToken %s, got %s for %s", tt.expectedToken, actualToken, tt.name)
				}
			}
		})
	}

}
