package Authentification

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	err = bcrypt.CompareHashAndPassword(bytes, []byte(password))
	if err != nil {
		return "", fmt.Errorf("failed comparison between hash and password: %w", err)
	}
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) (bool, error) {
	hashedPassword, err := GenerateHashPassword(password)
	if err != nil {
		return false, err
	}
	return hashedPassword == hash, err
}
