package handler

import (
	entity "github.com/bibishkin/notes-rest-api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Authorization of a user
// @Tags auth
// @Description authorizes user
// @Accept json
// @Produce json
// @Param user body entity.User true "user"
// @Success 200 {integer} int
// @Failure 400 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Router /auth/sign-up [post]
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

// @Summary Authentication of a user
// @Tags auth
// @Description authenticates user
// @Accept json
// @Produce json
// @Param user body entity.User true "user"
// @Success 200 {string} string
// @Failure 400 {object} errorMessage
// @Failure 401 {object} errorMessage
// @Router /auth/sign-in [post]
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
