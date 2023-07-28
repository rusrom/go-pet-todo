package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "github.com/rusrom/yt-todo"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) CreateUser(u todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, u.Name, u.Username, u.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthRepo) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
