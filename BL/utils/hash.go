package utils

import "golang.org/x/crypto/bcrypt"

// BcryptHash encrypts the given password using bcrypt
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}
