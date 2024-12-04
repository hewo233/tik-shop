package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the password using bcrypt.
func HashPassword(password string) (hashed *string, err error) {
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	t := string(h)
	hashed = &t
	return hashed, nil
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
