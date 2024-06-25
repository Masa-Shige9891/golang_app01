package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type APIServer struct {
	addr string
	// DB connection/orm
	// http client
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		handleGreet(w, r)
	})

	fmt.Printf("Server starting on address %s", s.addr)
	http.ListenAndServe(s.addr, router)
}
