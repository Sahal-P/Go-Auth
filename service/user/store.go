package user

import (
	"context"
	"fmt"

	"github.com/Sahal-P/Go-Auth/db"
	"github.com/Sahal-P/Go-Auth/types"
	"github.com/jackc/pgx/v4"
)

type Store struct {
	db *db.PostgreSQLStorage
}

func NewStore(db *db.PostgreSQLStorage) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	user := &types.User{}

	query := `SELECT id, first_name, last_name, email, password, created_at FROM users WHERE email = $1`

	err := s.db.Conn.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			// If no rows are found, return a nil user with no error
			return nil, nil
		}
		// Return any other error that occurred
		return nil, fmt.Errorf("failed to query user by email: %w", err)
	}

	// Return the user found in the database
	return user, nil
}

func (s *Store) CreateUser(user *types.User) (*types.User, error) {
	query := `INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id, created_at`

	err := s.db.Conn.QueryRow(context.Background(), query, user.FirstName, user.LastName, user.Email, user.Password).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}