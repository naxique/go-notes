package models

import "time"

type NoteRequest struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Body   string `json:"body"`
}

type NoteEditRequest struct {
	Name string `json:"name"`
	Body string `json:"body"`
}

type Note struct {
	ID        int        `json:"note_id"`
	UserID    int        `json:"user_id"`
	Name      string     `json:"name"`
	Body      string     `json:"body"`
	CreatedAt time.Time  `json:"created_at"`
	EditedAt  *time.Time `json:"edited_at"`
}

func NewNote(userId int, name, body string) *Note {
	return &Note{
		UserID:    userId,
		Name:      name,
		Body:      body,
		CreatedAt: time.Now().UTC(),
	}
}
