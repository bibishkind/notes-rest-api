package handler

import (
	"errors"
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
	limit := c.Query("limit")
	offset := c.Query("offset")

	var limitInt int

	if limit == "" {
		limitInt = -1
	} else {
		var err error
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			errorResponse(c, http.StatusBadRequest, errors.New("bad query"))
			return
		}
	}

	var offsetInt int

	if offset == "" {
		offsetInt = 0
	} else {
		var err error
		offsetInt, err = strconv.Atoi(offset)
		if err != nil {
			errorResponse(c, http.StatusBadRequest, errors.New("bad query"))
			return
		}
	}

	userId := c.GetInt("userId")

	lists, err := h.service.GetLists(userId, limitInt, offsetInt)
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
	}

	userId := c.GetInt("userId")

	err = h.service.DeleteList(userId, listId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
