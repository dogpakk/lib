package str

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	hashCost = 10
)

func RandomNumericCode(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	return string(result)
}

func RandomAlphaNumericCode(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	return string(result)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
