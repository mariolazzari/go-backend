package main

import (
	"net/http"
)

type api struct {
	addr string
}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ciao Mario"))
}

func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {

}

// func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	// w.Write([]byte("Ciao Mario!"))
// 	switch r.Method {
// 	case http.MethodGet:
// 		switch r.URL.Path {
// 		case "/":
// 			w.Write([]byte("index page"))
// 			return
// 		case "/users":
// 			w.Write([]byte("users page"))
// 			return
// 		default:
// 			w.Write([]byte("404 page"))
// 			return
// 		}
// 	}
// }

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

	server.ListenAndServe()
}
