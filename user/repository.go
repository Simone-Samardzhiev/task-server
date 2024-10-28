package user

import "database/sql"

// PostgresRepository is an implementation of Repository.
type PostgresRepository struct {
	database sql.DB
}

// CheckEmail will check in the database if the email is already in use.
func (p *PostgresRepository) CheckEmail(email *string) (bool, error) {
	query := "SELECT COUNT(id) FROM users WHERE email = $1"
	row := p.database.QueryRow(query, *email)
	var count int

	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	return count < 1, nil
}

// AddUser will add the user to the database.
func (p *PostgresRepository) AddUser(user *User) error {
	query := "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)"
	_, err := p.database.Exec(query, user.Id, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByEmail will get the user with a matching email.
func (p *PostgresRepository) GetUserByEmail(email *string) (*User, error) {
	query := "SELECT * FROM users WHERE email = $1"
	row := p.database.QueryRow(query, *email)
	var user User

	err := row.Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
