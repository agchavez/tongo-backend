package helpers

import (
	"regexp"

	"github.com/agchavez/tongo-backend/models"
)

type ListError struct {
	Errors []string `json:"errors"`
}

func ValidationUser(user models.User) (bool, ListError) {
	var listErrors ListError
	var errors []string
	var isValid bool

	if len(user.FirstName) < 3 {
		errors = append(errors, "El nombre debe tener al menos 3 caracteres")
	}
	if len(user.LastName) < 3 {
		errors = append(errors, "El apellido debe tener al menos 3 caracteres")
	}
	if user.Email == "" {
		listErrors.Errors = append(listErrors.Errors, "El email es obligatorio")
	} else {
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if re.FindAllStringSubmatch(user.Email, -1) == nil {
			listErrors.Errors = append(listErrors.Errors, "El email no es valido")
		}
	}
	if len(user.Password) < 5 {
		listErrors.Errors = append(listErrors.Errors, "El password debe tener al menos 5 caracteres")
	}

	if len(errors) > 0 {
		listErrors.Errors = append(listErrors.Errors, errors...)
		isValid = false
	} else {
		isValid = true
	}
	return isValid, listErrors

}
