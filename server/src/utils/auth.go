package utils

import (
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomString(length int) string {
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(bytes)
}

func EncryptPassword(password string) (string, error) {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPasswordBytes), err
}

func VerifySamePassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword))
	if err != nil {
		return false
	}
	return true
}
