package middleware

import "github.com/google/uuid"

// Authenticator defines methods for an authenticator.
type Authenticator interface {
	// CreateRefreshToken will create a new token with an id.
	CreateRefreshToken(*uuid.UUID) (*string, error)

	// CheckRefreshToken will check a refresh token and return its id.
	CheckRefreshToken(*string) (*uuid.UUID, error)

	// CreateAccessToken will return a new access token with an id.
	CreateAccessToken(*uuid.UUID) (*string, error)

	// CheckAccessToken will check an access token and return its id.
	CheckAccessToken(*string) (*uuid.UUID, error)
}
