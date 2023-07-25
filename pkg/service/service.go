package service

import (
	todo "github.com/rusrom/yt-todo"
	"github.com/rusrom/yt-todo/pkg/repository"
)

type UserAuthorization interface {
	CreateUser(u todo.User) (int, error)
	GenerateJWT(u todo.SignInInput) (string, error)
	ParseJWT(tokenString string) (int, error)
}

type TodoListProcessing interface {
	CreateNewList(l todo.ListTodo, userId int) (int, error)
	GetAllUserLists(userId int) ([]todo.ListTodo, error)
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
		UserAuthorization:  NewAuthService(r.UserAuthorization),
		TodoListProcessing: NewTodoListService(r.TodoListProcessing),
		//TodoItemProcessing: nil,
	}
}
