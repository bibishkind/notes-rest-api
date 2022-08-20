package handler

import (
	entity "github.com/bibishkind/notes-rest-api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Creation of a list
// @Security ApiKeyAuth
// @Tags lists
// @Description creates a list
// @Accept json
// @Produce json
// @Param list body entity.List true "list"
// @Success 201 {integer} int
// @Failure 400 {object} errorMessage
// @Failure 401 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Router /api/lists [post]
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

// @Summary Getting lists
// @Security ApiKeyAuth
// @Tags lists
// @Description gets lists
// @Accept json
// @Produce json
// @Param limit query string int "sets the limit"
// @Param offset query string int "sets the offset"
// @Success 200 {object} []entity.List
// @Failure 400 {object} errorMessage
// @Failure 401 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Router /api/lists [get]
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

// @Summary Getting list by id
// @Security ApiKeyAuth
// @Tags lists
// @Description gets list by id
// @Accept json
// @Produce json
// @Success 200 {object} entity.List
// @Failure 400 {object} errorMessage
// @Failure 401 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Router /api/lists/{list_id} [get]
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

// @Summary Updating list
// @Security ApiKeyAuth
// @Tags lists
// @Description updates a list
// @Accept json
// @Produce json
// @Success 204
// @Failure 400 {object} errorMessage
// @Failure 401 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Router /api/lists/{list_id} [put]
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

// @Summary Deletion of a list
// @Security ApiKeyAuth
// @Tags lists
// @Description deletes a list
// @Accept json
// @Produce json
// @Success 204
// @Failure 400 {object} errorMessage
// @Failure 401 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Router /api/lists/{list_id} [delete]
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
