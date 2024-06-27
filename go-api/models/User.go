package models

import (
	"phpguru.net/go-api/db"
	"phpguru.net/go-api/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) CreateUser() error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	query := `
        INSERT INTO users(email, password)
        VALUES(?, ?)
    `
	stmt, err := db.GetDB().Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	rs, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	id, err := rs.LastInsertId()
	u.Id = id
	return err
}

func (u *User) ValidateCredentials() (bool, error) {
	query := `SELECT id, email, password FROM users where email = ?`
	stmt, err := db.GetDB().Prepare(query)
	if err != nil {
		return false, err
	}
	var retrievePassword string
	row := stmt.QueryRow(u.Email)
	err = row.Scan(&u.Id, &u.Email, &retrievePassword)
	if err != nil {
		return false, err
	}
	// compare
	passwordIsValid := utils.CheckPasswordHash(retrievePassword, u.Password)
	if !passwordIsValid {
		return false, nil
	}
	return true, nil
}

func (u *User) GenerateAccessTokens() (string, error) {
	return utils.GenerateToken(u.Email, u.Id)
}
