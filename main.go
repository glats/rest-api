package main

import (
	"log"
	"net/http"

	"github.com/glats/rest-api/app/contact"
    "github.com/rs/cors"
	"github.com/gorilla/mux"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/contact", contact.RecieveContact).Methods("POST")
    handler := cors.Default().Handler(router)
	log.Printf("%v", http.ListenAndServe(":8080", handler))
}
