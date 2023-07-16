package handler

import "github.com/gin-gonic/gin"

type TodoHandler struct {
}

func (h *TodoHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		list := api.Group("/list")
		{
			list.GET("/", h.getAllLists)
			list.POST("/", h.createList)
			list.GET("/:list_id", h.getList)
			list.PUT("/:list_id", h.editList)
			list.DELETE("/:list_id", h.deleteList)

			item := api.Group(":list_id/item")
			{
				item.GET("/", h.getAllItems)
				item.POST("/", h.createItem)
				item.GET("/:item_id", h.getItem)
				item.PUT("/:item_id", h.editItem)
				item.DELETE("/:item_id", h.deleteItem)
			}
		}
	}

	return router
}
