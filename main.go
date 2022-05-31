package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

func main() {
	path := mux.NewRouter().StrictSlash(true)

	path.HandleFunc("/", home).Methods(http.MethodGet)
	path.HandleFunc("/createUser", createUser).Methods(http.MethodPost)
	path.HandleFunc("/users", getUsers).Methods(http.MethodGet)
	fmt.Println("Server is running on port 8090")
	log.Fatal(http.ListenAndServe(":8090", path))

}

type USER struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

type allUsers []USER

var users = allUsers{
	{
		ID:   "1",
		Name: "Lenovo",
		Age:  3,
	},
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to homepage!")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newEvent USER
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter valid data")
	}

	json.Unmarshal(reqBody, &newEvent)
	err = json.Unmarshal(reqBody, &newEvent)
	if err != nil {
		fmt.Fprintf(w, "unvalid data")
		return
	}
	users = append(users, newEvent)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "new user added successfully!\n")

	json.NewEncoder(w).Encode(newEvent)
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	fmt.Println("List all users")
	json.NewEncoder(w).Encode(users)
}
