package user

import "github.com/google/uuid"

// User is a model representing a user.
type User struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

// NewUser is a struct representing a new user.
type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
