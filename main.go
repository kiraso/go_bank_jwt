package main

import (
	"fmt"
	"log"
)




func main() {
	store,err := NewPostgresStorage()
	
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
	 log.Fatal()
	}
	
	fmt.Printf("%+v",store)
	 server := NewAPIServer(":3000",store)
	 server.Run()
}