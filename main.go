package main

import (
	"log"
	"net/http"

	"github.com/kabukky/httpscerts"
)


func main() {
	// Instantiate the mux router
	router := NewRouter()

	/* Build the HTTPS certs. This is currently only a self-cert for development, and
	a real HTTPS cert from a CA will be needed */
    err := httpscerts.Check("cert.pem", "key.pem")
    // If they are not available, generate new ones.
    if err != nil {
        err = httpscerts.Generate("cert.pem", "key.pem", ":8080")
        if err != nil {
            log.Fatal("Error: Couldn't create https certs.")
        }
    }

	log.Fatal(http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", router))
}