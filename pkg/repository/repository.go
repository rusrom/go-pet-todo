package repository

import (
	"github.com/jmoiron/sqlx"
	todo "github.com/rusrom/yt-todo"
)

type UserAuthorization interface {
	CreateUser(u todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
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
	return &TodoRepository{
		UserAuthorization: NewAuthRepo(db),
		//TodoListProcessing: nil,
		//TodoItemProcessing: nil,
	}
}
