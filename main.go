package main

import (
	"encoding/json"
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

	log.Printf("Server is running with version %s", os.Getenv("VERSION"))
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func logRequest(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("API Method: %s Endpoint: %s", r.Method, r.URL.Path)
		handler(w, r)
	}
}

func getHome(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hello from %s\n", os.Getenv("REGION"))

	jsonData, err := json.Marshal(os.Getenv("REGION"))
	if err != nil {
		http.Error(w, "Failed to retrieve region", http.StatusInternalServerError)
		return
	}
	w.Header()
	w.Write(append([]byte("Hello From "), jsonData...))
	w.Write(append([]byte("\n")))
}

func connectDB() error {
	u := os.Getenv("DB_USERNAME")
	p := os.Getenv("DB_PASSWORD")
	if u == "" && p == "" {
		log.Fatalln("No Database credentials found!")
	} else {
		log.Println("Database connected successfully")
	}
	return nil
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	connectDB()
	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	if len(users) < 1 {
		log.Println("No user exist in the database")
	}
	w.WriteHeader(http.StatusOK)
	log.Printf("Http Status Code: %d\n", http.StatusOK)
	w.Write(jsonData)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User
	connectDB()
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
