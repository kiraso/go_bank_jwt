package main

import (
	"encoding/json"
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/mux"
)


func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
type apiFunc func(http.ResponseWriter, *http.Request) error 

type ApiError struct{
	Error string
}

type APIServer struct {
	listenAddr string
}
func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}


func makeHttpHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle error
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
			return 
		}
		
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHttpHandler(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHttpHandler(s.handleGetAccount))
	log.Println("Server running on", s.listenAddr)
	http.ListenAndServe(s.listenAddr,router)
	//router.HandleFunc("/account"), s.handleGetAccount)
}
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}

	if r.Method == "DELETE" {
		return s.handleCreateAccount(w, r)
	}
	return fmt.Errorf("Method not allowed")
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	fmt.Println(id)
	//account := NeweAccount("John", "Doe")
	return WriteJSON(w, http.StatusOK, &Account{}) 

}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil 

}


func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil 
}