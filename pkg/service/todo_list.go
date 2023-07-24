package service

import (
	todo "github.com/rusrom/yt-todo"
	"github.com/rusrom/yt-todo/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoListProcessing
}

func NewTodoListService(repo repository.TodoListProcessing) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) CreateNewList(l todo.ListTodo, userId int) (int, error) {
	id, err := s.repo.CreateNewList(l, userId)
	if err != nil {
		return 0, err
	}
	return id, nil
}
