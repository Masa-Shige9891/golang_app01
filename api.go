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
	fmt.Printf("Server starting on address %s", s.addr)
	http.ListenAndServe(s.addr, router)
}
