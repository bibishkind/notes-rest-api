package handler

import (
	entity "github.com/bibishkin/bi-notes-rest-api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createNote(c *gin.Context) {

	listId, err := strconv.Atoi(c.Param("list_id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
	}

	var note entity.Note

	if err := c.BindJSON(&note); err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	userId := c.GetInt("userId")

	noteId, err := h.service.CreateNote(userId, listId, note)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": noteId,
	})
}

func (h *Handler) getNotes(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("list_id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
	}

	limit, offset, err := parseLimitAndOffset(c.Query("limit"), c.Query("offset"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	userId := c.GetInt("userId")

	notes, err := h.service.GetNotes(userId, listId, limit, offset)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"notes": notes,
	})
}

func (h *Handler) getNoteById(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("list_id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	noteId, err := strconv.Atoi(c.Param("note_id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	userId := c.GetInt("userId")

	note, err := h.service.GetNoteById(userId, listId, noteId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"note": note,
	})
}

func (h *Handler) updateNote(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("list_id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	noteId, err := strconv.Atoi(c.Param("note_id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	userId := c.GetInt("userId")

	var note entity.Note

	if err := c.BindJSON(&note); err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	err = h.service.UpdateNote(userId, listId, noteId, note)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) deleteNote(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("list_id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	noteId, err := strconv.Atoi(c.Param("note_id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	userId := c.GetInt("userId")

	err = h.service.DeleteNote(userId, listId, noteId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
