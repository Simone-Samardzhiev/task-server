package user

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

// PostgresRepository is an implementation of Repository.
type PostgresRepository struct {
	database *sql.DB
}

// CheckEmail will check in the database if the email is already in use.
func (p *PostgresRepository) CheckEmail(email *string) (bool, error) {
	query := "SELECT COUNT(id) FROM users WHERE email = $1"
	row := p.database.QueryRow(query, *email)
	log.Printf("Executing query in CheckEmail: %s | Parameters: %s", query, *email)

	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Printf("Error in CheckEmail line 22 | Error: %v", err)
		return false, err
	}

	return count < 1, nil
}

// AddUser will add the user to the database.
func (p *PostgresRepository) AddUser(user *User) error {
	query := "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)"
	_, err := p.database.Exec(query, user.Id, user.Email, user.Password)
	log.Printf("Execiting query in AddUser: %s | Parameters: %s, %s, %s", query, user.Id.String(), user.Email, user.Password)

	if err != nil {
		log.Printf("Error in AddUser line 36 | Error: %v", err)
		return err
	}

	return nil
}

// GetUserByEmail will get the user with a matching email.
func (p *PostgresRepository) GetUserByEmail(email *string) (*User, error) {
	query := "SELECT * FROM users WHERE email = $1"
	row := p.database.QueryRow(query, *email)
	log.Printf("Executing query in GetUserByEmail: %s | Parameters: %s", query, *email)

	var user User
	err := row.Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		log.Printf("Error in GetUserByEmail line 52 | Error: %v", err)
		return nil, err
	}

	return &user, nil
}

// NewPostgresRepository will create a new repository with a connection.
func NewPostgresRepository(database *sql.DB) *PostgresRepository {
	return &PostgresRepository{database}
}
