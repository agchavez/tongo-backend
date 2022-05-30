package routers

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/agchavez/tongo-backend/db"
	"github.com/agchavez/tongo-backend/models"
	"github.com/agchavez/tongo-backend/jwt"
)

// Router login funtion
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var m models.User

	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		http.Error(w, "Error se espertaba un json "+err.Error(), 400)

	}

	//Validations data
	var listErrors ListError

	if m.Email == "" {
		listErrors.Errors = append(listErrors.Errors, "El email es obligatorio")
	} else {
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if re.FindAllStringSubmatch(m.Email, -1) == nil {
			listErrors.Errors = append(listErrors.Errors, "El email no es valido")
		}
	}
	if len(m.Password) < 5 {
		listErrors.Errors = append(listErrors.Errors, "El password debe tener al menos 5 caracteres")
	}

	if len(listErrors.Errors) > 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(listErrors)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	 isExist, user := db.LoginValidateDB(m.Email, m.Password)

	 if !isExist {
		 http.Error(w, "Usuario y/o contraseña invalido", 400)
		 return 
	 }
	 jwtKey, err := jwt.GenerateJWT(user)
	 if err != nil {
		http.Error(w, "Error al generar el jwt"+err.Error(), 400)
		return 
	}
	resp := models.RespLoginModel{
		Name: user.FirstName + " "  + user.LastName,
		Email: user.Email,
		Token: jwtKey,
	}
	
	 w.Header().Set("Content-Type", "application/json")
	 w.WriteHeader(http.StatusBadRequest)
	 json.NewEncoder(w).Encode(resp)

}
