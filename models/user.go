package models

import (
	"errors"

	"example.com/golang_sqlite_api/db"
	"example.com/golang_sqlite_api/utils"
)

type User struct {
	ID       int64  ``
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := `
	INSERT INTO user (email, password)
	VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	user.ID = userId
	return err
}

func (user User) ValidateCredentials() error {
	retrievedPassword, err := user.bindUser()

	if err != nil {
		return errors.New("Credentials invalid")
	}

	isValid := utils.CheckPassword(user.Password, retrievedPassword)

	if !isValid {
		return errors.New("could not validate user")
	}

	return nil
}

func (user *User) bindUser() (string, error) {
	query := `
	SELECT id, password FROM user WHERE email = ?
	`
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)

	return retrievedPassword, err
}
