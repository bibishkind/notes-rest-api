package repository

import (
	"context"
	"fmt"
	notes "github.com/bibishkin/bi-notes-rest-api"
)

type AuthPostgres struct {
	db *PostgresDB
}

func NewAuthPostgres(db *PostgresDB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(ctx context.Context, username, password string) (int, error) {
	createUserQuery := fmt.Sprintf("INSERT INTO %s (username, password) VALUES ($1, $2) RETURNING id", usersTable)
	row := r.db.Pool.QueryRow(ctx, createUserQuery, username, password)
	
	var id int
	err := row.Scan(&id)

	return id, err
}

func (r *AuthPostgres) GetUser(ctx context.Context, username, password string) (notes.User, error) {
	getUserQuery := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 AND password=$2", usersTable)
	row := r.db.Pool.QueryRow(ctx, getUserQuery, username, password)

	var user notes.User
	err := row.Scan(&user.Id, &user.Username, &user.Password)

	return user, err
}
