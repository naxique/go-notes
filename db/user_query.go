package db

import (
	"notes/models"

	_ "github.com/lib/pq"
)

func (s *Database) CreateUser(user *models.UserRequest) error {
	_, err := s.db.Exec(`
		insert into users (username, password)	
		values ($1, $2)
	`, user.Username, user.Password)

	if err != nil {
		return err
	}

	return nil
}

func (s *Database) GetUserByUsername(username string) (*models.User, error) {
	rows, err := s.db.Query(`
		select * from users where username=$1
	`, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := &models.User{}
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Database) GetUser(id int) (*models.User, error) {
	rows, err := s.db.Query(`
		select * from users where id=$1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := &models.User{}
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Database) DeleteUser(userId int) error {
	_, err := s.db.Exec(`
		delete from users where id=$1
	`, userId)

	if err != nil {
		return err
	}

	return nil
}
