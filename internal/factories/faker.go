package factories

import (
	"time"

	"golang.org/x/exp/rand"
)

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var alphanumeric = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandChars(n int) string {
	rand.Seed(uint64(time.Now().UnixNano()))

	b := make([]rune, n)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(b)
}

func RandAlphanumeric(n int) string {
	rand.Seed(uint64(time.Now().UnixNano()))

	b := make([]rune, n)
	for i := range b {
		b[i] = alphanumeric[rand.Intn(len(alphanumeric))]
	}

	return string(b)
}
