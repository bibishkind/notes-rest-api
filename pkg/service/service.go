package service

import (
	entity "github.com/bibishkind/notes-rest-api"
	"github.com/bibishkind/notes-rest-api/pkg/repository"
)

type Service struct {
	repository repository.Implementation
}

type AuthImplementation interface {
	CreateUser(username string, password string) (int, error)
	GenerateJWT(username string, password string) (string, error)
	ParseJWT(tokenString string) (int, error)
}

type ListsImplementation interface {
	CreateList(userId int, list entity.List) (int, error)
	GetLists(userId int, limit int, offset int) ([]entity.List, error)
	GetListById(userId int, listId int) (*entity.List, error)
	UpdateList(userId int, listId int, list entity.List) error
	DeleteList(userId int, listId int) error
}

type NotesImplementation interface {
	CreateNote(userId int, listId int, note entity.Note) (int, error)
	GetNotes(userId int, listId int, limit int, offset int) ([]entity.Note, error)
	GetNoteById(userId int, listId int, noteId int) (*entity.Note, error)
	UpdateNote(userId int, listId int, noteId int, note entity.Note) error
	DeleteNote(userId int, listId int, noteId int) error
}

type Implementation interface {
	AuthImplementation
	ListsImplementation
	NotesImplementation
}

func NewService(repository repository.Implementation) Implementation {
	return &Service{repository: repository}
}
