package usecase

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"regexp"
)

// HashActivationLink : ...
func HashActivationLink(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 14)
	return string(bytes), err
}

var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
// RandomString : Function to generate Randam String
func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func IsEmpty(s string) bool {
	if s == "" {
		return true
	}
	return false
}

func IsEmptyID(i int) bool {
	if i == 0 {
		return true
	}
	return false
}


func ValidateEmail(s string) bool {
	if s == "" {
		return true
	}
	emailRegexp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegexp.MatchString(s) {
		return false
	}
	return true
}


