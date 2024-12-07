package main

import (
	"encoding/json"
	"fmt"
	"go/token"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	transferReq := new(TransferRequest)
	if err := json.NewDecoder(r.Body).Decode(transferReq); err != nil {
		return err 
	}
	defer r.Body.Close()

	return WriteJSON(w, http.StatusOK, transferReq)
}
func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func createJwt(account *Account) (*jwt.JWT, error)
func withJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calling JWT auth middleware")
		tokenString := r.Header.Get("x-jwt-token")
		_,err := validateJWT(tokenString)
		if err != nil {
			WriteJSON(w, http.StatusForbidden,ApiError{Error: "invalid token "})
			return 
		}
		handlerFunc(w, r)
	}
}

//jwt-go doc
func validateJWT(tokenString string) (*jwt.Token ,error) {
	secret := os.Getenv("JWT_SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}

type apiFunc func(http.ResponseWriter, *http.Request) error 

type ApiError struct{
	Error string `json:"error"`
}

type APIServer struct {
	listenAddr string
	store Storage
}
func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store: store,
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
	router.HandleFunc("/account/{id}", withJWTAuth(makeHttpHandler(s.handleGetAccountById)))
	router.HandleFunc("/transfer/", makeHttpHandler(s.handleTransfer))

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

	// if r.Method == "DELETE" {
	// 	return s.handleCreateAccount(w, r)
	// }
	return fmt.Errorf("Method not allowed")
}



//
func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts,err := s.store.GetAccounts()
	if err != nil {
		return err
	}
	
	return WriteJSON(w,http.StatusOK,accounts)
}

func (s *APIServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		id, err := getId(r)
	if err != nil {
		return err
	}

	account,err := s.store.GetAccountByID(id)
	if err != nil {
		return err
	}
	fmt.Println(id)
	//account := NeweAccount("John", "Doe")
	return WriteJSON(w, http.StatusOK,account) 
	}

	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w,r)
	}
	
	return fmt.Errorf("Method not allowed %s",r.Method)

}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountReq := new(createAccountRequest)
	if err :=  json.NewDecoder(r.Body).Decode(&createAccountReq); err != nil {
		return err
	}
	account :=  NewAccount(createAccountReq.FirstName,createAccountReq.LastName)
	if err := s.store.CreateAccount(account); err != nil {
		return err
	}
	tokenString,err := createJwt(account)
	if err != nil {	
		return err
	}
	return WriteJSON(w, http.StatusOK, account)

}


func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return err
	}
	if err := s.store.DeleteAccount(id); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, map[string]int{"deleted": id}) 
}

func getId(r *http.Request) (int,error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id,fmt.Errorf("invalid id given %s",idStr)
	}
	return id,nil
}

