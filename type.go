package main

import (
	"math/rand"
	"time"
)

type Account struct {
	ID int `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Number int64 `json:"number"`
	Balance int64 `json:"balance"`
	CreateAt time.Time `json:"createAt"`
}

type createAccountRequest struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

func NewAccount (firstName ,lastName string) *Account {
	return &Account{
		ID: rand.Intn(10000),
		FirstName: firstName,
		LastName: lastName,
		Number: int64(rand.Intn(10000)),
		CreateAt: time.Now().UTC(),
	}
}
