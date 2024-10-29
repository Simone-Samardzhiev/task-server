package user

import "github.com/google/uuid"

// User is a model representing a user.
type User struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

// WithoutIdUser is a struct representing a user without id.
type WithoutIdUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TokenGroup is a struct holding the two tokens.
type TokenGroup struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// RefreshToken is a struct representing the information about a token.
type RefreshToken struct {
	Id          uuid.UUID
	StringToken string
	UserId      uuid.UUID
}
