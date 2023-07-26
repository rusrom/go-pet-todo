package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "github.com/rusrom/yt-todo"
)

type TodoItemRepo struct {
	db *sqlx.DB
}

func NewTodoItemRepo(db *sqlx.DB) *TodoItemRepo {
	return &TodoItemRepo{db: db}
}

func (r *TodoItemRepo) CreateNewItem(i todo.ItemTodo, listId int) (int, error) {
	trx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	queryNewItem := fmt.Sprintf("INSERT INTO %s (title, description, done) values ($1, $2, $3) RETURNING id", itemsTable)
	row := trx.QueryRow(queryNewItem, i.Title, i.Description, i.Done)
	if err := row.Scan(&id); err != nil {
		_ = trx.Rollback()
		return 0, err
	}

	queryListItem := fmt.Sprintf("INSERT INTO %s (item_id, list_id) values ($1, $2)", listsItemsTable)
	_, err = trx.Exec(queryListItem, id, listId)
	if err != nil {
		_ = trx.Rollback()
		return 0, err
	}

	return id, trx.Commit()
}
