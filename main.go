package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func main() {
	http.HandleFunc("/", logRequest(getHome))
	http.HandleFunc("/users", logRequest(getUsersHandler))
	http.HandleFunc("/createUser", logRequest(createUserHandler))
	http.HandleFunc("/deleteUser", logRequest(deleteUserHandler))

	version := os.Getenv("VERSION")
	if version == "" {
		log.Println("VERSION environment variable not set")
		version = "unknown"
	}
	log.Printf("Server is running with version %s", version)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func logRequest(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("API Method: %s Endpoint: %s", r.Method, r.URL.Path)
		handler(w, r)
	}
}

func getHome(w http.ResponseWriter, r *http.Request) {
	region := os.Getenv("REGION")
	if region == "" {
		log.Println("REGION environment variable not set")
		region = "unknown"
	}

	log.Printf("Hello from %s\n", region)

	jsonData, err := json.Marshal(region)
	if err != nil {
		http.Error(w, "Failed to retrieve region", http.StatusInternalServerError)
		return
	}
	w.Write(append([]byte("Hello From "), jsonData...))
	w.Write([]byte("\n"))
}

func connectDB() error {
	u := os.Getenv("DB_USERNAME")
	p := os.Getenv("DB_PASSWORD")
	if u == "" || p == "" {
		return logError("No Database credentials found!")
	}
	log.Println("Database connected successfully")
	return nil
}

func logError(message string) error {
	log.Println(message)
	return fmt.Errorf(message)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := connectDB(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	if len(users) < 1 {
		log.Println("No users found in the database!")
	}
	w.WriteHeader(http.StatusOK)
	log.Printf("Http Status Code: %d\n", http.StatusOK)
	w.Write(jsonData)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User
	if err := connectDB(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	if username == "" {
		if len(users) > 0 {
			users = users[:len(users)-1]
		}
	} else {
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
