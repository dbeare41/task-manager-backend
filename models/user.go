package models

import (
	"errors"
	"my-task-manager/db"
	"my-task-manager/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) SaveUser() error {
	query := "INSERT INTO users(email,password) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedpass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.Email, hashedpass)
	if err != nil {
		return err
	}

	return err
}

func (u *User) VerifyUser() error {
	query := "SELECT id,password FROM users WHERE email =?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.Id, &retrievedPassword)
	if err != nil {
		return err
	}
	passwordvalidated := utils.ValidatePassword(u.Password, retrievedPassword)
	if !passwordvalidated {
		return errors.New("invalid password")
	}
	return nil
}
