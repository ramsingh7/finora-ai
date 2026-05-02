package user

import "errors"

// User is a persisted account (password hash never leaves the service layer).
type User struct {
	ID           string
	Email        string
	PasswordHash string
}

// ErrNotFound is returned when no row matches.
var ErrNotFound = errors.New("user not found")
