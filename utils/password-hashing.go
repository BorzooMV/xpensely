package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) ([]byte, error) {
	const bcryptCost = 5
	return bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
}

func ValidatePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
