package db

import "golang.org/x/crypto/bcrypt"

func EncripPassword(password string) (string, error) {
	temp := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), temp)
	return string(bytes), err

}
