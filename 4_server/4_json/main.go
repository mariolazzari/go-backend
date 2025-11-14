package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type api struct {
	addr string
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var users = []User{}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// encode users slice to json
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}
	err = insertUser(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func insertUser(u User) error {
	if u.FirstName == "" {
		return errors.New("First name is mandatory")
	}
	if u.LastName == "" {
		return errors.New("Last name is mandatory")
	}

	// storage validation
	for _, user := range users {
		if user.FirstName == u.FirstName && user.LastName == u.LastName {
			return errors.New("User already saved")
		}
	}
	users = append(users, u)

	return nil
}

func main() {
	api := &api{addr: ":8080"}

	// init mux
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
