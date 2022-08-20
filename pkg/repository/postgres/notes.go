package postgres

import (
	"context"
	"errors"
	"fmt"
	entity "github.com/bibishkind/notes-rest-api"
)

func (r *Repository) CreateNote(userId, listId int, note entity.Note) (int, error) {
	if err := r.checkList(userId, listId); err != nil {
		return 0, err
	}

	createNoteQuery := fmt.Sprintf("INSERT INTO %s (list_id, title, content) VALUES (%d, $1, $2) RETURNING id", notesTable, listId)
	row := r.pool.QueryRow(context.Background(), createNoteQuery, note.Title, note.Content)

	var noteId int
	if err := row.Scan(&noteId); err != nil {
		return 0, err
	}

	return noteId, nil
}

func (r *Repository) GetNotes(userId, listId, limit, offset int) ([]entity.Note, error) {
	if err := r.checkList(userId, listId); err != nil {
		return nil, err
	}

	var getNotesQuery string

	if limit < 0 {
		getNotesQuery = fmt.Sprintf("SELECT * FROM %s WHERE list_id=%d OFFSET %d", notesTable, listId, offset)
	} else {
		getNotesQuery = fmt.Sprintf("SELECT * FROM %s WHERE list_id=%d LIMIT %d OFFSET %d", notesTable, listId, limit, offset)
	}

	rows, err := r.pool.Query(context.Background(), getNotesQuery)
	if err != nil {
		return nil, err
	}

	var notes []entity.Note

	for rows.Next() {
		var note entity.Note
		if err := rows.Scan(&note.Id, &note.ListId, &note.Title, &note.Content); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (r *Repository) GetNoteById(userId, listId, noteId int) (*entity.Note, error) {
	err := r.checkList(userId, listId)
	if err != nil {
		return nil, err
	}

	getNoteQuery := fmt.Sprintf("SELECT * FROM %s WHERE list_id=%d AND id=%d", notesTable, listId, noteId)

	var note entity.Note
	row := r.pool.QueryRow(context.Background(), getNoteQuery)

	if err := row.Scan(&note.Id, &note.ListId, &note.Title, &note.Content); err != nil {
		return nil, err
	}

	return &note, nil
}

func (r *Repository) UpdateNote(userId, listId, noteId int, note entity.Note) error {
	err := r.checkList(userId, listId)
	if err != nil {
		return err
	}

	updateListQuery := fmt.Sprintf("UPDATE %s SET title=$1, content=$2 WHERE list_id=%d AND id=%d", notesTable, listId, noteId)
	_, err = r.pool.Exec(context.Background(), updateListQuery, note.Title, note.Content)

	return err
}

func (r *Repository) DeleteNote(userId, listId, noteId int) error {
	err := r.checkList(userId, listId)
	if err != nil {
		return err
	}

	deleteListQuery := fmt.Sprintf("DELETE FROM %s WHERE list_id=%d AND id=%d", notesTable, listId, noteId)
	_, err = r.pool.Exec(context.Background(), deleteListQuery)

	return err
}

func (r *Repository) checkList(userId, listId int) error {
	checkListQuery := fmt.Sprintf("SELECT FROM %s WHERE user_id=%d AND id=%d", listsTable, userId, listId)

	tag, err := r.pool.Exec(context.Background(), checkListQuery)
	if err != nil {
		return err
	} else if tag.RowsAffected() == 0 {
		return errors.New("user doesn't have this list")
	}

	return nil
}
