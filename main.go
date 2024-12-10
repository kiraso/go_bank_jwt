package main

import (
	"flag"
	"fmt"
	"log"
)



func seedAccount(store Storage,fname,lname, pw string) *Account{
	acc,err := NewAccount(fname,lname,pw)
	if err != nil{
		log.Fatal(err)
	}
	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}
	fmt.Println("new account => ", acc.Number)
	return acc
}

func seedAccounts(store Storage) {
    seedAccount(store, "Annie", "kiraso", "Annie587")
}
func main() {
	seed := flag.Bool("seed", false, "seed accounts db")
	flag.Parse()
	store,err := NewPostgresStorage()
	
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
	 log.Fatal()
	}
	if *seed {
		fmt.Println("seeding the database")
		seedAccounts(store)
	}
	
	fmt.Printf("%+v",store)
	 server := NewAPIServer(":3000",store)
	 server.Run()
}