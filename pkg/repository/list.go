package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "github.com/rusrom/yt-todo"
	"strings"
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

func (r *TodoListRepo) DeleteList(listId int, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s l USING %s ul WHERE l.id = ul.list_id AND l.id = $1 AND ul.user_id = $2",
		listsTable,
		usersListsTable,
	)
	_, err := r.db.Exec(query, listId, userId)
	return err
}

func (r *TodoListRepo) UpdateListData(listId int, userId int, updatedData *todo.UpdateListData) error {
	setVals := make([]string, 0)
	args := make([]interface{}, 0)
	argPosition := 1

	if updatedData.Title != nil {
		setVals = append(setVals, fmt.Sprintf("title=$%d", argPosition))
		args = append(args, *updatedData.Title)
		argPosition++
	}

	if updatedData.Description != nil {
		setVals = append(setVals, fmt.Sprintf("description=$%d", argPosition))
		args = append(args, *updatedData.Description)
		argPosition++
	}

	setPart := strings.Join(setVals, ", ")

	args = append(args, listId, userId)

	query := fmt.Sprintf(
		"UPDATE %s l SET %s FROM %s ul WHERE l.id=ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		listsTable, setPart, usersListsTable, argPosition, argPosition+1,
	)
	_, err := r.db.Exec(query, args...)
	return err
}
