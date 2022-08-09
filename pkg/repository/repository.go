package repository

import entity "github.com/bibishkin/bi-notes-rest-api"

type Repository interface {
	CreateUser(username, password string) (int, error)
	GetUser(username, password string) (*entity.User, error)
}
