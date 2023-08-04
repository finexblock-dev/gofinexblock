package utils

import "golang.org/x/crypto/bcrypt"

func CompareHash(hashed, value string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(value))
	return err == nil
}