package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
)

// SetupLogger initializes and returns a logger with standard flags.
func SetupLogger(logFilePath string) *log.Logger {
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	return log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
}

// GenerateRandomString generates a random string with a length between min and max characters.
// It uses a predefined character set and returns an error if the random generation fails.
func GenerateRandomString(min, max int) (string, error) {
	if min > max {
		return "", fmt.Errorf("min length cannot be greater than max length")
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	length, err := randomInt(min, max)
	if err != nil {
		return "", fmt.Errorf("failed to generate random length: %w", err)
	}

	return randomString(length, charset)
}

// randomInt generates a random integer between min and max (inclusive).
func randomInt(min, max int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()) + min, nil
}

// randomString generates a random string of the specified length using the given character set.
func randomString(length int, charset string) (string, error) {
	result := make([]byte, length)
	for i := range result {
		charIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random character index: %w", err)
		}
		result[i] = charset[charIndex.Int64()]
	}

	return string(result), nil
}
