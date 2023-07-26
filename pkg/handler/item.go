package handler

import (
	"github.com/gin-gonic/gin"
	todo "github.com/rusrom/yt-todo"
	"net/http"
	"strconv"
)

func (h *TodoHandler) getAllItems(c *gin.Context) {

}

func (h *TodoHandler) createItem(ctx *gin.Context) {
	userId, err := getAuthUserId(ctx)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(ctx.Param("list_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid list_id url path param")
		return
	}

	var newItem todo.ItemTodo
	err = ctx.BindJSON(&newItem)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Provide correct json in body")
		return
	}

	id, err := h.services.TodoItemProcessing.CreateNewItem(newItem, listId, userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "todo list doesnt exists or you are not an owner of this todo list")
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *TodoHandler) getItem(c *gin.Context) {

}

func (h *TodoHandler) editItem(c *gin.Context) {

}

func (h *TodoHandler) deleteItem(c *gin.Context) {

}
