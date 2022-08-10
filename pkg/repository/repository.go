package repository

import entity "github.com/bibishkin/bi-notes-rest-api"

type Repository interface {
	CreateUser(username, password string) (int, error)
	GetUser(username, password string) (*entity.User, error)

	CreateList(userId int, list entity.List) (int, error)
	GetLists(userId int, limit int, offset int) ([]entity.List, error)
	GetListById(userId, listId int) (*entity.List, error)
	UpdateList(userId, listId int, list entity.List) error
	DeleteList(userId, listId int) error
}
