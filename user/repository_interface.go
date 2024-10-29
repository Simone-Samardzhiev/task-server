package user

import "github.com/google/uuid"

// Repository that defines the methods for a user repository.
type Repository interface {
	// CheckEmail will check if the email is already in use.
	CheckEmail(*string) (bool, error)

	// AddUser will add a new user.
	AddUser(user *User) error

	// GetUserByEmail will return a user with a specific email.
	GetUserByEmail(*string) (*User, error)

	// AddToken will add a new token.
	AddToken(*RefreshToken) error

	// DeleteTokenById will delete a token with a specific id.
	DeleteTokenById(uuid *uuid.UUID) error

	// GetTokenById will get the token with a specific id.
	GetTokenById(uuid *uuid.UUID) (*RefreshToken, error)
}
