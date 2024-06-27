package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

const SALT string = "XajjQvNhvvRt5GSeFk1xFe"

func HashPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password+SALT), 10)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	return string(hash), nil
}

func CheckPasswordHash(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+SALT))
	return err == nil
}
