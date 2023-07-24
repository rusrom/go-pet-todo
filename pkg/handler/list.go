package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *TodoHandler) getAllLists(c *gin.Context) {

}

func (h *TodoHandler) createList(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *TodoHandler) getList(c *gin.Context) {

}

func (h *TodoHandler) editList(c *gin.Context) {

}

func (h *TodoHandler) deleteList(c *gin.Context) {

}
