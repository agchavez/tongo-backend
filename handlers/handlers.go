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
	router.HandleFunc("/user/", middlewares.CheckDB(routers.GetUser)).Methods("GET")
	router.HandleFunc("/user/", middlewares.CheckDB(routers.CreateUser)).Methods("POST")
	router.HandleFunc("/user/", middlewares.CheckDB(routers.UpdateUser)).Methods("PUT")
	router.HandleFunc("/auth/login/", middlewares.CheckDB(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}

	log.Fatal(http.ListenAndServe(":8000", router))
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
