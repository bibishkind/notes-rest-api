package repository

import entity "github.com/bibishkin/notes-rest-api"

type Repository interface {
	CreateUser(username, password string) (int, error)
	GetUser(username, password string) (*entity.User, error)

	CreateList(userId int, list entity.List) (int, error)
	GetLists(userId, limit, offset int) ([]entity.List, error)
	GetListById(userId, listId int) (*entity.List, error)
	UpdateList(userId, listId int, list entity.List) error
	DeleteList(userId, listId int) error

	CreateNote(userId, listId int, note entity.Note) (int, error)
	GetNotes(userId, listId, limit, offset int) ([]entity.Note, error)
	GetNoteById(userId, listId, noteId int) (*entity.Note, error)
	UpdateNote(userId, listId, noteId int, note entity.Note) error
	DeleteNote(userId, listId, noteId int) error
}
