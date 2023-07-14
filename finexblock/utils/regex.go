package utils

import "regexp"

func PasswordRegex(password string) bool {
	lowercaseRegex := regexp.MustCompile(`[a-z]`)
	if !lowercaseRegex.MatchString(password) {
		return false
	}

	// At least one uppercase letter
	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	if !uppercaseRegex.MatchString(password) {
		return false
	}

	// At least one number
	numberRegex := regexp.MustCompile(`\d`)
	if !numberRegex.MatchString(password) {
		return false
	}

	// At least one special character
	specialCharRegex := regexp.MustCompile(`[^a-zA-Z0-9]`)
	if !specialCharRegex.MatchString(password) {
		return false
	}

	// All requirements are met
	return true
}