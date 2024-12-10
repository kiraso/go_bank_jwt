package main

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct{
	Number int64 `json:"number"`
	Password string `json:"password"`	
}
type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount int `json:"amount"`
}
type Account struct {
	ID int `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Number int64 `json:"number"`
	EncryptedPassword string `json:"-"`
	Balance int64 `json:"balance"`
	CreateAt time.Time `json:"createAt"`
}

type createAccountRequest struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Password string `json:"password"`
}

func NewAccount (firstName ,lastName , password string) (*Account,error) {
	encrp,err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Account{
		//ID: rand.Intn(10000),
		FirstName: firstName,
		LastName: lastName,
		EncryptedPassword: string(encrp),
		Number: int64(rand.Intn(10000)),
		CreateAt: time.Now().UTC(),
	},nil
}
