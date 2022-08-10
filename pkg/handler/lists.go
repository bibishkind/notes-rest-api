package handler

import (
	entity "github.com/bibishkin/bi-notes-rest-api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createList(c *gin.Context) {
	var list entity.List

	if err := c.BindJSON(&list); err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	userId := c.GetInt("userId")

	listId, err := h.service.CreateList(userId, list)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": listId,
	})
}

func (h *Handler) getLists(c *gin.Context) {
	limit, offset, err := parseLimitAndOffset(c.Query("limit"), c.Query("offset"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	userId := c.GetInt("userId")

	lists, err := h.service.GetLists(userId, limit, offset)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"lists": lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("list_id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	userId := c.GetInt("userId")

	list, err := h.service.GetListById(userId, listId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"list": list,
	})
}

func (h *Handler) updateList(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("list_id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	var list entity.List

	if err := c.BindJSON(&list); err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	userId := c.GetInt("userId")

	err = h.service.UpdateList(userId, listId, list)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) deleteList(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("list_id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	userId := c.GetInt("userId")

	err = h.service.DeleteList(userId, listId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
