package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/agchavez/tongo-backend/middlewares"
	"github.com/agchavez/tongo-backend/routers"
)

// App is a struct to hold the router and the database connection
func Handlers() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/user/", middlewares.CheckDB(routers.UserRouter)).Methods("POST", "GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}

	log.Fatal(http.ListenAndServe(":8000", router))
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}