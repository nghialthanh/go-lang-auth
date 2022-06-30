package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	username string
	email    string
	password string
}

// returns the bcrypt hash from password string
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// compares a bcrypt hashed password with password submit by user
func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Remove space and encode some special characters of data, avoid injection security
func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}
