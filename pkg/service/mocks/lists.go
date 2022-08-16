package mocks

import entity "github.com/bibishkin/notes-rest-api"

func (m *MockedService) CreateList(userId int, list entity.List) (int, error) {
	args := m.Called(userId, list)
	return args.Get(0).(int), args.Error(1)
}

func (m *MockedService) GetLists(userId int, limit int, offset int) ([]entity.List, error) {
	args := m.Called(userId, limit, offset)
	return args.Get(0).([]entity.List), args.Error(1)
}

func (m *MockedService) GetListById(userId int, listId int) (*entity.List, error) {
	args := m.Called(userId, listId)
	return args.Get(0).(*entity.List), args.Error(1)
}

func (m *MockedService) UpdateList(userId int, listId int, list entity.List) error {
	args := m.Called(userId, listId, list)
	return args.Error(0)
}

func (m *MockedService) DeleteList(userId int, listId int) error {
	args := m.Called(userId, listId)
	return args.Error(0)
}
