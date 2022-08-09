package postgres

import (
	"context"
	"fmt"
	entity "github.com/bibishkin/bi-notes-rest-api"
)

func (r *Repository) CreateUser(username, password string) (int, error) {
	createUserQuery := fmt.Sprintf("INSERT INTO %s (username, password) VALUES ($1, $2) RETURNING id", usersTable)

	row := r.pool.QueryRow(context.Background(), createUserQuery, username, password)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) GetUser(username, password string) (*entity.User, error) {
	getUserQuery := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 AND password=$2", usersTable)
	fmt.Println(username, password)
	row := r.pool.QueryRow(context.Background(), getUserQuery, username, password)

	var user entity.User
	if err := row.Scan(&user.Id, &user.Username, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}
