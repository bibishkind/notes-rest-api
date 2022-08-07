package repository

import (
	"context"
	notes "github.com/bibishkin/bi-notes-rest-api"
)

type Authorization interface {
	CreateUser(ctx context.Context, username, password string) (int, error)
	GetUser(ctx context.Context, username, password string) (notes.User, error)
}

type NoteItem interface {
}

type NoteList interface {
}

type Repository struct {
	Authorization
	NoteList
	NoteItem
}

func NewRepository(db *PostgresDB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
