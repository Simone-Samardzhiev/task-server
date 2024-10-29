package user

// Service defines the methods for a user service.
type Service interface {
	// Login will return a refresh token.
	Login(*WithoutIdUser) (*string, error)

	// Register will register the user.
	Register(*WithoutIdUser) error

	// RefreshTokens will return a new refresh token and access token using the access token.
	RefreshTokens(*string) (*TokenGroup, error)
}
