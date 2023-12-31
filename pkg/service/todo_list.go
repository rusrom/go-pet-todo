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

func (s *TodoListService) GetAllUserLists(userId int) ([]todo.ListTodo, error) {
	lists, err := s.repo.GetAllUserLists(userId)
	if err != nil {
		return nil, err
	}
	return lists, nil
}

func (s *TodoListService) GetListDetail(listId int, userId int) (todo.ListTodo, error) {
	return s.repo.GetListDetail(listId, userId)
}

func (s *TodoListService) DeleteList(listId int, userId int) error {
	return s.repo.DeleteList(listId, userId)
}

func (s *TodoListService) UpdateListData(listId int, userId int, updatedData *todo.UpdateListData) error {
	if err := updatedData.ValidateFields(); err != nil {
		return err
	}
	return s.repo.UpdateListData(listId, userId, updatedData)
}
