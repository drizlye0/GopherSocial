package store

import (
	"context"
	"database/sql"
	"time"
)

func NewMockStore() Storage {
	return Storage{
		Users: UserMockStore{},
	}
}

type UserMockStore struct {
}

func (m UserMockStore) Create(ctx context.Context, tx *sql.Tx, user *User) error {
	return nil
}

func (m UserMockStore) GetByID(ctx context.Context, userID int64) (*User, error) {
	return &User{}, nil
}

func (m UserMockStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	return &User{}, nil
}

func (m UserMockStore) CreateAndInvite(ctx context.Context, user *User, token string, invitationExp time.Duration) error {
	return nil
}

func (m UserMockStore) Activate(ctx context.Context, token string) error {
	return nil
}

func (m UserMockStore) Delete(ctx context.Context, userID int64) error {
	return nil
}
