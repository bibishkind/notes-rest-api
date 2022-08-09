package handler

import (
	"github.com/gin-gonic/gin"
)

type errorMessage struct {
	Error string `json:"error"`
}

func errorResponse(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(code, errorMessage{Error: err.Error()})
}
