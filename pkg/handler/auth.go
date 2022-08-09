package handler

import (
	entity "github.com/bibishkin/bi-notes-rest-api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var user entity.User

	err := c.BindJSON(&user)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	id, err := h.service.CreateUser(user.Username, user.Password)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var user entity.User

	err := c.BindJSON(&user)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	token, err := h.service.GenerateJWT(user.Username, user.Password)
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"jwt": token,
	})

}
