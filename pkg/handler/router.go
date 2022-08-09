package handler

import "github.com/gin-gonic/gin"

func (h *Handler) GetRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.identifyUser)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
		}

		items := api.Group(":id/items")
		{
			items.POST("/", h.createNote)
			items.GET("/", h.getAllNotes)
			items.GET("/:item_id", h.getNoteById)
			items.PUT("/:item_id", h.updateNote)
			items.DELETE("/:item_id", h.deleteNote)
		}
	}

	return router
}
