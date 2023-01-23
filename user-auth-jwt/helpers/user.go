package helpers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash , err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}

func ComparePassword(password, hashedPassword string) error {
	 err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	fmt.Println("error", err)
	 return err
}