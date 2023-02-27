package storage

import (
	"fmt"
	"log"
	"simplewebserver/internal/app/models"
)

// Instance of User repository (model instance)
type UserRepository struct {
	storage *Storage
}

var (
	tableUser string = "users"
)

// Create User in DB
func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (login, password) VALUSE ($1, $2) RETURNING ID", tableUser)
	if err := ur.storage.db.QueryRow(query, u.Login, u.Password).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

// Find User in DB
func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	users, err := ur.SelectAll()
	var founded bool
	if err != nil {
		return nil, founded, err
	}
	var userFound *models.User
	for _, u := range users {
		if u.Login == login {
			userFound = u
			founded = true
			break
		}
	}
	return userFound, founded, nil
}

// Select all Users in DB
func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableUser)
	rows, err := ur.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Where to read
	users := make([]*models.User, 0)
	for rows.Next() {
		u := models.User{}
		err := rows.Scan(&u.ID, &u.Login, &u.Password)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &u)
	}
	return users, nil
}
