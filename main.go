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
	fmt.Printf("%+v",store)
	 server := NewAPIServer(":3000",store)
	 server.Run()
}