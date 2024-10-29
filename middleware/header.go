package middleware

import (
	"errors"
	"net/http"
	"strings"
)

// GetTokenFromHeader will return the token from the request header.
func GetTokenFromHeader(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")

	if header == "" {
		return "", errors.New("no authorization header")
	}

	if !strings.HasPrefix(header, "Bearer ") {
		return "", errors.New("invalid authorization header")
	}

	return strings.TrimPrefix(header, "Bearer "), nil
}
