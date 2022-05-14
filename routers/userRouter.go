package routers

import (
	"encoding/json"
	"net/http"

	"github.com/agchavez/tongo-backend/db"
	"github.com/agchavez/tongo-backend/helpers"
	"github.com/agchavez/tongo-backend/models"
)

type ListError struct {
	Errors []string `json:"errors"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")
	users := db.GetUsers(limit, offset)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var listErrors ListError
	var user models.User
	// Obtener el body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error se espertaba un json"+err.Error(), 400)
		return
	}

	isValid, listError := helpers.ValidationUser(user)
	if !isValid {
		listErrors.Errors = append(listErrors.Errors, listError.Errors...)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(listErrors)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Crear el usuario
	exist, _ := db.FindUserByEmail(user.Email)
	if exist {
		http.Error(w, "El usuario ya existe", 400)
		return
	}
	user = db.CreateUser(user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusCreated)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var listErrors ListError
	var user models.User

	//Verificar si el usuario existe
	isExist, _ := db.FindUserById(id)
	if !isExist {
		http.Error(w, "El usuario no existe", 400)
		return
	}
	// Obtener el body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error se espertaba un json"+err.Error(), 400)
		return
	}

	isValid, listError := helpers.ValidationUser(user)
	if !isValid {
		listErrors.Errors = append(listErrors.Errors, listError.Errors...)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(listErrors)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Actualizar el usuario
	userDB := db.UpdateUserById(id, user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userDB)
	w.WriteHeader(http.StatusCreated)
}
