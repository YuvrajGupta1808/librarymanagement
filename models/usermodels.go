package models

import (
	"database/sql"
	"yuvraj/config"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func CreateUser(user User) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}

	defer db.Close()
	query := "INSERT INTO users (id, name, email, username) VALUES ($1, $2, $3, $4)"
	_, err = db.Exec(query, user.ID, user.Name, user.Email, user.Username)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByID(id int) (*User, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	query := "SELECT id, name, email, username FROM users WHERE id = $1"
	row := db.QueryRow(query, id)

	var user User
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &user, nil
}

func UpdateUser(id string, user User) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}

	defer db.Close()
	query := "UPDATE users SET name = $1, email = $2, username = $3 WHERE id = $4"
	_, err = db.Exec(query, user.Name, user.Email, user.Username, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(id int) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}

	defer db.Close()
	query := "DELETE FROM users WHERE id = $1"
	_, err = db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func GetAllUsers() ([]User, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT id, name, email, username FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
