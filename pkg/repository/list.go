package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "github.com/rusrom/yt-todo"
)

type TodoListRepo struct {
	db *sqlx.DB
}

func NewTodoListRepo(db *sqlx.DB) *TodoListRepo {
	return &TodoListRepo{db: db}
}

func (r *TodoListRepo) CreateNewList(l todo.ListTodo, userId int) (int, error) {
	trx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	queryNewList := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", listsTable)
	row := r.db.QueryRow(queryNewList, l.Title, l.Description)
	if err := row.Scan(&id); err != nil {
		_ = trx.Rollback()
		return 0, err
	}

	queryUserList := fmt.Sprintf("INSERT INTO %s (user_id, list_id) values ($1, $2) RETURNING id", usersListsTable)
	_, err = trx.Exec(queryUserList, userId, id)
	if err != nil {
		_ = trx.Rollback()
		return 0, err
	}

	return id, trx.Commit()
}
