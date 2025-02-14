package randstr

import (
	"math/rand"
	"time"
)

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var alphanumeric = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func Chars(n int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]rune, n)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(b)
}

func Alphanumeric(n int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]rune, n)
	for i := range b {
		b[i] = alphanumeric[rand.Intn(len(alphanumeric))]
	}

	return string(b)
}

func Date() string {
	minDate := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	maxDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	delta := maxDate - minDate

	sec := rand.Int63n(delta) + minDate
	date := time.Unix(sec, 0)

	return date.Format("2006-01-02")
}
