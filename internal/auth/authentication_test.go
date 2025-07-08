package auth

import (
	"testing"
)

func TestCheckPasswordHash(t *testing.T) {
	tests := []struct {
		storedPassword  string
		loginPassword   string
		expectedSuccess bool
	}{
		{
			storedPassword:  "123456",
			loginPassword:   "123456",
			expectedSuccess: true,
		},
		{
			storedPassword:  "123456",
			loginPassword:   "000000",
			expectedSuccess: false,
		},
		{
			storedPassword:  "!@#$%",
			loginPassword:   "!@#$%",
			expectedSuccess: true,
		},
		{
			storedPassword:  "",
			loginPassword:   "",
			expectedSuccess: true,
		},
	}

	for _, test := range tests {
		storedHash, _ := HashPassword(test.storedPassword)
		var actualSuccess bool
		if err := CheckPasswordHash(test.loginPassword, storedHash); err != nil {
			actualSuccess = false
		} else {
			actualSuccess = true
		}
		if test.expectedSuccess != actualSuccess {
			t.Errorf("Stored password: \"%s\" + login password: \"%s\" Expect: \"%t\", Got: \"%t\"",
				test.storedPassword, test.loginPassword, test.expectedSuccess, actualSuccess)
		}
	}

}
