package routers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"github.com/agchavez/tongo-backend/models"
	"github.com/agchavez/tongo-backend/db"

)
func UserRouter(w http.ResponseWriter, r *http.Request) {
	type ErrorDescription struct {
		description string 
	}
	type LocalErrors struct {
		key ErrorDescription
	}
	if r.Method == "POST" {
		var user models.User

		var error []LocalErrors
		// El body se lee y se destruye el buffer
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
			return
		}

		// Validaciones de datos
		if len(user.Email) == 0 {
			error.append({
				description: "Error al guardar el archivo"
			})
			http.Error(w, "El email es obligatorio", 400)
			return
		}else {
			re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
			if re.FindStringSubmatch(user.Email) == nil {
				http.Error(w, "El correo electronico no es valido", 400)
				return	
			}
		}

		if len(user.Password) == 0 {
			http.Error(w, "El password es obligatorio", 400)
			return
		}

		found, _ := db.FindUserByEmail(user.Email)
		if found {
			http.Error(w, "Ya existe un usuario con el correo "+user.Email, 400)
			return
		}
		_, status := db.SaveNewUser(user)
		if status {
			http.Error(w, "Error al guardar el usuario", 400)
			return
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	} else if r.Method == "GET" {
		//params limit and offset
		limit := r.URL.Query().Get("limit")
		offset := r.URL.Query().Get("offset")
		users := db.GetUsers(limit, offset)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)

		w.WriteHeader(http.StatusOK)
	}
}
