package service

import (
	"context"
	"github.com/bibishkin/bi-notes-rest-api/pkg/repository"
)

type Authorization interface {
	CreateUser(ctx context.Context, username, password string) (int, error)
}

type NoteList interface {
}

type NoteItem interface {
}

type Service struct {
	Authorization
	NoteList
	NoteItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
	}
}
