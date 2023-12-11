package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"lesson11/customTypes"
	"net/http"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		payload := struct {
			Users []*customTypes.User `json:"users"`
		}{}

		payload.Users = users

		b, err := json.Marshal(payload)
		if err != nil {
			http.Error(w, "serialization error", http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, string(b))
		}
	case http.MethodPost:
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

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	default:
		http.Error(w, "wrong method", http.StatusMethodNotAllowed)
	}
}
