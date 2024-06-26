package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

type APIServer struct {
	addr    string
	storage *Storage // DB connection/orm

	// http client
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		storage: &Storage{
			db: db,
		},
	}
}

func (s *APIServer) Run() {
	//s.migrate()

	router := mux.NewRouter()

	router.HandleFunc("/hello", s.handleGreet)
	router.HandleFunc("/post/{id}", s.handleGetPost)
	router.HandleFunc("/post/create", s.handleCreatePost)

	fmt.Printf("Server starting on address %s", s.addr)
	http.ListenAndServe(s.addr, router)
}

//func (s *APIServer) migrate() {
//	s.db.AutoMigrate(&Post{})
//}
