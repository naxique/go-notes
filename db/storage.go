package db

import (
	"database/sql"
	"log"
	"notes/models"

	"github.com/gofor-little/env"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateUser(user *models.UserRequest) error
	GetUser(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)

	CreateNote(userId int, note models.Note) error
	GetAllNotes(userId int) ([]*models.Note, error)
	GetNote(noteId int) (*models.Note, error)
	EditNote(noteId int, note *models.NoteEditRequest) error
	DeleteNote(noteId int) error
}

type Database struct {
	db *sql.DB
}

func (s *Database) InitStorage() {
	connStr, err := env.MustGet("POSTGRES")
	if err != nil {
		log.Fatalln(err)
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	s.db = db

	createUsersTableQuery := `
		create table if not exists users(
			id serial primary key,
			username varchar(50) not null unique,
			password varchar(50) not null
		) 
	`
	_, err = s.db.Exec(createUsersTableQuery)
	if err != nil {
		log.Fatalln(err)
	}

	createNotesTableQuery := `
		create table if not exists notes(
			id serial primary key,
			user_id integer references users(id) not null,
			name varchar(100) not null,
			body text,
			created_at timestamp,
			edited_at timestamp
		) 
	`
	_, err = s.db.Exec(createNotesTableQuery)
	if err != nil {
		log.Fatalln(err)
	}
}
