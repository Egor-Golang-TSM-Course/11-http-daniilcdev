package main

import (
	"fmt"
	"lesson11/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HandleDefault)
	http.HandleFunc("/time", handlers.HandleTime)
	http.HandleFunc("/user", handlers.HandleUser)

	fmt.Println("Starting server on 8080 port")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
