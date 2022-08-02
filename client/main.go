package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

func main() {
	res, err := http.Get("http://localhost:8080/users")

	if err != nil {
		fmt.Println(err.Error())
		return 
	}

	if res.StatusCode != 200 {
		fmt.Println("Error", res.StatusCode)
		return
	}

	body, err := io.ReadAll(res.Body)

	var response []User
	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("Failed data", err.Error())
		return
	}

	fmt.Println(response)
}