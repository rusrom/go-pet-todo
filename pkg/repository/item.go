package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "github.com/rusrom/yt-todo"
	"strings"
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

func (r *TodoItemRepo) GetListItems(listId int, userId int) ([]todo.ItemTodo, error) {
	var listItems []todo.ItemTodo
	query := fmt.Sprintf(
		`SELECT i.id, i.title, i.description, i.done FROM %s i 
    			INNER JOIN %s li ON i.id = li.item_id
    			INNER JOIN %s ul ON li.list_id = ul.list_id
                WHERE li.list_id = $1 AND ul.user_id = $2
                ORDER BY i.id`,
		itemsTable,
		listsItemsTable,
		usersListsTable,
	)
	if err := r.db.Select(&listItems, query, listId, userId); err != nil {
		return nil, err
	}
	return listItems, nil
}

func (r *TodoItemRepo) GetItemDetail(itemId int, userId int) (todo.ItemTodo, error) {
	var listItem todo.ItemTodo
	query := fmt.Sprintf(
		`SELECT i.id, i.title, i.description, i.done FROM %s i
    			INNER JOIN %s li ON i.id = li.item_id
				INNER JOIN %s ul ON li.list_id = ul.list_id
				WHERE i.id = $1 AND ul.user_id = $2`,
		itemsTable,
		listsItemsTable,
		usersListsTable,
	)

	err := r.db.Get(&listItem, query, itemId, userId)
	return listItem, err
}

func (r *TodoItemRepo) DeleteItem(itemId int, userId int) error {
	query := fmt.Sprintf(
		`DELETE FROM %s i
       			USING %s li, %s ul
				WHERE i.id = li.item_id AND li.list_id = ul.list_id AND li.item_id = $1 AND ul.user_id = $2`,
		itemsTable,
		listsItemsTable,
		usersListsTable,
	)
	_, err := r.db.Exec(query, itemId, userId)

	return err
}

func (r *TodoItemRepo) UpdateItem(updatedData *todo.UpdateItemData, itemId int, userId int) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argPosition := 1

	if updatedData.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argPosition))
		args = append(args, *updatedData.Title)
		argPosition++
	}

	if updatedData.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argPosition))
		args = append(args, *updatedData.Description)
		argPosition++
	}

	if updatedData.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argPosition))
		args = append(args, *updatedData.Done)
		argPosition++
	}

	setPart := strings.Join(setValues, ", ")

	args = append(args, itemId, userId)

	query := fmt.Sprintf(
		`UPDATE %s i SET %s
				FROM %s li, %s ul
				WHERE i.id = li.item_id AND li.list_id = ul.list_id AND li.item_id = $%d AND ul.user_id = $%d`,
		itemsTable, setPart, listsItemsTable, usersListsTable, argPosition, argPosition+1,
	)
	_, err := r.db.Exec(query, args...)
	return err
}
