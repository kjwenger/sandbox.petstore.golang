package main

import (
	sw "github.com/kjwenger/sandbox.petstore.golang/go"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()
	
	log.Fatal(http.ListenAndServe(":8080", router))
}
