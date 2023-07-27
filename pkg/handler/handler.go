package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rusrom/yt-todo/pkg/service"
)

type TodoHandler struct {
	services *service.TodoService
}

func NewTodoHandler(services *service.TodoService) *TodoHandler {
	return &TodoHandler{services: services}
}

func (h *TodoHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		list := api.Group("/list")
		{
			list.GET("/", h.getAllLists)
			list.POST("/", h.createList)
			list.GET("/:list_id", h.getList)
			list.PUT("/:list_id", h.editList)
			list.DELETE("/:list_id", h.deleteList)

			item := list.Group(":list_id/item")
			{
				item.GET("/", h.getAllItems)
				item.POST("/", h.createItem)
			}
		}
		item := api.Group("/item")
		{
			item.GET("/:item_id", h.getItem)
			item.PUT("/:item_id", h.editItem)
			item.DELETE("/:item_id", h.deleteItem)
		}
	}

	return router
}
