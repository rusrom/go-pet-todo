package repository

import "github.com/jmoiron/sqlx"

type UserAuthorization interface {
}

type TodoListProcessing interface {
}

type TodoItemProcessing interface {
}

type TodoRepository struct {
	UserAuthorization
	TodoListProcessing
	TodoItemProcessing
}

func NewTodoRepository(db *sqlx.DB) *TodoRepository {
	return &TodoRepository{}
}
