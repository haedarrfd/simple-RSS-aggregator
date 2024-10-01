package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Get an API key from headers of an incoming HTTP request
func GetAPIKey(headers http.Header) (string, error) {
	// Retrieve the value of Authorization header
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no authorization header included")
	}

	// Split the string of authHeader into two parts (ApiKey and the actual ApiKey)
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}
