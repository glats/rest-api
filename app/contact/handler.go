package contact

import (
	"encoding/json"
	"log"
	"net/http"
)

// RecieveContact funciton to get contact from a form
func RecieveContact(w http.ResponseWriter, r *http.Request) {
	var contact Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)
	log.Println("pasé por acá")
	json.NewEncoder(w).Encode(contact)
}
