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
	AccessToken string `json:"accessToken"`
	Password    string `json:"password"`
}
