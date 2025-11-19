package cryptox

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func IsHashed(str string) bool {
	if len(str) == 60 && (strings.HasPrefix(str, "$2a$") || strings.HasPrefix(str, "$2b$") || strings.HasPrefix(str, "$2y$")) {
		return true
	}
	return false
}

func HashMake(str string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hashed)
}

func HashCheck(plain string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
