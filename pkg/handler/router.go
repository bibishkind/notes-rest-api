package handler

import (
	_ "github.com/bibishkin/bi-notes-rest-api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

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

		notes := lists.Group(":list_id/notes")
		{
			notes.POST("/", h.createNote)
			notes.GET("/", h.getNotes)
			notes.GET("/:note_id", h.getNoteById)
			notes.PUT("/:note_id", h.updateNote)
			notes.DELETE("/:note_id", h.deleteNote)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
