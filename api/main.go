package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", getUsers)
	fmt.Println("Server running on PORT 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

func getUsers(x http.ResponseWriter, y *http.Request){

	if y.Method != "GET"{
		http.Error(x, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return 
	}

	x.Header().Set("Content=Type", "application/json")
	json.NewEncoder(x).Encode([]User{
		{
		ID:1,
		Name:"José",
	},
		{
		ID:2,
		Name:"Allana",
	},
		{
		ID:3,
		Name:"Fulano",
	},
})
}