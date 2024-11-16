package helpers

import (
	"log"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		return
	}
}

func CheckValidEmail(mail string) bool {
	res, err := emailverifier.NewVerifier().Verify(mail)
	if err != nil {
		log.Println(err)
	}
	return res.Syntax.Valid
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println(err)
	}
	return string(bytes)
}

func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println("Invalid password:", err)
		return false
	}
	return true
}

func SliceContains[T comparable](slice []T, value T) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
