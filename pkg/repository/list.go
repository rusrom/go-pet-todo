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

func (r *TodoListRepo) GetAllUserLists(userId int) ([]todo.ListTodo, error) {
	var todoLists []todo.ListTodo
	query := fmt.Sprintf(
		"SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1 ORDER BY tl.id",
		listsTable,
		usersListsTable,
	)
	err := r.db.Select(&todoLists, query, userId)
	return todoLists, err
}

func (r *TodoListRepo) GetListDetail(listId int, userId int) (todo.ListTodo, error) {
	var listDetail todo.ListTodo

	query := fmt.Sprintf("SELECT l.id, l.title, l.description FROM %s l INNER JOIN %s ul ON l.id = ul.list_id WHERE l.id = $1 AND user_id = $2",
		listsTable,
		usersListsTable,
	)
	err := r.db.Get(&listDetail, query, listId, userId)

	return listDetail, err
}
