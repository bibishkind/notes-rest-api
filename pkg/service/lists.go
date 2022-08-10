package service

import entity "github.com/bibishkin/bi-notes-rest-api"

func (s *Service) CreateList(userId int, list entity.List) (int, error) {
	return s.repository.CreateList(userId, list)
}

func (s *Service) GetLists(userId, limit, offset int) ([]entity.List, error) {
	return s.repository.GetLists(userId, limit, offset)
}

func (s *Service) GetListById(userId, listId int) (*entity.List, error) {
	return s.repository.GetListById(userId, listId)
}

func (s *Service) UpdateList(userId, listId int, list entity.List) error {
	return s.repository.UpdateList(userId, listId, list)
}

func (s *Service) DeleteList(userId, listId int) error {
	return s.repository.DeleteList(userId, listId)
}
