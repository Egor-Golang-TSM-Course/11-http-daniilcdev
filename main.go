package main

import (
	"fmt"
	"net/http"
	"time"
)

func serveDefault(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "hello!")
	} else {
		serveNotFound(w, r)
	}
}

func serveNotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
func serveTime(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, time.Now().Format(time.RFC1123))
}

func main() {

	http.HandleFunc("/", serveDefault)
	http.HandleFunc("/time", serveTime)

	fmt.Println("Starting server on 8080 port")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
