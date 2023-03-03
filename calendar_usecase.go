package usecase

import (
	"time"

	"example.com/m/dto"
	"example.com/m/entity"
	"example.com/m/repository"
)

type CalendarUsecase interface {
	InputNotes(data dto.InputNotes) error
	FetchNotes(data dto.NoteRequest) (*[]entity.Notes, error)
	FetchNotesDay(data dto.NoteRequestDay) (*[]entity.Notes, error)
	DeleteNote(data dto.NoteDeleteRequest) error
	UpdateNotes(data dto.UpdateNotesRequest) error
}

type calendarUsecaseImpl struct {
	calendarRepo repository.CalendarRepo
}

type CalendarUsecaseConfig struct {
	CalendarRepo repository.CalendarRepo
}

func NewCalendarUsecase(h CalendarUsecaseConfig) CalendarUsecase {
	return &calendarUsecaseImpl{
		calendarRepo: h.CalendarRepo,
	}
}

func (c *calendarUsecaseImpl) InputNotes(data dto.InputNotes) error {

	dateString := data.NoteDate

	date, err := time.Parse("2006-01-02", dateString)

	if err != nil {
		return err
	}

	notesInput := &entity.Notes{
		Title:    data.Title,
		Note:     data.Note,
		Priority: data.Priority,
		NoteDate: date,
	}

	c.calendarRepo.InputNotes(notesInput)
	err = c.calendarRepo.InputNotes(notesInput)

	if err != nil {
		return err
	}

	return nil
}

func (c *calendarUsecaseImpl) UpdateNotes(data dto.UpdateNotesRequest) error {

	notesInput := &entity.Notes{
		NoteId:   data.NoteID,
		Title:    data.Title,
		Note:     data.Note,
		Priority: data.Priority,
	}

	err := c.calendarRepo.UpdateNote(notesInput)

	if err != nil {
		return err
	}

	return nil
}

func (c *calendarUsecaseImpl) FetchNotes(data dto.NoteRequest) (*[]entity.Notes, error) {

	notes, err := c.calendarRepo.FetchNotes(data)

	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (c *calendarUsecaseImpl) FetchNotesDay(data dto.NoteRequestDay) (*[]entity.Notes, error) {

	notes, err := c.calendarRepo.FetchNotesDay(data)

	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (c *calendarUsecaseImpl) DeleteNote(data dto.NoteDeleteRequest) error {

	err := c.calendarRepo.DeleteNote(data)

	if err != nil {
		return err
	}

	return nil
}
