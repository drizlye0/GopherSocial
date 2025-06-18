package cache

import (
	"context"

	"github.com/drizlye0/GopherSocial/internal/store"
)

func NewMockStore() Storage {
	return Storage{
		Users: UserMockStore{},
	}
}

type UserMockStore struct {
}

func (m UserMockStore) Get(ctx context.Context, userID int64) (*store.User, error) {
	return &store.User{}, nil
}

func (m UserMockStore) Set(ctx context.Context, user *store.User) error {
	return nil
}
