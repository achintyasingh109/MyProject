package main

import (
	"fmt"
	"net/http"
)

type server struct {
	addr string
}

type api struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello From Anmol"))

	switch r.Method {
	//case "GET":
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("index Page"))
			return
		case "/users":
			w.Write([]byte("User Page"))
			return
		case "/admin":
			w.Write([]byte("Admin Page"))
			return
		default:
			w.Write([]byte("Error 404"))
			return
		}

	}

}

func (api *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello From Anmol"))
}
func (api *api) createUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created User..."))
}
func (api *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users List..."))
}

func main() {
	fmt.Println("Hello World")
	// s := &server{addr: ":8010"}
	// http.ListenAndServe(s.addr, s)
	api := &api{addr: "8010"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: api,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUserHandler)

	srv.ListenAndServe()

}
