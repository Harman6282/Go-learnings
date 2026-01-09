package store

import (
	"context"
	"database/sql"
)

type Posts interface {
	Create(context.Context, *Post) error
}

type Users interface {
	Create(context.Context, *User) error
}

type Storage struct {
	Posts
	Users
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		Users: &UserStore{db},
	}
}
