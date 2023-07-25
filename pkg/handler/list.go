package handler

import (
	"github.com/gin-gonic/gin"
	todo "github.com/rusrom/yt-todo"

	//todo "github.com/rusrom/yt-todo"
	"net/http"
)

type allTodoLists struct {
	Data []todo.ListTodo
}

func (h *TodoHandler) getAllLists(c *gin.Context) {
	userId, err := getAuthUserId(c)
	if err != nil {
		return
	}

	allUserLists, err := h.services.GetAllUserLists(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, allTodoLists{
		Data: allUserLists,
	})
}

func (h *TodoHandler) createList(c *gin.Context) {
	userId, err := getAuthUserId(c)
	if err != nil {
		return
	}

	var newTodo todo.ListTodo
	err = c.BindJSON(&newTodo)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Provide correct json in body")
		return
	}

	id, err := h.services.CreateNewList(newTodo, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *TodoHandler) getList(c *gin.Context) {

}

func (h *TodoHandler) editList(c *gin.Context) {

}

func (h *TodoHandler) deleteList(c *gin.Context) {

}
