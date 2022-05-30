package jwt

import (
	"time"

	"github.com/agchavez/tongo-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// Function to generate a jwt from a user model
func GenerateJWT(m models.User) (string, error) {
	myKey := []byte("AANMNIO,jncxv._$%&//dfgd")
	payload := jwt.MapClaims{
		"email": m.Email,
		"name":  m.FirstName,
		"id":    m.ID.Hex(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil

}
