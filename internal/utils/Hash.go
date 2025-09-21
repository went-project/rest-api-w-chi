package utils

import "golang.org/x/crypto/bcrypt"

func Hash(text string) string {
	cost := 12
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(text), cost)
	if err != nil {
		panic(err)
	}
	return string(hashedBytes)
}
