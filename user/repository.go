package user

import (
	"database/sql"
	"github.com/google/uuid"
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
	log.Printf("Executing query in user-PostgresRepository-CheckEmail: %s | Parameters: %s", query, *email)

	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Printf("Error in user-PostgresRepository-CheckEmail: %v", err)
		return false, err
	}

	return count < 1, nil
}

// AddUser will add the user to the database.
func (p *PostgresRepository) AddUser(user *User) error {
	query := "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)"
	_, err := p.database.Exec(query, user.Id, user.Email, user.Password)
	log.Printf("Execiting query in user-PostgresRepository-AddUser: %s | Parameters: %s, %s, %s", query, user.Id.String(), user.Email, user.Password)

	if err != nil {
		log.Printf("Error in user-PostgresRepository-AddUser: %v", err)
		return err
	}

	return nil
}

// GetUserByEmail will get the user with a matching email.
func (p *PostgresRepository) GetUserByEmail(email *string) (*User, error) {
	query := "SELECT * FROM users WHERE email = $1"
	row := p.database.QueryRow(query, *email)
	log.Printf("Executing query in user-PostgresRepository-GetUserByEmail: %s | Parameters: %s", query, *email)

	var user User
	err := row.Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		log.Printf("Error in user-PostgresRepository-GetUserByEmail: %v", err)
		return nil, err
	}

	return &user, nil
}

// AddToken will add a new token to the table.
func (p *PostgresRepository) AddToken(refreshToken *RefreshToken) error {
	query := "INSERT INTO tokens(id, token, user_id) VALUES ($1, $2, $3)"
	_, err := p.database.Exec(query, refreshToken.Id, refreshToken.UserId, refreshToken.Id)
	log.Printf("Executing query in user-PostgresRepository-AddToken: %s | Parameters: %s, %s, %s", query, refreshToken.Id.String(), refreshToken.StringToken, refreshToken.UserId.String())
	if err != nil {
		log.Printf("Error in user-PostgresRepository-AddToken: %v", err)
	}
	return err
}

// DeleteTokenById will delete the token with a specified id from the table.
func (p *PostgresRepository) DeleteTokenById(id *uuid.UUID) error {
	query := "DELETE FROM tokens WHERE id = $1"
	_, err := p.database.Exec(query, *id)
	log.Printf("Execiting query in user-PostgresRepository-DeleteToken: %s | Parameters: %s", query, *id)

	if err != nil {
		log.Printf("Error in user-PostgresRepository-DeleteToken: %v", err)
	}
	return err
}

func (p *PostgresRepository) DeleteRefreshTokenById(id *uuid.UUID) error {
	query := "DELETE FROM tokens WHERE user_id = $1"
	_, err := p.database.Exec(query, *id)
	log.Printf("Execiting query in user-PostgresRepository-DeleteToken: %s | Parameters: %s", query, *id)
	if err != nil {
		log.Printf("Error in user-PostgresRepository-DeleteToken: %v", err)
	}
	return err
}

func (p *PostgresRepository) GetTokenById(id *uuid.UUID) (*RefreshToken, error) {
	query := "SELECT id, token, user_id FROM tokens WHERE id = $1"
	row := p.database.QueryRow(query, *id)
	log.Printf("Execiting query in user-PostgresRepository-GetTokenById: %s | Parameters: %s", query, id.String())

	var refreshToken RefreshToken
	err := row.Scan(&refreshToken.Id, &refreshToken.StringToken, &refreshToken.UserId)
	if err != nil {
		log.Printf("Error in user-PostgresRepository-GetTokenById: %v", err)
	}
	return &refreshToken, err
}

// NewPostgresRepository will create a new repository with a connection.
func NewPostgresRepository(database *sql.DB) *PostgresRepository {
	return &PostgresRepository{database}
}
