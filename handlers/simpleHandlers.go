package handlers

import (
	"fmt"
	"lesson11/customTypes"
	"net/http"
	"time"
)

var users []*customTypes.User = make([]*customTypes.User, 0, 4)

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "hello!")
	} else {
		HandleNotFound(w, r)
	}
}

func HandleTime(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, time.Now().Format(time.RFC1123))
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
