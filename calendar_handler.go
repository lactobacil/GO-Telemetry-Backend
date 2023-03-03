package handler

import (
	"net/http"

	"example.com/m/dto"
	"example.com/m/usecase"
	"github.com/gin-gonic/gin"
)

type HandlerCalendarImpl struct {
	calendarUsecase usecase.CalendarUsecase
}

type HandlerCalendarConfig struct {
	CalendarUsecase usecase.CalendarUsecase
}

func NewCalendarHandler(c HandlerCalendarConfig) *HandlerCalendarImpl {
	return &HandlerCalendarImpl{
		calendarUsecase: c.CalendarUsecase,
	}
}

func (h *HandlerCalendarImpl) AddNotes(c *gin.Context) {
	var reqInputNotes dto.InputNotes

	if err := c.ShouldBindJSON(&reqInputNotes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
	}

	err := h.calendarUsecase.InputNotes(reqInputNotes)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success added notes",
	})
}

func (h *HandlerCalendarImpl) FetchNotes(c *gin.Context) {
	var reqDates dto.NoteRequest

	if err := c.ShouldBindJSON(&reqDates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
	}

	notesResult, err := h.calendarUsecase.FetchNotes(reqDates)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": notesResult,
	})
}

func (h *HandlerCalendarImpl) FetchNotesDay(c *gin.Context) {
	var reqDates dto.NoteRequestDay

	if err := c.ShouldBindJSON(&reqDates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
	}

	notesResult, err := h.calendarUsecase.FetchNotesDay(reqDates)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": notesResult,
	})
}

func (h *HandlerCalendarImpl) DeleteNotes(c *gin.Context) {
	var reqDeleteNote dto.NoteDeleteRequest

	if err := c.ShouldBindJSON(&reqDeleteNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
	}

	err := h.calendarUsecase.DeleteNote(reqDeleteNote)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

}

func (h *HandlerCalendarImpl) UpdateNotes(c *gin.Context) {
	var reqUpdateNote dto.UpdateNotesRequest

	if err := c.ShouldBindJSON(&reqUpdateNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
	}

	err := h.calendarUsecase.UpdateNotes(reqUpdateNote)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

}
