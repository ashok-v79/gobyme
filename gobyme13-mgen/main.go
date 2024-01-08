// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akv/me13/database"
	"github.com/akv/me13/userservice"
)

func main() {
	// Initialize a simple in-memory database (replace with real DB initialization)
	db := database.NewInMemoryDB()

	// Create an instance of UserService with the database.
	userService := userservice.NewUserService(db)

	// Setup HTTP server and handlers.
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("id")
		userName, err := userService.GetUserName(userID)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "User Name: %s", userName)
	})

	// Start the HTTP server.
	log.Println("Starting server on :8088")
	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
