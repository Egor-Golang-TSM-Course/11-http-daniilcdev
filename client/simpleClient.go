package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func CallServer() {
	segments := []string{"/", "/time", "/random"}

	for _, path := range segments {
		url := "http://127.0.0.1:8080" + path
		fmt.Println("client: calling", url)
		r, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		defer r.Body.Close()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("client: response", string(body))
	}

}
func main() {
	CallServer()
}
