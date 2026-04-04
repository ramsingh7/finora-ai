package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresStore implements Store against PostgreSQL.
type PostgresStore struct {
	pool *pgxpool.Pool
}

// NewPostgresStore wraps a pgx pool.
func NewPostgresStore(pool *pgxpool.Pool) *PostgresStore {
	return &PostgresStore{pool: pool}
}

// GetByEmail returns a user by unique email.
func (s *PostgresStore) GetByEmail(ctx context.Context, email string) (User, error) {
	const q = `SELECT id::text, email, password_hash FROM users WHERE lower(email) = lower($1)`
	var u User
	err := s.pool.QueryRow(ctx, q, email).Scan(&u.ID, &u.Email, &u.PasswordHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return User{}, ErrNotFound
		}
		return User{}, err
	}
	return u, nil
}
