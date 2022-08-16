package mocks

import entity "github.com/bibishkin/notes-rest-api"

func (m *MockedService) CreateNote(userId int, listId int, note entity.Note) (int, error) {
	args := m.Called(userId, listId, note)
	return args.Get(0).(int), args.Error(1)
}

func (m *MockedService) GetNotes(userId int, listId int, limit int, offset int) ([]entity.Note, error) {
	args := m.Called(userId, listId, limit, offset)
	return args.Get(0).([]entity.Note), args.Error(1)
}

func (m *MockedService) GetNoteById(userId int, listId int, noteId int) (*entity.Note, error) {
	args := m.Called(userId, listId, noteId)
	return args.Get(0).(*entity.Note), args.Error(1)
}

func (m *MockedService) UpdateNote(userId int, listId int, noteId int, note entity.Note) error {
	args := m.Called(userId, listId, noteId)
	return args.Error(0)
}

func (m *MockedService) DeleteNote(userId int, listId int, noteId int) error {
	args := m.Called(userId, listId, noteId)
	return args.Error(0)
}
