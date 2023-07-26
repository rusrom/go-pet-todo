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
	CreateNewList(l todo.ListTodo, userId int) (int, error)
	GetAllUserLists(userId int) ([]todo.ListTodo, error)
	GetListDetail(listId int, userId int) (todo.ListTodo, error)
	DeleteList(listId int, userId int) error
	UpdateListData(listId int, userId int, updatedData *todo.UpdateListData) error
}

type TodoItemProcessing interface {
	CreateNewItem(i todo.ItemTodo, listId int) (int, error)
	GetListItems(listId int, userId int) ([]todo.ItemTodo, error)
}

type TodoRepository struct {
	UserAuthorization
	TodoListProcessing
	TodoItemProcessing
}

func NewTodoRepository(db *sqlx.DB) *TodoRepository {
	return &TodoRepository{
		UserAuthorization:  NewAuthRepo(db),
		TodoListProcessing: NewTodoListRepo(db),
		TodoItemProcessing: NewTodoItemRepo(db),
	}
}
