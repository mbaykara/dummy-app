package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// User represents a user object
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func main() {
	http.HandleFunc("/", logRequest(getHome))
	http.HandleFunc("/users", logRequest(getUsersHandler))
	http.HandleFunc("/createUsers", logRequest(createUserHandler))
	http.HandleFunc("/deleteUser", logRequest(deleteUserHandler))

	log.Printf("Server is running with version %s", os.Getenv("VERSION"))
	log.Fatal(http.ListenAndServe(":8090", nil))
}

// logRequest is a middleware that logs the API call and then executes the given handler
func logRequest(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("API call: %s %s", r.Method, r.URL.Path)
		handler(w, r)
	}
}

// getHome handles the GET request for the home page
func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to homepage!"))
	fmt.Printf("App version: %s\n", os.Getenv("VERSION"))
}

// getUsersHandler handles the GET request for the /users endpoint
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// createUserHandler handles the POST request for the /createUsers endpoint
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	newUser.ID = len(users) + 1
	users = append(users, newUser)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

// deleteUserHandler handles the DELETE request for the /deleteUser endpoint
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	if username == "" {
		// Remove the last user if no username is provided
		if len(users) > 0 {
			users = users[:len(users)-1]
		}
	} else {
		// Find and remove the user by username
		for i, user := range users {
			if user.Username == username {
				users = append(users[:i], users[i+1:]...)
				break
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}
