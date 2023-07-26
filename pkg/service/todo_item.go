package service

import (
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
		// list does not exist or not belong to user
		return 0, err
	}
	return s.repoItem.CreateNewItem(item, listId)
}
