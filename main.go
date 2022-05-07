package main

import (
	"fmt"
	"log"
	"os"

	"github.com/agchavez/tongo-backend/db"
	"github.com/agchavez/tongo-backend/handlers"
	"github.com/joho/godotenv"
)

func main() {

	// Cargar variables de entorno del archivo .env
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secretKey := os.Getenv("DB_HOST")
	fmt.Println(secretKey)
	// Conexión a la base de datos
	if db.CheckConnection() == 0 {
		log.Fatal("Error connecting to database")
		return
	}
	handlers.Handlers()
}
