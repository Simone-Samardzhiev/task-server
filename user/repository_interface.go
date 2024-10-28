package user

// Repository that defines the methods for a user repository.
type Repository interface {
	// CheckEmail will check if the email is already in use.
	CheckEmail(*string) (bool, error)

	// AddUser will add a new user.
	AddUser(user *User) error

	// GetUserByEmail will return a user with a specific email.
	GetUserByEmail(*string) (*User, error)
}
