package store

import (
	"context"
	"database/sql"
	"errors"
)


var (
	ErrNotFound = errors.New("Record not found")
)

type Posts interface {
	GetByID(context.Context, int64) (*Post, error)
	Create(context.Context, *Post) error
	Delete(context.Context, int64) error
	Update(context.Context, *Post) error
}

type Users interface {
	Create(context.Context, *User) error
}

type Comments interface {
	GetByPostID(context.Context, int64) ([]Comment, error)
}

type Storage struct {
	Posts
	Users
	Comments
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		Users: &UserStore{db},
		Comments: &CommentStore{db},
	}
}
