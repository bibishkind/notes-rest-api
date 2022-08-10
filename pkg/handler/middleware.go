package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) identifyUser(c *gin.Context) {
	headers := c.Request.Header
	authHeader := headers["Authorization"]

	if len(authHeader) == 0 {
		errorResponse(c, http.StatusUnauthorized, errors.New("empty auth header"))
		return
	}

	authHeaderSlice := strings.Split(authHeader[0], " ")

	if len(authHeaderSlice) != 2 {
		errorResponse(c, http.StatusUnauthorized, errors.New("invalid authorization header"))
		return
	}

	jwt := authHeaderSlice[1]
	userId, err := h.service.ParseJWT(jwt)
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, err)
		return
	}

	c.Set("userId", userId)
}
