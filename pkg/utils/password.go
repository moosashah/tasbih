package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hashPassword)
}

func checkPassword(password, hashedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword)); err != nil {
		return false
	}
	return true
}
