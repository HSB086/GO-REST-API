package models

import (
	"errors"

	"haseeb.khan/event-booking/database"
	"haseeb.khan/event-booking/utils"
)

type User struct {
	ID       int64
	Email    string
	Password string
}

func (user User) Save() error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	var userid int64
	err = database.DB.QueryRow(query, user.Email, hashedPassword).Scan(&userid)
	if err != nil {
		return err
	}

	return err
}

func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = $1"
	row := database.DB.QueryRow(query, user.Email)
	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)
	if err != nil {
		return err
	}

	isPasswordValid := utils.CheckHashedPassword(user.Password, retrievedPassword)
	if !isPasswordValid {
		return errors.New("invalid credentials")
	}

	return nil
}
