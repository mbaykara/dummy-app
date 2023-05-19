package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// User represents a user object
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func main() {
	http.HandleFunc("/", getHome)
	http.HandleFunc("/users", getUsersHandler)
	http.HandleFunc("/createUsers", createUserHandler)
	http.HandleFunc("/deleteUser", deleteUserHandler)
	log.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
//get home page
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to homepage!")
}
// getUsersHandler handles the GET request for /users endpoint
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

// createUserHandler handles the POST request for /createUsers endpoint
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

// deleteUserHandler handles the DELETE request for /deleteUser endpoint
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
