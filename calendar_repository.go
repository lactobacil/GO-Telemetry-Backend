package repository

import (
	"database/sql"
	"fmt"

	"example.com/m/dto"
	"example.com/m/entity"
)

type CalendarRepo interface {
	InputNotes(newNote *entity.Notes) error
	FetchNotes(dateRequest dto.NoteRequest) (*[]entity.Notes, error)
	DeleteNote(noteTitle dto.NoteDeleteRequest) error
	FetchNotesDay(dateRequest dto.NoteRequestDay) (*[]entity.Notes, error)
	UpdateNote(noteUpdate *entity.Notes) error
}

type calendarRepoImpl struct {
	db *sql.DB
}

type CalendarRepositoryConfig struct {
	DB *sql.DB
}

func NewCalendarRepository(c CalendarRepositoryConfig) CalendarRepo {
	return &calendarRepoImpl{
		db: c.DB,
	}
}

func (r *calendarRepoImpl) InputNotes(newNote *entity.Notes) error {

	insertNotes := `insert into "calendar_notes"("title", "note", "priority", "note_date") values($1, $2, $3, $4)`
	_, err := r.db.Exec(insertNotes, newNote.Title, newNote.Note, newNote.Priority, newNote.NoteDate)

	if err != nil {
		return err
	}

	return nil
}

func (r *calendarRepoImpl) FetchNotes(dateRequest dto.NoteRequest) (*[]entity.Notes, error) {
	var notes []entity.Notes
	rows, err := r.db.Query(`SELECT * FROM calendar_notes 
	WHERE EXTRACT(MONTH FROM note_date) = $1 
	AND EXTRACT(YEAR FROM note_date) = $2
	ORDER BY priority ASC`, dateRequest.Month, dateRequest.Year)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		note := new(entity.Notes)
		err = rows.Scan(&note.NoteId, &note.Title, &note.Note, &note.Priority, &note.NoteDate)
		notes = append(notes, *note)

		fmt.Println(notes)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	return &notes, nil
}

func (r *calendarRepoImpl) FetchNotesDay(dateRequest dto.NoteRequestDay) (*[]entity.Notes, error) {
	var notes []entity.Notes
	rows, err := r.db.Query(`SELECT * FROM calendar_notes 
	WHERE EXTRACT(MONTH FROM note_date) = $1 
	AND EXTRACT(YEAR FROM note_date) = $2
	AND EXTRACT(DAY FROM note_date) = $3
	ORDER BY note_id DESC`, dateRequest.Month, dateRequest.Year, dateRequest.Day)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		note := new(entity.Notes)
		err = rows.Scan(&note.NoteId, &note.Title, &note.Note, &note.Priority, &note.NoteDate)
		notes = append(notes, *note)

		fmt.Println(notes)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	return &notes, nil
}

func (r *calendarRepoImpl) UpdateNote(noteUpdate *entity.Notes) error {
	_, err := r.db.Query(`UPDATE calendar_notes
						  SET title = $1, 
						  note = $2,
						  priority = $3
	                      WHERE note_id = $4`, noteUpdate.Title, noteUpdate.Note, noteUpdate.Priority, noteUpdate.NoteId)

	if err != nil {
		return err
	}

	return nil
}

func (r *calendarRepoImpl) DeleteNote(noteTitle dto.NoteDeleteRequest) error {

	_, err := r.db.Query(`DELETE FROM calendar_notes WHERE note_id = $1`, noteTitle.NoteID)

	if err != nil {
		return err
	}

	return nil
}
