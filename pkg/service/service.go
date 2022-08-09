package service

import "github.com/bibishkin/bi-notes-rest-api/pkg/repository"

type Service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *Service {
	return &Service{repository: repository}
}
