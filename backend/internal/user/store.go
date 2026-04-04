package user

import "context"

// Store loads users for authentication and domain logic.
type Store interface {
	GetByEmail(ctx context.Context, email string) (User, error)
}
