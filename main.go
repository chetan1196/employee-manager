package main

import (
	"employee-manager/store"
	"fmt"
	"log"
	"net/http"

	"employee-manager/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new employee store
	store := store.NewInMemoryStore()

	router := mux.NewRouter()

	// Register routes
	handlers.RegisterRoutes(router, store)

	// Start the server
	fmt.Println("Server listening on port 7075...")
	log.Fatal(http.ListenAndServe(":7075", router))
}
