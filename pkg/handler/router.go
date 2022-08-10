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
			lists.GET("/", h.getLists)
			lists.GET("/:list_id", h.getListById)
			lists.PUT("/:list_id", h.updateList)
			lists.DELETE("/:list_id", h.deleteList)
		}

		items := api.Group(":id/items")
		{
			items.POST("/", h.createNote)
			items.GET("/", h.getNotes)
			items.GET("/:note_id", h.getNoteById)
			items.PUT("/:note_id", h.updateNote)
			items.DELETE("/:note_id", h.deleteNote)
		}
	}

	return router
}
