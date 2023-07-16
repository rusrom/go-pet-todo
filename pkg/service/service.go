package service

import "github.com/rusrom/yt-todo/pkg/repository"

type UserAuthorization interface {
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
	return &TodoService{}
}
