package handler

import (
	"github.com/gin-gonic/gin"
	todo "github.com/rusrom/yt-todo"
	"net/http"
	"strconv"
)

type listAllItems struct {
	Data []todo.ItemTodo `json:"data"`
}

func (h *TodoHandler) getAllItems(ctx *gin.Context) {
	userId, err := getAuthUserId(ctx)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(ctx.Param("list_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid list_id url path param")
		return
	}

	listItems, err := h.services.GetListItems(listId, userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, listAllItems{
		Data: listItems,
	})
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
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *TodoHandler) getItem(ctx *gin.Context) {
	userId, err := getAuthUserId(ctx)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(ctx.Param("item_id"))
	if err != nil {
		//newErrorResponse(c, http.StatusBadRequest, err.Error())
		newErrorResponse(ctx, http.StatusBadRequest, "invalid item_id url path param")
		return
	}

	itemDetail, err := h.services.TodoItemProcessing.GetItemDetail(itemId, userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, itemDetail)
}

func (h *TodoHandler) editItem(c *gin.Context) {

}

func (h *TodoHandler) deleteItem(ctx *gin.Context) {
	userId, err := getAuthUserId(ctx)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(ctx.Param("item_id"))
	if err != nil {
		//newErrorResponse(c, http.StatusBadRequest, err.Error())
		newErrorResponse(ctx, http.StatusBadRequest, "invalid item_id url path param")
		return
	}

	err = h.services.TodoItemProcessing.DeleteItem(itemId, userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{Status: "deleted"})
}
