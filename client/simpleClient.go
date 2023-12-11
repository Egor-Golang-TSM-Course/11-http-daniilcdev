package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"lesson11/customTypes"
	"log"
	"net/http"
)

func callServer() {
	segments := []string{"/", "/time", "/random"}
	for _, path := range segments {
		body := get(path)

		fmt.Println("client: response", string(body))
	}

	users := []customTypes.User{
		{Name: "Daniil", Age: 30},
		{Name: "John", Age: 89},
		{Name: "Alex", Age: 14},
	}

	listUsers()

	for i, user := range users {
		fmt.Println("User: #", i+1, user)

		fmt.Println("client: adding new user", user)
		data, err := json.Marshal(user)
		if err != nil {
			log.Fatal(err)
		}

		_ = post("/user", data)
		listUsers()
	}
}

func get(endpoint string) []byte {
	const serverAddr = "http://127.0.0.1:8080"
	fmt.Println("client: [GET] calling endpoint", endpoint)
	r, err := http.Get(serverAddr + endpoint)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body
}

func post(endpoint string, data []byte) []byte {
	const serverAddr = "http://127.0.0.1:8080"
	fmt.Println("client: [POST] calling endpoint", endpoint)
	r, err := http.Post(serverAddr+endpoint, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body
}

func listUsers() {
	body := get("/user")

	payload := struct {
		Users []customTypes.User `json:"users"`
	}{}

	err := json.Unmarshal(body, &payload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("client: all users on server", payload.Users)
}

func main() {
	callServer()
}
