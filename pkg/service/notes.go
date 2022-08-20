package service

import entity "github.com/bibishkind/notes-rest-api"

func (s *Service) CreateNote(userId, listId int, note entity.Note) (int, error) {
	return s.repository.CreateNote(userId, listId, note)
}

func (s *Service) GetNotes(userId, listId, limit, offset int) ([]entity.Note, error) {
	return s.repository.GetNotes(userId, listId, limit, offset)
}

func (s *Service) GetNoteById(userId, listId, noteId int) (*entity.Note, error) {
	return s.repository.GetNoteById(userId, listId, noteId)
}

func (s *Service) UpdateNote(userId, listId, noteId int, note entity.Note) error {
	return s.repository.UpdateNote(userId, listId, noteId, note)
}

func (s *Service) DeleteNote(userId, listId, noteId int) error {
	return s.repository.DeleteNote(userId, listId, noteId)
}
