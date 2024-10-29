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
