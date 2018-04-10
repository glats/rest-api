package main

import (
	"log"
	"net/http"

	"github.com/glats/rest-api/app/contact"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// our main function
func main() {
	corsObj := handlers.AllowedOrigins([]string{"*"})
	router := mux.NewRouter()
	router.HandleFunc("/contact", contact.RecieveContact).Methods("POST")
	log.Printf("%v", http.ListenAndServe(":8080", handlers.CORS(corsObj)(router)))
}
