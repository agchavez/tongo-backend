package db

import (
	"github.com/agchavez/tongo-backend/models"
	"golang.org/x/crypto/bcrypt"
)

// Funtion validate email and password with user login
func LoginValidateDB(email string, password string) (bool, models.User) {

	isexists, user := FindUserByEmail(email)

	if !isexists {
		return false, user
	}

	passwordBytes := []byte(password)
	passwordDb := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordDb, passwordBytes)

	if err != nil {
		return false, user
	}

	return true, user

}
