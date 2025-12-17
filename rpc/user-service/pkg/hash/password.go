package hash

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the password using bcrypt.
func HashPassword(password string) (hashed string, err error) {

	if len(password) <= 0 {
		return "", errors.New("password cannot be empty")
	}

	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(h), nil
}

// CheckPassword compares the provided password with the stored hashed password.
func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// isHashed checks if the string is already hashed.
func isHashed(password string) bool {
	return len(password) == 60 // bcrypt hashes are always 60 characters long.
}
