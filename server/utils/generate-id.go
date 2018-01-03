package utils

import "math/rand"

const (
	size    = 5
	simbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-+"
)

// GenerateID generates a random ID with 5 characters
func GenerateID() string {
	id := make([]byte, size, size)
	for i := range id {
		id[i] = simbols[rand.Intn(len(simbols))]
	}

	return string(id)
}
