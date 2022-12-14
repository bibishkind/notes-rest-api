package repository

import entity "github.com/bibishkind/notes-rest-api"

type AuthImplementation interface {
	CreateUser(username, password string) (int, error)
	GetUser(username, password string) (*entity.User, error)
}

type ListsImplementation interface {
	CreateList(userId int, list entity.List) (int, error)
	GetLists(userId, limit, offset int) ([]entity.List, error)
	GetListById(userId, listId int) (*entity.List, error)
	UpdateList(userId, listId int, list entity.List) error
	DeleteList(userId, listId int) error
}

type NotesImplementation interface {
	CreateNote(userId, listId int, note entity.Note) (int, error)
	GetNotes(userId, listId, limit, offset int) ([]entity.Note, error)
	GetNoteById(userId, listId, noteId int) (*entity.Note, error)
	UpdateNote(userId, listId, noteId int, note entity.Note) error
	DeleteNote(userId, listId, noteId int) error
}

type Implementation interface {
	AuthImplementation
	ListsImplementation
	NotesImplementation
}
