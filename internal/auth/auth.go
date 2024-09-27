package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API key from the headers of an HTTP request
func GetAPIKey(headers http.Header) (string, error) {
	// Get the authorization from the request headers, if anything goes wrong return an error
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no authorization header included")
	}

	// Split the authorization into two parts (ApiKey and the actual ApiKey),
	// if anything goes wrong return an error
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}
