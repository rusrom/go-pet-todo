package service

import (
	"errors"
	todo "github.com/rusrom/yt-todo"
	"github.com/rusrom/yt-todo/pkg/repository"
)

type TodoItemService struct {
	repoItem repository.TodoItemProcessing
	repoList repository.TodoListProcessing
}

func NewTodoItemService(repoItem repository.TodoItemProcessing, repoList repository.TodoListProcessing) *TodoItemService {
	return &TodoItemService{
		repoItem: repoItem,
		repoList: repoList,
	}
}

func (s *TodoItemService) CreateNewItem(item todo.ItemTodo, listId int, userId int) (int, error) {
	_, err := s.repoList.GetListDetail(listId, userId)
	if err != nil {
		return 0, errors.New("todo list doesn't exist or you are not an owner of todo list")
	}
	return s.repoItem.CreateNewItem(item, listId)
}

func (s *TodoItemService) GetListItems(listId int, userId int) ([]todo.ItemTodo, error) {
	_, err := s.repoList.GetListDetail(listId, userId)
	if err != nil {
		return nil, errors.New("todo list doesn't exist or you are not an owner of todo list")
	}
	return s.repoItem.GetListItems(listId, userId)
}

func (s *TodoItemService) GetItemDetail(itemId int, userId int) (todo.ItemTodo, error) {
	return s.repoItem.GetItemDetail(itemId, userId)
}

func (s *TodoItemService) DeleteItem(itemId int, userId int) error {
	return s.repoItem.DeleteItem(itemId, userId)
}

func (s *TodoItemService) UpdateItem(updatedData *todo.UpdateItemData, itemId int, userId int) error {
	return s.repoItem.UpdateItem(updatedData, itemId, userId)
}
