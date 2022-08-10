package postgres

import (
	"context"
	"fmt"
	entity "github.com/bibishkin/bi-notes-rest-api"
)

func (r *Repository) CreateList(userId int, list entity.List) (int, error) {
	createListQuery := fmt.Sprintf("INSERT INTO %s (user_id, title, description) VALUES (%d, $1, $2) RETURNING id", listsTable, userId)
	row := r.pool.QueryRow(context.Background(), createListQuery, list.Title, list.Description)

	var listId int
	if err := row.Scan(&listId); err != nil {
		return 0, err
	}

	return listId, nil
}

func (r *Repository) GetLists(userId int, limit int, offset int) ([]entity.List, error) {
	var getListsQuery string

	if limit < 0 {
		getListsQuery = fmt.Sprintf("SELECT * FROM %s WHERE user_id=%d OFFSET %d", listsTable, userId, offset)
	} else {
		getListsQuery = fmt.Sprintf("SELECT * FROM %s WHERE user_id=%d LIMIT %d OFFSET %d", listsTable, userId, limit, offset)
	}

	rows, err := r.pool.Query(context.Background(), getListsQuery)
	if err != nil {
		return nil, err
	}

	var lists []entity.List

	for rows.Next() {
		var list entity.List
		if err := rows.Scan(&list.Id, &list.UserId, &list.Title, &list.Description); err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}

	return lists, nil
}

func (r *Repository) GetListById(userId, listId int) (*entity.List, error) {
	getListQuery := fmt.Sprintf("SELECT * FROM %s WHERE user_id = %d AND id = %d", listsTable, userId, listId)

	var list entity.List
	row := r.pool.QueryRow(context.Background(), getListQuery)

	if err := row.Scan(&list.Id, &list.UserId, &list.Title, &list.Description); err != nil {
		return nil, err
	}

	return &list, nil
}

func (r *Repository) UpdateList(userId, listId int, list entity.List) error {
	updateListQuery := fmt.Sprintf("UPDATE %s SET title=$1, description=$2 WHERE user_id=%d AND id=%d", listsTable, userId, listId)
	_, err := r.pool.Exec(context.Background(), updateListQuery, list.Title, list.Description)
	return err
}

func (r *Repository) DeleteList(userId, listId int) error {
	deleteListQuery := fmt.Sprintf("DELETE FROM %s WHERE user_id=%d AND id=%d", listsTable, userId, listId)
	_, err := r.pool.Exec(context.Background(), deleteListQuery)
	return err
}
