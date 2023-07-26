package todo

import "errors"

type ListTodo struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id     int
	UserId string
	ListId string
}

type ItemTodo struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type ItemList struct {
	Id     int
	ListId string
	ItemId string
}

type UpdateListData struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (d *UpdateListData) ValidateFields() error {
	if d.Title == nil && d.Description == nil {
		return errors.New("update data consists only from nil values")
	}
	return nil
}
