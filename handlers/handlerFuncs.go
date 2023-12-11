package handlers

import (
	"encoding/json"
	"fmt"
	"io"
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

func HandleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		payload := struct {
			Users []*customTypes.User `json:"users"`
		}{}

		payload.Users = users

		b, err := json.Marshal(payload)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, string(b))
		}
	case http.MethodPost:
		w.WriteHeader(http.StatusOK)
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		newUser := customTypes.User{}

		err = json.Unmarshal(b, &newUser)
		if err != nil {
			http.Error(w, "bad format", http.StatusBadRequest)
			return
		}

		users = append(users, &newUser)

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	default:
		http.Error(w, "wrong method", http.StatusMethodNotAllowed)
	}
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
