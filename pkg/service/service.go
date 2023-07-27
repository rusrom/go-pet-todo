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
	GetListDetail(listId int, userId int) (todo.ListTodo, error)
	DeleteList(listId int, userId int) error
	UpdateListData(listId int, userId int, updatedData *todo.UpdateListData) error
}

type TodoItemProcessing interface {
	CreateNewItem(i todo.ItemTodo, listId int, userId int) (int, error)
	GetListItems(listId int, userId int) ([]todo.ItemTodo, error)
	GetItemDetail(itemId int, userId int) (todo.ItemTodo, error)
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
		TodoItemProcessing: NewTodoItemService(r.TodoItemProcessing, r.TodoListProcessing),
	}
}
