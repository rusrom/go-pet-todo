package handler

import (
	"github.com/gin-gonic/gin"
	todo "github.com/rusrom/yt-todo"
	"strconv"

	//todo "github.com/rusrom/yt-todo"
	"net/http"
)

type allTodoLists struct {
	Data []todo.ListTodo
}

// @Summary Get All Lists
// @Security ApiKeyAuth
// @Tags lists
// @Description get all lists
// @ID get-all-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} allTodoLists
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/list [get]
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

// @Summary Create todo list
// @Security ApiKeyAuth
// @Tags lists
// @Description create todo list
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body todo.ListTodo true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/list [post]
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

// @Summary Get List Detail
// @Security ApiKeyAuth
// @Tags lists
// @Description get list by id
// @ID get-list-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} todo.ListTodo
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/list/:list_id [get]
func (h *TodoHandler) getList(c *gin.Context) {
	userId, err := getAuthUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("list_id"))
	if err != nil {
		//newErrorResponse(c, http.StatusBadRequest, err.Error())
		newErrorResponse(c, http.StatusBadRequest, "invalid list_id url path param")
		return
	}

	listDetail, err := h.services.GetListDetail(listId, userId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, listDetail)
}

func (h *TodoHandler) editList(c *gin.Context) {
	userId, err := getAuthUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("list_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list_id url path param")
		return
	}

	var updatedData todo.UpdateListData
	if err = c.BindJSON(&updatedData); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "provide correct json in body")
		return
	}

	err = h.services.UpdateListData(listId, userId, &updatedData)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "updated",
	})
}

func (h *TodoHandler) deleteList(c *gin.Context) {
	userId, err := getAuthUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("list_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list_id url path param")
		return
	}

	err = h.services.DeleteList(listId, userId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "deleted",
	})
}
