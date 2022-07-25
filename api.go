package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Order struct {
	Item       string
	Restaurant string
}
type User struct {
	UserID          string
	HasMembership   bool
	PreviousOrders  []Order
	CurrentLocation string
}

var users []User

func main() {
	handleRequests()
}

func handleRequests() {
	newUser := User{
		UserID:          "1234",
		HasMembership:   false,
		PreviousOrders:  []Order{},
		CurrentLocation: "Guntur",
	}
	users = append(users, newUser)

	router := mux.NewRouter()
	
	//Reading
	router.HandleFunc("/getUsers", handleGetUsers)
	router.HandleFunc("/user/{id}", handleGetSingleUser)

	//Creating
	router.HandleFunc("/addUser", handleAddUser).Methods("POST")

	//Updating
	router.HandleFunc("/user/updateMembership/{id}", handleUpdateMembership).Methods("PATCH")

	//Deleting
	router.HandleFunc("/user/delete/{id}", handleDeleteUser).Methods("DELETE")
	http.HandleFunc("/", handleHome)

	log.Fatal(http.ListenAndServe(":3000", router))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Reached the home route!!! "))
	fmt.Println("End point hit home page")
}

func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	w.Write(res)
	fmt.Println("Fetched all users")
}

func handleGetSingleUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]
	fmt.Println(userId)
	for _, v := range users {
		if v.UserID == userId {
			res, err := json.Marshal(v)
			if err != nil {
				panic(err)
			}
			w.Write(res)
		}
	}
	fmt.Println("Fetched a single user")
}

func handleAddUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(reqBody)
	var newUser User
	json.Unmarshal(reqBody, &newUser)
	users = append(users, newUser)
	json.NewEncoder(w).Encode(newUser)
	fmt.Println("Posting a new user")
}

func handleUpdateMembership(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var updatedUser User
	json.Unmarshal(reqBody, &updatedUser)
	for k, v := range users {
		if v.UserID == userId {
			newUsers := append(users[:k], updatedUser)
			newUsers = append(newUsers, users[k+1:]...)

			json.NewEncoder(w).Encode(updatedUser)
		}
	}

	fmt.Println("Updated the membership")
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]
	for k, v := range users {
		if v.UserID == userId {
			users = append(users[:k], users[k+1:]...)
		}
	}
	fmt.Println("Deleted the user")
}
