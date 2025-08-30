package validator

import (
	"net/mail"
	"unicode"
)

func ValidateEmail(email string) bool {
	if email == "" {
		return false
	}
	
	if len(email) > 5 && len(email) < 30 {
		return false
	}
	
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return true
}

func ValidatePassword(password string) bool {
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	if len(password) < 8 {
		return false
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSpecial
}
