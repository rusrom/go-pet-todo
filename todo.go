package todo

type ListTodo struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required""`
	Description string `json:"description"`
}

type UserList struct {
	Id     int
	UserId string
	ListId string
}

type ItemTodo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ItemList struct {
	Id     int
	ListId string
	ItemId string
}
