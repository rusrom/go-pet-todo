package service

import (
	todo "github.com/rusrom/yt-todo"
	"github.com/rusrom/yt-todo/pkg/repository"
)

type UserAuthorization interface {
	CreateUser(u todo.User) (int, error)
	GenerateJWT(u todo.SignInInput) (string, error)
}

type TodoListProcessing interface {
}

type TodoItemProcessing interface {
}

type TodoService struct {
	UserAuthorization
	TodoListProcessing
	TodoItemProcessing
}

func NewTodoService(r *repository.TodoRepository) *TodoService {
	return &TodoService{
		UserAuthorization: NewAuthService(r.UserAuthorization),
		//TodoListProcessing: nil,
		//TodoItemProcessing: nil,
	}
}
