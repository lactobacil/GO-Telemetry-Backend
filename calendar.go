package entity

import "time"

type Notes struct {
	NoteId   int       `json:"note_id" gorm:"primaryKey"`
	Title    string    `json:"title" gorm:"primaryKey"`
	Note     string    `json:"note" gorm:"not null"`
	Priority string    `json:"priority" gorm:"not null"`
	NoteDate time.Time `json:"note_date" gorm:"not null"`
}
