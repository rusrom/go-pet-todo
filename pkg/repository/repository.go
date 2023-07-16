package repository

type UserAuthorization interface {
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

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}
