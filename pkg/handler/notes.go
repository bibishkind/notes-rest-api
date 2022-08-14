package handler

import (
	entity "github.com/bibishkin/notes-rest-api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Creation of a note
// @Security ApiKeyAuth
// @Tags notes
// @Description creates a note
// @Accept json
// @Produce json
// @Param note body entity.Note true "note"
// @Success 201 {integer} int
// @Failure 400 {object} errorMessage
// @Failure 401 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Router /api/lists/{list_id}/notes [post]
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

// @Summary Getting notes
// @Security ApiKeyAuth
// @Tags notes
// @Description gets notes
// @Accept json
// @Produce json
// @Param limit query string int "sets the limit"
// @Param offset query string int "sets the offset"
// @Success 200 {object} []entity.Note
// @Failure 400 {object} errorMessage
// @Failure 401 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Router /api/lists/{list_id}/notes [get]
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

// @Summary Getting note by id
// @Security ApiKeyAuth
// @Tags notes
// @Description gets note by id
// @Accept json
// @Produce json
// @Success 200 {object} entity.Note
// @Failure 400 {object} errorMessage
// @Failure 401 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Router /api/lists/{list_id}/notes/{note_id} [get]
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

// @Summary Updating note
// @Security ApiKeyAuth
// @Tags notes
// @Description updates a note
// @Accept json
// @Produce json
// @Success 204
// @Failure 400 {object} errorMessage
// @Failure 401 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Router /api/lists/{list_id}/notes/{note_id} [put]
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

// @Summary Deletion of a note
// @Security ApiKeyAuth
// @Tags notes
// @Description deletes a note
// @Accept json
// @Produce json
// @Success 204
// @Failure 400 {object} errorMessage
// @Failure 401 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Router /api/lists/{list_id}/notes/{note_id} [delete]
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
