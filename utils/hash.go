package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	bytes, err := bcrypt.GenerateFromPassword(bytePassword, 14)
	hashedPassword := string(bytes)
	return hashedPassword, err
}

func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
