package db

import (
	"notes/models"
	"time"

	_ "github.com/lib/pq"
)

func (s *Database) CreateNewNote(userId int, note *models.NoteRequest) error {
	_, err := s.db.Exec(`
		insert into notes (user_id, name, body, created_at)	
		values ($1, $2, $3, $4)
	`, userId, note.Name, note.Body, time.Now().UTC())

	if err != nil {
		return err
	}

	return nil
}

func (s *Database) GetAllNotes(userId int) ([]*models.Note, error) {
	rows, err := s.db.Query(`
		select * from notes where user_id=$1
	`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notes := []*models.Note{}
	for rows.Next() {
		note := new(models.Note)
		if err := rows.Scan(
			&note.ID,
			&note.UserID,
			&note.Name,
			&note.Body,
			&note.CreatedAt,
			&note.EditedAt); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}

func (s *Database) GetNote(noteId int) (*models.Note, error) {
	rows, err := s.db.Query(`
		select * from notes where id=$1
	`, noteId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	note := &models.Note{}
	for rows.Next() {
		if err := rows.Scan(
			&note.ID,
			&note.UserID,
			&note.Name,
			&note.Body,
			&note.CreatedAt,
			&note.EditedAt); err != nil {
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return note, nil
}

func (s *Database) EditNote(noteId int, note *models.NoteEditRequest) error {
	_, err := s.db.Exec(`
		update notes set (name, body, edited_at) = ($2, $3, $4)
		where id=$1
	`, noteId, note.Name, note.Body, time.Now().UTC())

	if err != nil {
		return err
	}

	return nil
}

func (s *Database) DeleteNote(noteId int) error {
	_, err := s.db.Exec(`
		delete from notes where id=$1
	`, noteId)

	if err != nil {
		return err
	}

	return nil
}
