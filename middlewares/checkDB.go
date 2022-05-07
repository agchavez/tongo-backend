package middlewares

import (
	"log"
	"net/http"

	"github.com/agchavez/tongo-backend/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	log.Println("Checking database connection")
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error connecting to database"))
			return
		}
		next.ServeHTTP(w, r)
	}
}
